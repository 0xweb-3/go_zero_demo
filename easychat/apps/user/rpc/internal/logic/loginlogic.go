package logic

import (
	"context"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/user/models"
	"github.com/0xweb-3/go_zero_demo/easychat/pkg/ctxdata"
	"github.com/0xweb-3/go_zero_demo/easychat/pkg/encrypt"
	"github.com/0xweb-3/go_zero_demo/easychat/pkg/xerr"
	"github.com/pkg/errors"
	"time"

	"github.com/0xweb-3/go_zero_demo/easychat/apps/user/rpc/internal/svc"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

//var (
//	ErrPhoneIsNotFound = errors.New("用户未注册")
//	ErrPasswordError   = errors.New("用户密码错误")
//)

var (
	ErrPhoneIsNotFound = xerr.New(xerr.SERVER_COMMON_ERROR, "手机号未注册")
	ErrPasswordError   = xerr.New(xerr.SERVER_COMMON_ERROR, "用户密码错误")
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	// 1. 验证用户是否注册，根据手机号
	userEntity, err := l.svcCtx.UsersModel.FindByPhone(l.ctx, in.GetPhone())
	if err != nil {
		if !errors.Is(err, models.ErrNotFound) {
			return nil, errors.WithStack(ErrPhoneIsNotFound)
		}
		return nil, errors.Wrapf(xerr.NewDBErr(), "find user by phone err %v, req %v", err, in.GetPhone())
	}
	// 2. 验证密码
	if !encrypt.ValidatePasswordHash(in.GetPassword(), userEntity.Password.String) {
		return nil, errors.WithStack(ErrPasswordError)
	}

	// 生成token
	now := time.Now().Unix()
	token, err := ctxdata.GetJwtToken(l.svcCtx.Config.Jwt.AccessSecret, now,
		l.svcCtx.Config.Jwt.AccessExpire, userEntity.Id)

	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "ctxdata get jwt token err %v", err)
	}
	return &user.LoginResp{
		Token:  token,
		Expire: now + l.svcCtx.Config.Jwt.AccessExpire,
	}, nil
}
