package user

import (
	"github.com/0xweb-3/go_zero_demo/easychat/apps/im/ws/internal/svc"
	websocketx "github.com/0xweb-3/go_zero_demo/easychat/apps/im/ws/websocket"
	"github.com/gorilla/websocket"
)

func OnLine(svc *svc.ServiceContext) websocketx.HandlerFunc {
	return func(srv *websocketx.Server, conn *websocket.Conn, msg *websocketx.Message) {
		uids := srv.GetUsers()
		u := srv.GetUsers(conn)
		err := srv.Send(websocketx.NewMessage(u[0], uids), conn)
		srv.Info("err", err)
	}
}
