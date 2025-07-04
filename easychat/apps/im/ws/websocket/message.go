package websocket

type FrameType uint8

const (
	FrameData FrameType = 0x0
	FramePing FrameType = 0x1
	//FrameAck   FrameType = 0x2
	//FrameNoAck FrameType = 0x3
	//FrameErr   FrameType = 0x9

	//FrameHeaders      FrameType = 0x1
	//FramePriority     FrameType = 0x2
	//FrameRSTStream    FrameType = 0x3
	//FrameSettings     FrameType = 0x4
	//FramePushPromise  FrameType = 0x5
	//FrameGoAway       FrameType = 0x7
	//FrameWindowUpdate FrameType = 0x8
	//FrameContinuation FrameType = 0x9
)

type Message struct {
	FrameType `json:"frameType"`
	Method    string `json:"method"` // 请求的方法
	FormId    string `json:"formId"` // 比较数据来源
	Data      any    `json:"data"`   // 数据
}

func NewMessage(fromId string, data any) *Message {
	return &Message{
		FrameType: FrameData,
		FormId:    fromId,
		Data:      data,
	}
}
