package mongo

import (
	"context"
	"fmt"
	"github.com/1111000110/go_service/services/member-service/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var Mongoclient *mongo.Client

func Init() {
	// 设置 MongoDB 连接字符串
	clientOptions := options.Client().ApplyURI(config.AppConfig.Mongo.URI)
	var err error
	// 连接到 MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// 确保连接成功
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("成功连接到 MongoDB!")
	Mongoclient = client
}

// 选择数据库和集合
func GetClient(database string, collec string) *mongo.Collection {
	collection := Mongoclient.Database(database).Collection(collec)
	return collection
}
func GetMemberCollection() *mongo.Collection {
	return GetClient("member", "members")
}
