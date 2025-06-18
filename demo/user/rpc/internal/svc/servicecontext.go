package svc

import (
	"github.com/0xweb-3/go_zero_demo/demo/user/rpc/internal/config"
	"github.com/0xweb-3/go_zero_demo/demo/user/rpc/models"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	UserModel models.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config: c,

		UserModel: models.NewUsersModel(sqlConn, c.Cache),
	}
}
