package parser

import (
	"context"
	"gitee.com/grandeep/org-svc/src/models"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
)

func DecodeUserProto(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(*pb_user_v1.UserProto)
	return models.User{
		BaseModel: models.BaseModel{
			ID:            int(r.Id.Id),
		},
		UserName: r.UserName,
		LoginName: r.LoginName,
		Password: r.Password,
		Mobile: int(r.Mobile),
	}, nil
}

func EncodeUserProto(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(models.User)
	return &pb_user_v1.UserProto{
		UserName: r.UserName,
		LoginName: r.LoginName,
		Password: r.Password,
		Mobile: int64(r.Mobile),
	}, nil
}
