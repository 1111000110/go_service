package service

import (
	"context"
	"github.com/1111000110/go_service/services/post-service/internal/repository/mongo"
)

func DeletePostByPid(ctx context.Context, pid int64) error {
	err := mongo.DeletePostByPid(ctx, pid)
	return err
}
