package mongo

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func MongoDeletePostByMid(ctx context.Context, mid int64) error {
	collection := GetPostCollection()
	filter := bson.M{"mid": mid}
	result, err := collection.DeleteMany(ctx, filter)
	if err != nil {
		log.Printf("Error deleting post by PID: %v", err)
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("post not found")
	}
	return nil
}

func MongoDeletePostByPid(ctx context.Context, pid int64, mid int64) error {
	collection := GetPostCollection()
	filter := bson.M{"pid": pid, "mid": mid}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Printf("Error deleting post by PID: %v", err)
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("post not found")
	}
	return nil
}
