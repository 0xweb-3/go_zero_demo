package svc

import (
	"github.com/0xweb-3/go_zero_demo/easychat/apps/user/models"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/user/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	models.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config: c,

		UsersModel: models.NewUsersModel(sqlConn, c.Cache),
	}
}
