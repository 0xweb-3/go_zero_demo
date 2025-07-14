package websocket

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"sync"
	"time"
)

type AckType int

const (
	NoAck AckType = iota + 1
	OnlyAck
	RigorAck
)

func (t AckType) ToString() string {
	switch t {
	case OnlyAck:
		return "OnlyAck"
	case RigorAck:
		return "RigorAck"
	}
	return "NoAck"
}

type Server struct {
	sync.RWMutex

	opt            *serverOption
	authentication Authentication

	routes map[string]HandlerFunc // key 是具体方法 val是被执行的实际方法
	addr   string
	Patten string

	connToUser map[*Conn]string
	userToConn map[string]*Conn

	upgrader websocket.Upgrader
	logx.Logger
}

func NewServer(addr string, opts ...ServerOptions) *Server {
	opt := NewServerOption(opts...)
	fmt.Printf("Auth implementation type: %T\n", opt.Authentication)
	return &Server{
		routes:         make(map[string]HandlerFunc),
		addr:           addr,
		Patten:         opt.Patten,
		opt:            &opt,
		authentication: opt.Authentication,
		connToUser:     make(map[*Conn]string),
		userToConn:     make(map[string]*Conn),
		upgrader:       websocket.Upgrader{},
		Logger:         logx.WithContext(context.Background()),
	}
}

func (s *Server) ServerWs(w http.ResponseWriter, r *http.Request) {
	// 服务中的错误捕获并记录
	defer func() {
		if r := recover(); r != nil {
			s.Errorf("server Handler ws recover err %v", r)
		}
	}()

	// 获取一个连接对象
	conn := NewConn(s, w, r)
	if conn == nil {
		return
	}
	//conn, err := s.upgrader.Upgrade(w, r, nil)
	//if err != nil {
	//	s.Errorf("upgrade err %v", err)
	//	return
	//}

	// 连接的鉴权
	if !s.authentication.Auth(w, r) {
		//conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprint("没有连接权限")))
		s.Send(&Message{
			FrameType: FrameData,
			Data:      fmt.Sprint("没有连接权限"),
		}, conn)
		conn.Close()
		return
	}

	// 记录连接
	s.addConn(conn, r)

	// 根据连接对象 执行任务处理
	go s.handlerConn(conn)
}

func (s *Server) handlerConn(conn *Conn) {
	uids := s.GetUsers(conn)
	conn.Uid = uids[0]

	// 处理任务
	go s.handlerWrite(conn)
	if s.isAck(nil) {
		go s.readAck(conn)
	}

	for {
		// 获取请求信息
		_, msg, err := conn.ReadMessage()
		if err != nil {
			s.Errorf("websocket conn read message err %v", err)
			s.Close(conn)
			return
		}
		var message Message
		if err = json.Unmarshal(msg, &message); err != nil {
			s.Errorf("json unmarshal, msg %v, err %v", string(msg), err)
			s.Close(conn)
			return
		}

		// 根据不同消息类型进行处理
		switch message.FrameType {
		case FramePing:
			s.Send(&Message{
				FrameType: FramePing,
			}, conn)
		case FrameData:
			// 根据请求的method分发路由并执行
			if handler, ok := s.routes[message.Method]; ok {
				handler(s, conn, &message)
			} else {
				s.Send(&Message{
					FrameType: FrameData,
					Data:      fmt.Sprintf("不存在执行方法 %v, method: %v", err, message.Method),
				}, conn)
				//conn.WriteMessage(websocket.TextMessage,
				//	[]byte(fmt.Sprintf("不存在执行方法 %v, method: %v", err, message.Method)))
			}
		}
		// 给客户端ack应答
		if s.isAck(&message) {
			s.Infof("conn message read ack msg %v", message)
			conn.appendMsgMq(&message)
		} else {
			conn.message <- &message
		}
	}
}

func (s *Server) isAck(message *Message) bool {
	if message == nil {
		return s.opt.ack != NoAck
	}

	return s.opt.ack != NoAck && message.FrameType != FrameNoAck
}

// 读取的ack确认
func (s *Server) readAck(conn *Conn) {
	for {
		select {
		case <-conn.done: // 判断连接是否关闭
			s.Info("close message ack uid %v", conn.Uid)
			return
		default:
		}
		// 从队列中读取新消息
		conn.messageMu.Lock()
		if len(conn.readMessage) == 0 {
			conn.messageMu.Unlock()
			// 增加睡眠
			time.Sleep(100 * time.Millisecond)
			continue
		}

		// 读取第一条消息
		message := conn.readMessage[0]

		// 判断ack方式
		switch s.opt.ack {
		case OnlyAck:
			// 直接回复客户端
			s.Send(&Message{
				FrameType: FrameAck,
				Id:        message.Id,
				AckSeq:    message.AckSeq + 1,
				AckTime:   time.Now(),
			}, conn)
			// 进行业务处理
			// 确认完毕将消息从队列中移除
			conn.readMessage = conn.readMessage[1:]
			conn.messageMu.Unlock()
			conn.message <- message
		case RigorAck:
			// 先回复
			conn.readMessage[0].AckSeq++
			conn.readMessage[0].AckTime = time.Now()
			s.Send(&Message{
				FrameType: FrameAck,
				Id:        message.Id,
				AckSeq:    message.AckSeq,
				AckTime:   time.Now(),
			}, conn)
			s.Info("message ack RigorAck send mid %v, seq %v, time %v",
				message.Id, message.AckSeq, message.AckTime)
			conn.messageMu.Unlock()
			continue
		}
		// 再验证

		// 1. 客户端返回结果，再一次确认
		// 获取客户端序列号
		msgSeq := conn.readMessageSeq[message.Id]
		if msgSeq.AckSeq > message.AckSeq {
			// 确认
			conn.readMessage = conn.readMessage[1:]
			conn.messageMu.Unlock()

			conn.message <- message
			s.Info("message ack RigorAck success mid %v", message.Id)
			continue
		}

		// 2. 客户端没有确认，考虑超过了ack的确认时间
		val := s.opt.ackTimeOut - time.Since(message.AckTime)
		if !message.AckTime.IsZero() && val <= 0 {
			// 2.2 超过结束确认
			delete(conn.readMessageSeq, message.Id)
			conn.readMessage = conn.readMessage[1:]
			conn.messageMu.Unlock()
			continue
		}
		// 2.1 未超过，重新发送
		conn.messageMu.Unlock()
		s.Send(&Message{
			FrameType: FrameAck,
			Id:        message.Id,
			AckSeq:    message.AckSeq,
			AckTime:   time.Now(),
		}, conn)
		time.Sleep(100 * time.Millisecond)
	}
}

// 任务的处理
func (s *Server) handlerWrite(conn *Conn) {
	for {
		select {
		case <-conn.done:
			// 连接关闭
			return
		case message := <-conn.message:
			switch message.FrameType {
			case FramePing:
				s.Send(&Message{
					FrameType: FramePing,
				}, conn)
			case FrameData:
				// 根据请求的method分发路由并执行
				if handler, ok := s.routes[message.Method]; ok {
					handler(s, conn, message)
				} else {
					s.Send(&Message{
						FrameType: FrameData,
						Data:      fmt.Sprintf("不存在执行方法, method: %v", message.Method),
					}, conn)
				}
			}
			// ack消息的清理
			if s.isAck(message) {
				conn.messageMu.Lock()
				delete(conn.readMessageSeq, message.Id)
				conn.messageMu.Unlock()
			}
		}
	}
}

func (s *Server) AddRouters(rs []Route) {
	for _, r := range rs {
		s.routes[r.Method] = r.Handler
	}
}

func (s *Server) Start() {
	http.HandleFunc(s.Patten, s.ServerWs)
	s.Info(http.ListenAndServe(s.addr, nil))
}

func (s *Server) Stop() {
	fmt.Println("websocket server stoping 。。。。。")
}

func (s *Server) addConn(conn *Conn, req *http.Request) {
	uid := s.authentication.UserId(req)
	s.RWMutex.Lock()
	defer s.RWMutex.Unlock()

	// 处理重复登录问题
	if c := s.userToConn[uid]; c != nil {
		// 关闭上一个连接
		c.Close()
	}

	s.connToUser[conn] = uid
	s.userToConn[uid] = conn
}

func (s *Server) GetConn(uid string) *Conn {
	s.RWMutex.RLock()
	defer s.RWMutex.RUnlock()
	return s.userToConn[uid]
}

func (s *Server) GetConns(uids ...string) []*Conn {
	if len(uids) == 0 {
		return nil
	}

	s.RWMutex.RLock()
	defer s.RWMutex.RUnlock()

	res := make([]*Conn, 0, len(uids))
	for _, uid := range uids {
		res = append(res, s.userToConn[uid])
	}
	return res
}

func (s *Server) GetUsers(conns ...*Conn) []string {
	s.RWMutex.RLock()
	defer s.RWMutex.RUnlock()

	var res []string
	if len(conns) == 0 {
		// 获取全部
		res = make([]string, 0, len(s.connToUser))
		for _, uid := range s.connToUser {
			res = append(res, uid)
		}
	} else {
		// 获取部分
		res = make([]string, 0, len(conns))
		for _, conn := range conns {
			res = append(res, s.connToUser[conn])
		}
	}

	return res
}

func (s *Server) Close(conn *Conn) {
	s.RWMutex.Lock()
	defer s.RWMutex.Unlock()

	uid := s.connToUser[conn]
	if uid == "" {
		// 已经被关闭
		return
	}
	delete(s.connToUser, conn)
	delete(s.userToConn, uid)
	conn.Close()
}

func (s *Server) SendByUserId(msg any, sendIds ...string) error {
	if len(sendIds) == 0 {
		return nil
	}
	return s.Send(msg, s.GetConns(sendIds...)...)
}

func (s *Server) Send(msg any, conns ...*Conn) error {
	if len(conns) == 0 {
		return nil
	}
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	for _, conn := range conns {
		if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
			return err
		}
	}
	return nil
}
