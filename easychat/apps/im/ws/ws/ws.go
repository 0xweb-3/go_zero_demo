package ws

import "github.com/0xweb-3/go_zero_demo/easychat/pkg/constants"

type Msg struct {
	constants.MsgType `mapstructure:"msgType"`
	Content           string `mapstructure:"content"`
}

type Chat struct {
	ConversationId string             `mapstructure:"conversationId"`
	ChatType       constants.ChatType `mapstructure:"chatType"`
	SendId         string             `mapstructure:"sendId"`
	RecvId         string             `mapstructure:"recvId"`
	SendTime       int64              `mapstructure:"sendTime"`
	Msg            `mapstructure:"msg"`
}
