package user

import (
	"net/http"

	"github.com/0xweb-3/go_zero_demo/easychat/apps/user/api/internal/logic/user"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/user/api/internal/svc"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 用户登入
func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
