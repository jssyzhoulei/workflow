package parser

import (
	"context"
	"errors"
	"fmt"
	"gitee.com/grandeep/org-svc/src/models"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
	"reflect"
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
		Ststus:    int64(user.Status),
	}, nil
}

func DecodeUserModel(ctx context.Context, res interface{}) (interface{}, error) {
	user, ok := res.(*pb_user_v1.UserProto)
	if !ok {
		return nil, errors.New("error type")
	}
	return models.User{
		BaseModel: models.BaseModel{
			ID: int(user.Id.Id),
		},
		UserName:  user.UserName,
		LoginName: user.LoginName,
		Password:  user.Password,
		Mobile:    int(user.Mobile),
		Status:    int(user.Ststus),
	}, nil
}

func DecodeUsers(_ context.Context, i interface{}) (interface{}, error) {
	if users, ok := i.(*pb_user_v1.Users); ok {
		return users, nil
	}
	err := fmt.Errorf(clientEncodeErr, reflect.TypeOf(i))
	return nil, err
}

func EncodeUserPage(ctx context.Context, i interface{}) (request interface{}, err error) {
	if userPage, ok := i.(*pb_user_v1.UserPage); ok {
		return userPage, nil
	}
	err = fmt.Errorf(clientEncodeErr, reflect.TypeOf(i))
	return
}

func DecodeUsersPage(ctx context.Context, i interface{}) (response interface{}, err error) {
	if usersPage, ok := i.(*pb_user_v1.UsersPage); ok {
		return usersPage, nil
	}
	 err= fmt.Errorf(clientEncodeErr, reflect.TypeOf(i))
	return
}
