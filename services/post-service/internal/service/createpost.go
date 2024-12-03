package service

import (
	"context"
	"fmt"
	"github.com/1111000110/go_service/services/post-service/internal/repository/mongo"
	"github.com/1111000110/go_service/services/post-service/internal/repository/redis"
	"github.com/1111000110/go_service/shared/models"
	"github.com/1111000110/go_service/shared/proto/api/postapi"

	"math/rand"
	"strconv"
	"time"
)

func CreatePost(ctx context.Context, param *postapi.CreatePostRequest) (*postapi.CreatePostResponse, error) {
	post := models.Post{
		Tid:   param.Tid,
		Mid:   param.Mid,
		Pid:   generateUniqueInt64(),
		Title: param.Title,
		Desc:  param.Desc,
		Ct:    time.Now().Unix(),
		Ut:    time.Now().Unix(),
	}
	err := mongo.MongoCreatePost(ctx, &post)
	if err != nil {
		return nil, err
	}
	err = redis.CacheSetPost(ctx, &post)
	if err != nil {
		return nil, err
	}
	resp := &postapi.CreatePostResponse{
		Post: post,
	}
	return resp, nil
}
func generateUniqueInt64() int64 {
	timestamp := time.Now().Unix()
	timestampStr := strconv.FormatInt(timestamp, 10)
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(1000)
	randomStr := fmt.Sprintf("%03d", randomNum)
	finalStr := timestampStr + randomStr
	finalInt, err := strconv.ParseInt(finalStr, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Failed to convert finalStr to int64: %v", err))
	}
	return finalInt
}
