package mongo

import (
	"context"
	"github.com/1111000110/go_service/services/post-service/internal/model"
	"log"
)

func MongoCreatePost(ctx context.Context, post *model.Post) error {
	collection := GetPostCollection()
	// 插入文档
	_, err := collection.InsertOne(ctx, post)
	if err != nil {
		log.Printf("Error creating post: %v", err)
		return err
	}
	return nil
}
