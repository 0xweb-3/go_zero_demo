package config

import "github.com/zeromicro/go-zero/core/service"

type Config struct {
	service.ServiceConf

	ListenOn string

	JwtAuth struct {
		AccessSecret string
	}
}
