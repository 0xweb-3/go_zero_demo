package group

import (
	"context"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/social/rpc/socialclient"
	"github.com/0xweb-3/go_zero_demo/easychat/pkg/constants"
	"github.com/0xweb-3/go_zero_demo/easychat/pkg/ctxdata"

	"github.com/0xweb-3/go_zero_demo/easychat/apps/social/api/internal/svc"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupPutInHandleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 申请进群处理
func NewGroupPutInHandleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupPutInHandleLogic {
	return &GroupPutInHandleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupPutInHandleLogic) GroupPutInHandle(req *types.GroupPutInHandleRep) (resp *types.GroupPutInHandleResp, err error) {
	// todo: add your logic here and delete this line

	_, err = l.svcCtx.Social.GroupPutInHandle(l.ctx, &socialclient.GroupPutInHandleReq{
		GroupReqId:   req.GroupReqId,
		GroupId:      req.GroupId,
		HandleUid:    ctxdata.GetUID(l.ctx),
		HandleResult: req.HandleResult,
	})

	if constants.HandlerResult(req.HandleResult) != constants.PassHandlerResult {
		return
	}

	// todo: 通过后的业务

	return
}
