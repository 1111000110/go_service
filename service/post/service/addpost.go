package service

import (
	"fmt"
	"github.com/1111000110/go_service/api/api/postapi"
	"github.com/1111000110/go_service/api/apidata/postapidata"
	"math/rand"
	"strconv"
	"time"
)

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
func AddPost(param *postapi.AddPostParam) (*postapidata.Post, error) {
	return &postapidata.Post{
		Pid:  generateUniqueInt64(),
		Mid:  param.Mid,
		Tid:  param.Tid,
		Ct:   time.Now().Unix(),
		Ut:   time.Now().Unix(),
		Text: param.Text,
	}, nil
}
