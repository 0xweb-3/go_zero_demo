package main

import (
	"flag"
	"fmt"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/im/ws/internal/config"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/im/ws/internal/handler"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/im/ws/internal/svc"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/im/ws/websocket"
	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/dev/im.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	if err := c.SetUp(); err != nil {
		panic(err)
	}

	srv := websocket.NewServer(c.ListenOn)

	defer srv.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(srv, ctx)

	fmt.Println("start websocket server addr:", c.ListenOn)
	srv.Start()
}
