package main

import (
	"flag"
	"fmt"
	"github.com/0xweb-3/go_zero_demo/easychat/pkg/resultx"
	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/0xweb-3/go_zero_demo/easychat/apps/user/api/internal/config"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/user/api/internal/handler"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/user/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/dev/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	httpx.SetErrorHandlerCtx(resultx.ErrHandler(c.Name))
	httpx.SetOkHandler(resultx.OkHandler)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
