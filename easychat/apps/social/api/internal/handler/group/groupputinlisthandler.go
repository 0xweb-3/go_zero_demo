package group

import (
	"net/http"

	"github.com/0xweb-3/go_zero_demo/easychat/apps/social/api/internal/logic/group"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/social/api/internal/svc"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/social/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 申请进群列表
func GroupPutInListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GroupPutInListRep
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := group.NewGroupPutInListLogic(r.Context(), svcCtx)
		resp, err := l.GroupPutInList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
