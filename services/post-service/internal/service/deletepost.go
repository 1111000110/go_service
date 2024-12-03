package service

import (
	"context"
	"github.com/1111000110/go_service/services/post-service/internal/repository/mongo"
	"github.com/1111000110/go_service/services/post-service/internal/repository/redis"
	"github.com/1111000110/go_service/shared/proto/api/postapi"
)

func DeletePostByPid(ctx context.Context, param *postapi.DeletePostByPidRequest) error {
	err := mongo.MongoDeletePostByPid(ctx, param.Pid, param.Mid)
	if err != nil {
		return err
	}
	err = redis.CacheDeletePostByPid(ctx, param.Pid)
	return err
}
