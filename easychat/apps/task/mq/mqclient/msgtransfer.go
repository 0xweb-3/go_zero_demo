package mqclient

import (
	"context"
	"encoding/json"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/task/mq/mq"
	"github.com/zeromicro/go-queue/kq"
)

type MsgChatTransferClient interface {
	Push(msg *mq.MsgChatTransfer) error
}

type msgChatTransferClient struct {
	pusher *kq.Pusher
}

func NewMsgChatTransferClient(addr []string, topic string, opts ...kq.PushOption,
) MsgChatTransferClient {
	return &msgChatTransferClient{pusher: kq.NewPusher(addr, topic)}
}

func (c *msgChatTransferClient) Push(msg *mq.MsgChatTransfer) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	return c.pusher.Push(context.Background(), string(body))
}
