package svc

import (
	"github.com/0xweb-3/go_zero_demo/easychat/apps/user/models"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/user/rpc/internal/config"
	"github.com/0xweb-3/go_zero_demo/easychat/pkg/constants"
	"github.com/0xweb-3/go_zero_demo/easychat/pkg/ctxdata"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

type ServiceContext struct {
	Config config.Config
	*redis.Redis
	models.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:     c,
		Redis:      redis.MustNewRedis(c.Redisx),
		UsersModel: models.NewUsersModel(sqlConn, c.Cache),
	}
}

func (svc *ServiceContext) SetRootToken() error {
	validDuration := int64(10 * 365 * 24 * 60 * 60) // 10year
	systemToken, err := ctxdata.GetJwtToken(svc.Config.Jwt.AccessSecret, time.Now().Unix(), validDuration, constants.SystemRoot)
	if err != nil {
		return err
	}
	err = svc.Redis.Set(string(constants.RedisKeySystemRootJwtToken), systemToken)
	return err
}
