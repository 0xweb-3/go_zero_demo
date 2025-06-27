package group

import (
	"context"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/social/rpc/socialclient"
	"github.com/0xweb-3/go_zero_demo/easychat/pkg/ctxdata"

	"github.com/0xweb-3/go_zero_demo/easychat/apps/social/api/internal/svc"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupPutInLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 申请进群
func NewGroupPutInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupPutInLogic {
	return &GroupPutInLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupPutInLogic) GroupPutIn(req *types.GroupPutInRep) (resp *types.GroupPutInResp, err error) {
	// todo: add your logic here and delete this line
	uid := ctxdata.GetUID(l.ctx)

	_, err = l.svcCtx.Social.GroupPutin(l.ctx, &socialclient.GroupPutinReq{
		GroupId:    req.GroupId,
		ReqId:      uid,
		ReqMsg:     req.ReqMsg,
		ReqTime:    req.ReqTime,
		JoinSource: int32(req.JoinSource),
	})
	return
}
