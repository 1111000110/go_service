package service

import (
	"context"
	"github.com/1111000110/go_service/services/post-service/internal/repository/mongo"
	"github.com/1111000110/go_service/services/post-service/internal/repository/redis"
)

func DeletePostByMid(ctx context.Context, mid int64) error {
	posts, err := mongo.MongoGetPostsByMids(ctx, []int64{mid})
	if err != nil {
		return err
	}
	err = mongo.MongoDeletePostByMid(ctx, mid)
	if err != nil {
		return err
	}
	for _, post := range *posts {
		err = redis.CacheDeletePostByPid(ctx, post.Pid)
		if err != nil {
			return err
		}
	}
	return err
}
func DeletePostByPid(ctx context.Context, mid int64, pid int64) error {
	err := mongo.MongoDeletePostByPid(ctx, pid, mid)
	if err != nil {
		return err
	}
	err = redis.CacheDeletePostByPid(ctx, pid)
	return err
}
