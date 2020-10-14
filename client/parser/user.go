package parser

import (
	"context"
	"gitee.com/grandeep/org-svc/src/models"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
)

func EncodeUserModel(ctx context.Context, req interface{}) (interface{}, error) {
	user := req.(models.User)
	return &pb_user_v1.UserProto{
		UserName:  user.UserName,
		LoginName: user.LoginName,
		Password:  user.Password,
		Mobile:    int64(user.Mobile),
	}, nil
}

func DecodeUserModel(ctx context.Context, res interface{}) (interface{}, error) {
	user := res.(*pb_user_v1.UserProto)
	return models.User{
		BaseModel: models.BaseModel{
			ID:     int(user.Id.Id),
		},
		UserName:  user.UserName,
		LoginName: user.LoginName,
		Password:  user.Password,
		Mobile:    int(user.Mobile),
	}, nil
}
