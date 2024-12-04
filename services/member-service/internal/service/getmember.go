package service

import (
	"context"
	"github.com/1111000110/go_service/services/member-service/internal/model"
	"github.com/1111000110/go_service/services/member-service/internal/repository/mongo"
)

func GetMemberByMid(ctx context.Context, mid int64) (*model.Member, error) {
	return mongo.MongoGetMemberByMid(ctx, mid)
}
