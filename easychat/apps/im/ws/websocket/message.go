package websocket

type Message struct {
	Method string `json:"method"` // 请求的方法
	FormId string `json:"formId"` // 比较数据来源
	Data   any    `json:"data"`   // 数据
}

func NewMessage(fromId string, data any) *Message {
	return &Message{
		FormId: fromId,
		Data:   data,
	}
}
