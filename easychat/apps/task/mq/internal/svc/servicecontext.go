package svc

import (
	"github.com/0xweb-3/go_zero_demo/easychat/apps/im/immodels"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/im/ws/websocket"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/task/mq/internal/config"
	"github.com/0xweb-3/go_zero_demo/easychat/pkg/constants"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"net/http"
)

type ServiceContext struct {
	config.Config

	websocket.Client
	*redis.Redis

	immodels.ChatLogModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	svc := &ServiceContext{
		Config:       c,
		Redis:        redis.MustNewRedis(c.Redisx),
		ChatLogModel: immodels.NewChatLogModel(c.Mongo.Url, c.Mongo.Db),
	}

	token, err := svc.GetSystemToken()
	if err != nil {
		panic(err)
	}
	header := http.Header{}
	header.Set("Authorization", token)
	svc.Client = websocket.NewClient(c.Ws.Host, websocket.WithClientHeader(header))

	return svc
}

func (svc *ServiceContext) GetSystemToken() (string, error) {
	return svc.Redis.Get(string(constants.RedisKeySystemRootJwtToken))
}
