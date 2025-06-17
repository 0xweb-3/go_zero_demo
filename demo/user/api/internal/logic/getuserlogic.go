package logic

import (
	"context"
	"github.com/0xweb-3/go_zero_demo/demo/user/rpc/userclient"

	"github.com/0xweb-3/go_zero_demo/demo/user/api/internal/svc"
	"github.com/0xweb-3/go_zero_demo/demo/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.UserReq) (resp *types.UserResp, err error) {
	userResp, err := l.svcCtx.User.GetUser(l.ctx, &userclient.GetUserReq{
		Id: "1",
	})

	resp = &types.UserResp{
		Id:    userResp.GetId(),
		Name:  userResp.GetName(),
		Phone: userResp.GetPhone(),
	}
	return
}
