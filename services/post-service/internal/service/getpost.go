package service

import (
	"context"
	"github.com/1111000110/go_service/services/post-service/internal/repository/mongo"
	"github.com/1111000110/go_service/services/post-service/internal/repository/redis"
	"github.com/1111000110/go_service/shared/models"
)

func GetPostByPids(ctx context.Context, pids []int64) (*[]models.Post, error) {
	// 尝试从 Redis 缓存中获取数据
	data, err := redis.CacheGetPostsByPids(ctx, pids)
	if err != nil {
		return nil, err
	}
	if *data != nil {
		return data, nil
	}
	data, err = mongo.MongoGetPostsByPids(ctx, pids)
	if err != nil {
		return nil, err
	}
	err = redis.CacheSetPosts(ctx, data)
	return data, err
}
func GetPostByMids(ctx context.Context, mids []int64) (*[]models.Post, error) {
	data, err := mongo.MongoGetPostsByMids(ctx, mids)
	return data, err
}
