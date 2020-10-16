package parser

import (
	"context"
	"gitee.com/grandeep/org-svc/src/models"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
)

func DecodeUserProto(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(*pb_user_v1.UserProto)
	user := models.User{
		//BaseModel: models.BaseModel{
		//	ID:            int(r.Id.Id),
		//},
		UserName: r.UserName,
		LoginName: r.LoginName,
		Password: r.Password,
		Mobile: int(r.Mobile),
		GroupID: int(r.GroupId),
		UserType: int(r.UserType),
	}
	return user, nil
}

func EncodeUserProto(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(models.User)
	return &pb_user_v1.UserProto{
		Id: &pb_user_v1.Index{
			Id:                   int64(r.ID),
		},
		UserName: r.UserName,
		LoginName: r.LoginName,
		Password: r.Password,
		Mobile: int64(r.Mobile),
	}, nil
}
