package svc

import (
	"github.com/0xweb-3/go_zero_demo/demo/user/api/internal/config"
	"github.com/0xweb-3/go_zero_demo/demo/user/api/internal/middleware"
	"github.com/0xweb-3/go_zero_demo/demo/user/rpc/userclient"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	userclient.User

	LoginVerification rest.Middleware // 在demo/user/api/internal/handler/routes.go下对应的名称
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		User: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),

		// 配置中间件
		LoginVerification: middleware.NewLoginVerificationMiddleware().Handle,
	}
}
