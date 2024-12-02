package service

import (
	"context"
	"github.com/1111000110/go_service/services/post-service/internal/model"
	"github.com/1111000110/go_service/services/post-service/internal/repository/mongo"
)

func GetPostByPids(ctx context.Context, pids []int64) (*[]model.Post, error) {
	data, err := mongo.GetPostsByPids(ctx, pids)
	return data, err
}
func GetPostByMids(ctx context.Context, mids []int64) (*[]model.Post, error) {
	data, err := mongo.GetPostsByMids(ctx, mids)
	return data, err
}
