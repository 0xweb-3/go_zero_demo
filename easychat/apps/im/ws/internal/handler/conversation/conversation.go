package conversation

import (
	"context"
	"fmt"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/im/ws/internal/logic"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/im/ws/internal/svc"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/im/ws/websocket"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/im/ws/ws"
	"github.com/0xweb-3/go_zero_demo/easychat/pkg/constants"
	"github.com/mitchellh/mapstructure"
	"time"
)

func Chat(svc *svc.ServiceContext) websocket.HandlerFunc {
	return func(srv *websocket.Server, conn *websocket.Conn, msg *websocket.Message) {
		// 接收客户端发送的消息
		data := new(ws.Chat)
		if err := mapstructure.Decode(msg.Data, &data); err != nil {
			srv.Send(websocket.NewErrMessage(err), conn)
			return
		}
		switch data.ChatType {
		case constants.SingleChatType:
			err := logic.NewConversation(context.Background(), srv, svc).SingleChat(data, conn.Uid)
			if err != nil {
				srv.Send(websocket.NewErrMessage(err), conn)
				return
			}
			fmt.Println("数据：", conn.Uid, "--", data.RecvId)
			srv.SendByUserId(websocket.NewMessage(conn.Uid, ws.Chat{
				ConversationId: data.ConversationId,
				ChatType:       data.ChatType,
				SendId:         conn.Uid,
				RecvId:         data.RecvId,
				SendTime:       time.Now().UnixNano(),
				Msg:            data.Msg,
			}), data.RecvId)
		}

	}
}
