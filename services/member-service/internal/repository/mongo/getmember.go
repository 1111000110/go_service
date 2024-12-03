package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoGetLastMid(ctx context.Context) (int64, error) {
	collection := GetMemberCollection()
	sortOptions := options.FindOne().SetSort(bson.D{{"mid", -1}})
	var result struct {
		Mid int64 `bson:"mid"`
	}
	err := collection.FindOne(ctx, bson.D{}, sortOptions).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return 0, nil
		}
		return 0, err
	}
	return result.Mid, nil
}
