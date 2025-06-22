package logic

import (
	"github.com/0xweb-3/go_zero_demo/easychat/apps/user/rpc/internal/config"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/user/rpc/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
	"path/filepath"
)

var svcCtx *svc.ServiceContext

func init() {
	var c config.Config
	conf.MustLoad(filepath.Join("../../etc/dev/user.yaml"), &c)
	svcCtx = svc.NewServiceContext(c)
}
