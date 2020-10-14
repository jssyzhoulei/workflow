package parser

import (
	"context"
	"gitee.com/grandeep/org-svc/src/models"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
)

func EncodeUserModel(ctx context.Context, req interface{}) (interface{}, error) {
	user := req.(models.User2)
	return &pb_user_v1.UserProto{
		UserId:                   int64(user.UserId),
		UserName:                 user.UserName,
	}, nil
}

func DecodeUserModel(ctx context.Context, res interface{}) (interface{}, error) {
	user := res.(*pb_user_v1.UserProto)
	return models.User2{
		UserId:                   user.UserId,
		UserName:                 user.UserName,
	}, nil
}
