package user

import (
	"context"
	"fmt"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/user/api/internal/svc"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/user/api/internal/types"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/user/rpc/user"
	"github.com/0xweb-3/go_zero_demo/easychat/pkg/ctxdata"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户信息
func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	uid := ctxdata.GetUID(l.ctx)
	if uid == "" {
		return nil, fmt.Errorf("uid is empty")
	}

	userResp, err := l.svcCtx.User.GetUserInfo(l.ctx, &user.GetUserInfoReq{Id: uid})

	if err != nil {
		return nil, err
	}
	var res types.User
	copier.Copy(&res, userResp.User)
	return &types.UserInfoResp{
		Info: res,
	}, nil
}
