package parser

import (
	"context"
	"errors"
	"gitee.com/grandeep/org-svc/src/models"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
)

func EncodeUserModel(ctx context.Context, req interface{}) (interface{}, error) {
	user, ok := req.(models.User)
	if !ok {
		return nil, errors.New("error type")
	}
	return &pb_user_v1.UserProto{
		UserName:  user.UserName,
		LoginName: user.LoginName,
		Password:  user.Password,
		Mobile:    int64(user.Mobile),
		GroupId:   int64(user.GroupID),
		UserType:  int64(user.UserType),
	}, nil
}

func DecodeUserModel(ctx context.Context, res interface{}) (interface{}, error) {
	user, ok := res.(*pb_user_v1.UserProto)
	if !ok {
		return nil, errors.New("error type")
	}
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
