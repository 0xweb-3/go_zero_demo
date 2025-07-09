package svc

import "github.com/0xweb-3/go_zero_demo/easychat/apps/task/mq/internal/config"

type ServiceContext struct {
	config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
