package msgTransfer

import (
	"context"
	"encoding/json"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/im/immodels"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/im/ws/websocket"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/task/mq/internal/svc"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/task/mq/mq"
	"github.com/0xweb-3/go_zero_demo/easychat/pkg/constants"
	"github.com/zeromicro/go-zero/core/logx"
)

type MsgChatTransfer struct {
	logx.Logger
	svc *svc.ServiceContext
}

func NewMsgChatTransfer(svc *svc.ServiceContext) *MsgChatTransfer {
	return &MsgChatTransfer{
		Logger: logx.WithContext(context.Background()),
		svc:    svc,
	}
}

func (m *MsgChatTransfer) Consume(ctx context.Context, key, value string) error {
	var (
		data mq.MsgChatTransfer
	)

	if err := json.Unmarshal([]byte(value), &data); err != nil {
		return err
	}

	// 记录数据
	if err := m.addChatLog(ctx, &data); err != nil {
		return err
	}
	// 推送消息
	return m.svc.Client.Snd(websocket.Message{
		FrameType: websocket.FrameData,
		Method:    "push",
		FormId:    constants.SystemRoot,
		Data:      data,
	})
}

func (m *MsgChatTransfer) addChatLog(ctx context.Context, data *mq.MsgChatTransfer) error {
	// 记录到聊天到数据库
	chatLog := immodels.ChatLog{
		ConversationID: data.ConversationId,
		SendId:         data.SendId,
		RecvId:         data.RecvId,
		ChatType:       constants.ChatType(data.ChatType),
		MsgFrom:        0,
		MsgType:        constants.MsgType(data.MsgType),
		MsgContent:     data.Content,
		SendTime:       data.SendTime,
	}
	return m.svc.ChatLogModel.Insert(ctx, &chatLog)
}
