package logic

import (
	"context"
	"errors"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/user/models"
	"github.com/0xweb-3/go_zero_demo/easychat/pkg/ctxdata"
	"github.com/0xweb-3/go_zero_demo/easychat/pkg/encrypt"
	"time"

	"github.com/0xweb-3/go_zero_demo/easychat/apps/user/rpc/internal/svc"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrPhoneIsNotFound = errors.New("用户未注册")
	ErrPasswordError   = errors.New("用户密码错误")
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
			return nil, ErrPhoneIsNotFound
		}
		return nil, err
	}
	// 2. 验证密码
	if !encrypt.ValidatePasswordHash(in.GetPassword(), userEntity.Password.String) {
		return nil, ErrPasswordError
	}

	// 生成token
	now := time.Now().Unix()
	token, err := ctxdata.GetJwtToken(l.svcCtx.Config.Jwt.AccessSecret, now,
		l.svcCtx.Config.Jwt.AccessExpire, userEntity.Id)

	if err != nil {
		return nil, err
	}
	return &user.LoginResp{
		Token:  token,
		Expire: now + l.svcCtx.Config.Jwt.AccessExpire,
	}, nil
}
