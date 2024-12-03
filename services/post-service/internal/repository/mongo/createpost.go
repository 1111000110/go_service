package mongo

import (
	"context"
	"github.com/1111000110/go_service/shared/models"
	"log"
)

func MongoCreatePost(ctx context.Context, post *models.Post) error {
	collection := GetPostCollection()
	// 插入文档
	_, err := collection.InsertOne(ctx, post)
	if err != nil {
		log.Printf("Error creating post: %v", err)
		return err
	}
	return nil
}
