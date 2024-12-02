package mongo

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

// DeletePostByPid 根据 PID 删除 Post
func DeletePostByPid(ctx context.Context, pid int64) error {
	// 获取 MongoDB 的 collection
	collection := GetPostCollection()
	// 构建查询条件
	filter := bson.M{"_id": pid}

	// 执行删除操作
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Printf("Error deleting post by PID: %v", err)
		return err
	}

	// 检查是否成功删除
	if result.DeletedCount == 0 {
		return errors.New("post not found")
	}

	return nil
}
