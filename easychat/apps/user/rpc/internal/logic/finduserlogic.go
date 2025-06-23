package logic

import (
	"context"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/user/models"
	"github.com/jinzhu/copier"

	"github.com/0xweb-3/go_zero_demo/easychat/apps/user/rpc/internal/svc"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserLogic {
	return &FindUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindUserLogic) FindUser(in *user.FindUserReq) (*user.FindUserResp, error) {
	var (
		userEntitys []*models.Users
		err         error
	)

	if in.Phone != "" {
		userEntity, err := l.svcCtx.UsersModel.FindByPhone(l.ctx, in.GetPhone())
		if err == nil {
			userEntitys = append(userEntitys, userEntity)
		}
	} else if in.Name != "" {
		userEntitys, err = l.svcCtx.UsersModel.ListByName(l.ctx, in.GetName())
	} else if in.Ids != nil {
		userEntitys, err = l.svcCtx.UsersModel.ListByIds(l.ctx, in.GetIds())
	}
	if err != nil {
		return nil, err
	}
	var resp []*user.UserEntity
	copier.Copy(&resp, &userEntitys)
	return &user.FindUserResp{
		User: resp,
	}, nil
}
