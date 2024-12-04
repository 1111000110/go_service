package postapi

import (
	"context"
	"github.com/1111000110/go_service/shared/proto/api"
)

// 配置文件写入，临时这样写
const Url = "http://127.0.0.1:8080"

type PostApi struct {
	api.BaseApi
}

func NewMemberApi(ctx context.Context) *PostApi {
	return &PostApi{
		api.BaseApi{
			Ctx: ctx,
			Url: Url,
		},
	}
}
