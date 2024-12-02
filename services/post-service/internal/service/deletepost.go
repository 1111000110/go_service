package service

import (
	"context"
	"github.com/1111000110/go_service/services/post-service/internal/repository/mongo"
	"github.com/1111000110/go_service/services/post-service/internal/repository/redis"
)

func DeletePostByPid(ctx context.Context, pid int64) error {
	err := redis.CacheDeletePostByPid(ctx, pid)
	if err != nil {
		return err
	}
	err = mongo.MongoDeletePostByPid(ctx, pid)
	return err
}
