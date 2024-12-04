package service

import (
	"context"
	"errors"
	"github.com/1111000110/go_service/services/member-service/internal/repository/mongo"
	"github.com/1111000110/go_service/services/member-service/internal/utils"
	"github.com/1111000110/go_service/shared/proto/api/memberapi"
	"github.com/1111000110/go_service/shared/proto/api/postapi"
)

func DeleteMember(ctx context.Context, param *memberapi.DeletememberRequest) (*memberapi.DeletememberResponse, error) {
	data, err := mongo.MongoGetMemberByMid(ctx, param.Mid)
	if err != nil {
		return nil, err
	}
	if !utils.ComparePhone(param.Phone, data.Phone) {
		return nil, errors.New("phone err")
	}
	if !utils.ComparePasswords(param.Password, data.Password) {
		return nil, errors.New("password err")
	}
	deletedata, err := postapi.NewMemberApi(ctx).DeletePostByMid(&postapi.DeletePostByMidRequest{Mid: param.Mid})
	if err != nil || deletedata.IsDeleted == false {
		if err != nil {
			return nil, err
		}
		return nil, errors.New("delete post err")
	}
	err = mongo.MongoDeleteMemberByMid(ctx, param.Mid)
	if err != nil {
		return nil, err
	}
	return &memberapi.DeletememberResponse{
		IsDeleted: true,
	}, nil
}
