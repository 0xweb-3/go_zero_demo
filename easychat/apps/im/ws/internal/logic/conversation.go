package logic

import (
	"context"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/im/immodels"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/im/ws/internal/svc"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/im/ws/websocket"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/im/ws/ws"
	"github.com/0xweb-3/go_zero_demo/easychat/pkg/wuid"
	"time"
)

type Conversation struct {
	ctx context.Context
	srv *websocket.Server
	svc *svc.ServiceContext
}

func NewConversation(ctx context.Context, srv *websocket.Server, svc *svc.ServiceContext) *Conversation {
	return &Conversation{
		ctx: ctx,
		srv: srv,
		svc: svc,
	}
}

func (l *Conversation) SingleChat(data *ws.Chat, userId string) error {
	if data.ConversationId == "" {
		data.ConversationId = wuid.CombineId(data.RecvId, data.SendId)
	}

	// 记录到聊天到数据库
	chatLog := immodels.ChatLog{
		ConversationID: data.ConversationId,
		SendId:         data.SendId,
		RecvId:         data.RecvId,
		ChatType:       data.ChatType,
		MsgFrom:        0,
		MsgType:        data.MsgType,
		MsgContent:     data.Content,
		SendTime:       time.Now().UnixNano(),
	}

	err := l.svc.ChatLogModel.Insert(l.ctx, &chatLog)
	return err
}
