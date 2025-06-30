package svc

import "github.com/0xweb-3/go_zero_demo/easychat/apps/im/ws/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
