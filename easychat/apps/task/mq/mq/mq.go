package mq

type MsgChatTransfer struct {
	ConversationId string `json:"conversationId"`
	ChatType       int    `json:"chatType"`
	SendId         string `json:"sendId"`
	RecvId         string `json:"recvId"`
	SendTime       int64  `json:"sendTime"`
	MsgType        int    `json:"msgType"`
	Content        string `json:"content"`
}
