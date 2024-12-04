package memberapi

import (
	"context"
	"github.com/1111000110/go_service/shared/proto/api"
)

// 配置文件写入，临时这样写
const Url = "http://127.0.0.1:8081"

type MemberApi struct {
	api.BaseApi
}

func NewMemberApi(ctx context.Context) *MemberApi {
	return &MemberApi{
		api.BaseApi{
			Ctx: ctx,
			Url: Url,
		},
	}
}
