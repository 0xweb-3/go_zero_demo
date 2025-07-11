package svc

import (
	"github.com/0xweb-3/go_zero_demo/easychat/apps/im/immodels"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/im/ws/internal/config"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/task/mq/mqclient"
)

type ServiceContext struct {
	Config config.Config

	immodels.ChatLogModel

	mqclient.MsgChatTransferClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                c,
		ChatLogModel:          immodels.NewChatLogModel(c.Mongo.Url, c.Mongo.Db),
		MsgChatTransferClient: mqclient.NewMsgChatTransferClient(c.MsgChatTransfer.Addrs, c.MsgChatTransfer.Topic),
	}
}
