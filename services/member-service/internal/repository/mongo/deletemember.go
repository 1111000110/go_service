package mongo

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func MongoDeleteMemberByMid(ctx context.Context, mid int64) error {
	collection := GetMemberCollection()
	filter := bson.M{"mid": mid}
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
