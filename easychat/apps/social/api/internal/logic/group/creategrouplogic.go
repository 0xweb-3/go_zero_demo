package group

import (
	"context"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/social/rpc/socialclient"
	"github.com/0xweb-3/go_zero_demo/easychat/pkg/ctxdata"

	"github.com/0xweb-3/go_zero_demo/easychat/apps/social/api/internal/svc"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创群
func NewCreateGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGroupLogic {
	return &CreateGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateGroupLogic) CreateGroup(req *types.GroupCreateReq) (resp *types.GroupCreateResp, err error) {
	// todo: add your logic here and delete this line
	uid := ctxdata.GetUID(l.ctx)

	// 创建群
	_, err = l.svcCtx.Social.GroupCreate(l.ctx, &socialclient.GroupCreateReq{
		Name:       req.Name,
		Icon:       req.Icon,
		CreatorUid: uid,
	})
	if err != nil {
		return nil, err
	}

	return
}
