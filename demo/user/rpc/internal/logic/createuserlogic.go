package logic

import (
	"context"
	"github.com/0xweb-3/go_zero_demo/demo/user/rpc/models"

	"github.com/0xweb-3/go_zero_demo/demo/user/rpc/internal/svc"
	"github.com/0xweb-3/go_zero_demo/demo/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *user.CreateUserReq) (*user.CreateUserResp, error) {
	_, err := l.svcCtx.UserModel.Insert(l.ctx, &models.Users{
		Id:    in.GetId(),
		Name:  in.GetName(),
		Phone: in.GetPhone(),
	})
	return &user.CreateUserResp{}, err
}
