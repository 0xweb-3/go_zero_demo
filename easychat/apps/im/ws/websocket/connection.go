package websocket

import (
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
	"time"
)

type Conn struct {
	idleMu sync.Mutex
	Uid    string

	*websocket.Conn
	s *Server

	messageMu      sync.Mutex          // 消息队列的锁
	readMessage    []*Message          // 读消息队列
	readMessageSeq map[string]*Message // 读消息队列的序列化
	message        chan *Message       // ACK确认消息发送 给任务处理部分

	idle              time.Time     // 最近一次活跃时间（空闲开始的时间点）
	maxConnectionIdle time.Duration // 允许的最大空闲时间
	done              chan struct{} // 通知 keepalive 协程退出
}

func NewConn(s *Server, w http.ResponseWriter, r *http.Request) *Conn {
	c, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		s.Errorf("Upgrade err %v", err)
		return nil
	}
	conn := &Conn{
		Conn:              c,
		s:                 s,
		idle:              time.Now(),
		maxConnectionIdle: s.opt.MaxConnectionIdle,
		readMessage:       make([]*Message, 0, 2),
		readMessageSeq:    make(map[string]*Message, 2),
		message:           make(chan *Message, 1), // 1 为了减少数据投递中的阻塞情况，保障收发的顺序性
		done:              make(chan struct{}),
	}
	go conn.keepalive()
	return conn
}

func (c *Conn) appendMsgMq(msg *Message) {
	c.messageMu.Lock()
	defer c.messageMu.Unlock()

	// 读取队列中
	if m, ok := c.readMessageSeq[msg.Id]; ok {
		// 已经有消息的记录，改消息已经存在ack的确认
		if len(c.readMessage) == 0 {
			// 队列中已经没有消息，说明ack过程已经完成，消息被删除
			return
		}

		// 当前ack需要大于已经记录的ack 序号，才说说明是执行后续流程
		if m.AckSeq >= msg.AckSeq {
			// 可能是客户端重复发送的ack消息
			return
		}
		c.readMessageSeq[msg.Id] = msg

		return
	}
	// 还没没有进行ack的确认，避免客户端重复发送多余的ack消息
	if msg.FrameType == FrameAck {
		return
	}

	c.readMessage = append(c.readMessage, msg)
	c.readMessageSeq[msg.Id] = msg
}

func (c *Conn) ReadMessage() (messageType int, p []byte, err error) {
	messageType, p, err = c.Conn.ReadMessage()
	c.idleMu.Lock()
	defer c.idleMu.Unlock()
	c.idle = time.Time{} // 表示有连接非空闲
	return
}

func (c *Conn) WriteMessage(messageType int, data []byte) error {
	c.idleMu.Lock()
	defer c.idleMu.Unlock()
	err := c.Conn.WriteMessage(messageType, data)
	c.idle = time.Now()
	return err
}

func (c *Conn) Close() error {
	select {
	case <-c.done:
	default:
		close(c.done)
	}
	return c.Conn.Close()
}

func (c *Conn) keepalive() {
	idleTimer := time.NewTimer(c.maxConnectionIdle) // 对空对定时器进行初始化

	defer func() {
		idleTimer.Stop()
	}()

	for {
		select {
		case <-idleTimer.C:
			c.idleMu.Lock()
			idle := c.idle
			if idle.IsZero() { // The connection is non-idle.
				c.idleMu.Unlock()
				idleTimer.Reset(c.maxConnectionIdle)
				continue
			}
			val := c.maxConnectionIdle - time.Since(idle)
			c.idleMu.Unlock()
			if val <= 0 {
				c.s.Close(c)
				return
			}
			idleTimer.Reset(val)
		case <-c.done:
			return
		}
	}
}
