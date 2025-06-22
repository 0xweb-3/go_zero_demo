package logic

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/user/models"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/user/rpc/internal/svc"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/user/rpc/user"
	"github.com/0xweb-3/go_zero_demo/easychat/pkg/ctxdata"
	"github.com/0xweb-3/go_zero_demo/easychat/pkg/encrypt"
	"github.com/0xweb-3/go_zero_demo/easychat/pkg/wuid"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrPhoneIsRegister = errors.New("用户已经注册过")
	ErrPasswordIsEmpty = errors.New("密码为空")
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterResp, error) {
	// 1. 验证用户是否注册，根据手机号
	userEntity, err := l.svcCtx.UsersModel.FindByPhone(l.ctx, in.GetPhone())
	if err != nil && !errors.Is(err, models.ErrNotFound) {
		return nil, err
	}

	if userEntity != nil {
		return nil, ErrPhoneIsRegister
	}
	id, err := wuid.GetSonyflakeID(nil)
	if err != nil {
		return nil, nil
	}
	// 定义用户数据
	userEntity = &models.Users{
		Id:       fmt.Sprintf("%d", id),
		Avatar:   in.GetAvatar(),
		Nickname: in.GetNickname(),
		Phone:    in.GetPhone(),
		Password: sql.NullString{},
		Sex: sql.NullInt64{
			Int64: int64(in.GetSex()), // 存实际值
			Valid: true,               // 为 true 表示写进去
		},
	}

	if len(in.GetPassword()) == 0 {
		return nil, ErrPasswordIsEmpty
	}

	genPassword, err := encrypt.GeneratePasswordHash(in.GetPassword())
	if err != nil {
		return nil, err
	}

	userEntity.Password = sql.NullString{
		String: genPassword,
		Valid:  true,
	}

	_, err = l.svcCtx.UsersModel.Insert(l.ctx, userEntity)
	if err != nil {
		return nil, err
	}

	// 生成token
	now := time.Now().Unix()
	token, err := ctxdata.GetJwtToken(l.svcCtx.Config.Jwt.AccessSecret, now,
		l.svcCtx.Config.Jwt.AccessExpire, userEntity.Id)

	if err != nil {
		return nil, err
	}
	return &user.RegisterResp{
		Token:  token,
		Expire: now + l.svcCtx.Config.Jwt.AccessExpire,
	}, nil
}
