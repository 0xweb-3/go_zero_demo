package logic

import (
	"context"
	"fmt"

	"github.com/0xweb-3/go_zero_demo/demo/user/rpc/internal/svc"
	"github.com/0xweb-3/go_zero_demo/demo/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

type User struct {
	Id    string
	Name  string
	Phone string
}

var users = map[string]*User{
	"1": {
		Id:    "1",
		Name:  "xin",
		Phone: "13997742711",
	},
	"2": {
		Id:    "2",
		Name:  "bing",
		Phone: "15102724511",
	},
}

func (l *GetUserLogic) GetUser(in *user.GetUserReq) (*user.GetUserResp, error) {
	if u, ok := users[in.GetId()]; ok {
		return &user.GetUserResp{
			Id:    u.Id,
			Name:  u.Name,
			Phone: u.Phone,
		}, nil
	}

	return nil, fmt.Errorf("GetUser err")
}
