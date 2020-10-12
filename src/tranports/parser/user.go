package parser

import (
	"context"
	"gitee.com/grandeep/org-svc/src/models"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
)

func DecodeUserProto(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(*pb_user_v1.UserProto)
	return models.User{
		UserId: r.UserId,
		UserName: r.UserName,
	}, nil
}

func EncodeUserProto(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(models.User)
	return &pb_user_v1.UserProto{
		UserId:   r.UserId,
		UserName: r.UserName,
	}, nil
}
