package service

import (
	"context"
	"github.com/1111000110/go_service/services/member-service/internal/model"
	"github.com/1111000110/go_service/services/member-service/internal/repository/mongo"
	"github.com/1111000110/go_service/services/member-service/internal/utils"
	"github.com/1111000110/go_service/shared/proto/api/memberapi"
)

func CreateMember(ctx context.Context, param *memberapi.CreateMemberRequest) (*memberapi.CreateMemberResponse, error) {
	lastmid, err := mongo.MongoGetLastMid(ctx)
	mid := lastmid + 1
	if err != nil {
		return nil, err
	}
	phone, err := utils.EncryptPhoneNumber(param.Phone)
	if err != nil {
		return nil, err
	}
	password := utils.EncryptPassword(param.Password)
	memberdata := &model.Member{
		Username: param.Username,
		Mid:      mid,
		Name:     param.Name,
		Birthday: param.Birthday,
		Sex:      param.Sex,
		Phone:    phone,
		Password: password,
	}
	err = mongo.MongoCreatePost(ctx, memberdata)
	if err != nil {
		return nil, err
	}
	resp := &memberapi.CreateMemberResponse{
		Mid: mid,
	}
	return resp, nil
}
