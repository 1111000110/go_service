package mongo

import (
	"context"
	"github.com/1111000110/go_service/services/post-service/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func GetPostsByPids(ctx context.Context, pids []int64) (*[]model.Post, error) {
	// 获取 MongoDB 的 collection
	collection := GetPostCollection()
	// 创建查询条件，使用 $in 操作符来查询多个 pid
	filter := bson.M{"pid": bson.M{"$in": pids}} // 使用 $in 查询多个 id

	// 执行查询，获取多个结果
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Printf("Error finding posts by IDs: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)
	var posts []model.Post
	// 解码查询结果
	for cursor.Next(ctx) {
		var post model.Post
		if err := cursor.Decode(&post); err != nil {
			log.Printf("Error decoding post: %v", err)
			return nil, err
		}
		posts = append(posts, post)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Cursor error: %v", err)
		return nil, err
	}

	return &posts, nil
}
func GetPostsByMids(ctx context.Context, mids []int64) (*[]model.Post, error) {
	collection := GetPostCollection()
	filter := bson.M{"mid": bson.M{"$in": mids}}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Printf("Error finding posts by mids: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)
	var posts []model.Post
	for cursor.Next(ctx) {
		var post model.Post
		if err := cursor.Decode(&post); err != nil {
			log.Printf("Error decoding post: %v", err)
			return nil, err
		}
		posts = append(posts, post)
	}
	if err := cursor.Err(); err != nil {
		log.Printf("Cursor error: %v", err)
		return nil, err
	}
	return &posts, nil
}
