// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package handler

import (
	"net/http"

	"github.com/0xweb-3/go_zero_demo/demo/user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/user",
				Handler: getUserHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.LoginVerification},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/userinfo",
					Handler: getUserInfoHandler(serverCtx),
				},
			}...,
		),
	)
}
