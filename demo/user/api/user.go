package main

import (
	"flag"
	"fmt"
	"github.com/0xweb-3/go_zero_demo/demo/user/api/internal/config"
	"github.com/0xweb-3/go_zero_demo/demo/user/api/internal/handler"
	"github.com/0xweb-3/go_zero_demo/demo/user/api/internal/svc"
	"github.com/0xweb-3/go_zero_demo/easychat/pkg/configserver"
	"github.com/zeromicro/go-zero/core/proc"
	"github.com/zeromicro/go-zero/rest"
	"sync"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")
var wg sync.WaitGroup

func main() {
	flag.Parse()

	var c config.Config
	//conf.MustLoad(*configFile, &c)

	err := configserver.NewConfigServer(*configFile, configserver.NewSail(&configserver.Config{
		ETCDEndpoints:  "192.168.21.5:2379",
		ProjectKey:     "98c6f2c2287f4c73cea3d40ae7ec3ff2",
		Namespace:      "user",
		Configs:        "user-rpc.yaml",
		ConfigFilePath: "./sail_conf",
		LogLevel:       "DEBUG",
	})).MustLoad(&c, func(bytes []byte) error {
		var c config.Config
		configserver.LoadFromJsonBytes(bytes, &c)
		fmt.Println("更新后的配置", c)

		proc.WrapUp() // 停止服务对请求的接收

		// 重新启动服务
		wg.Add(1)
		go func(c config.Config) {
			defer wg.Done()
			Run(c)
		}(c)

		return nil
	})

	if err != nil {
		panic(err)
	}

	wg.Add(1)
	go func(c config.Config) {
		defer wg.Done()
		Run(c)
	}(c)
	wg.Wait()
}

func Run(c config.Config) {
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
