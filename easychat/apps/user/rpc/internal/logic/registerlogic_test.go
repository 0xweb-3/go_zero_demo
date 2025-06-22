package logic

import (
	"context"
	"github.com/0xweb-3/go_zero_demo/easychat/apps/user/rpc/user"
	"testing"
)

func TestRegisterLogic_Register(t *testing.T) {
	type args struct {
		in *user.RegisterReq
	}
	tests := []struct {
		name      string
		args      args
		wantPrint bool
		wantWrr   bool
	}{
		{
			name: "1", args: args{
				in: &user.RegisterReq{
					Phone:    "15102724511",
					Nickname: "xin",
					Password: "123456",
					Avatar:   "avatar.jpg",
					Sex:      1,
				},
			},
			wantPrint: true,
			wantWrr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewRegisterLogic(context.Background(), svcCtx)
			got, err := l.Register(tt.args.in)
			if err != nil {
				t.Errorf("Register() error = %v", err)
				return
			}
			if tt.wantPrint {
				t.Log(tt.name, got)
			}
		})
	}
}
