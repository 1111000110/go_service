package mongo

import (
	"context"
	"github.com/1111000110/go_service/services/member-service/internal/model"
	"log"
)

func MongoCreatePost(ctx context.Context, member *model.Member) error {
	collection := GetMemberCollection()
	// 插入文档
	_, err := collection.InsertOne(ctx, member)
	if err != nil {
		log.Printf("Error creating post: %v", err)
		return err
	}
	return nil
}
