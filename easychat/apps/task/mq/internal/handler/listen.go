package handler

import (
	"github.com/0xweb-3/go_zero_demo/easychat/apps/task/mq/internal/handler/msgTransfer"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/task/mq/internal/svc"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
)

type Listen struct {
	svc *svc.ServiceContext
}

func NewListen(svc *svc.ServiceContext) *Listen {
	return &Listen{svc: svc}
}

func (l *Listen) Service() []service.Service {
	return []service.Service{
		// 这里可以加载多个消费者
		kq.MustNewQueue(l.svc.Config.MsgChatTransfer, msgTransfer.NewMsgChatTransfer(l.svc)),
	}
}
