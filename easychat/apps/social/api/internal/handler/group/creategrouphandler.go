package group

import (
	"net/http"

	"github.com/0xweb-3/go_zero_demo/easychat/apps/social/api/internal/logic/group"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/social/api/internal/svc"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/social/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 创群
func CreateGroupHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GroupCreateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := group.NewCreateGroupLogic(r.Context(), svcCtx)
		resp, err := l.CreateGroup(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
