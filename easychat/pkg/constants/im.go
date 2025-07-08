package constants

type MsgType int

const (
	TextMsgType MsgType = iota + 1
)

type ChatType int

const (
	GroupChatType ChatType = iota + 1
	SingleChatType
)
