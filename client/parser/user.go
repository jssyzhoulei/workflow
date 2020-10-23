package parser

import (
	"context"
	"errors"
	"fmt"
	"gitee.com/grandeep/org-svc/src/models"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
	"reflect"
)

func EncodeUserRoleDTO(ctx context.Context, req interface{}) (interface{}, error) {
	r, ok := req.(models.UserRolesDTO)
	if !ok {
		return models.User{}, fmt.Errorf(clientEncodeErr, reflect.TypeOf(req))
	}
	userProto := &pb_user_v1.UserProto{
		UserName: r.UserName,
		LoginName: r.LoginName,
		Password: r.Password,
		Mobile: int64(r.Mobile),
		GroupId: int64(r.GroupID),
		UserType: int64(r.UserType),
		Ststus: int64(r.Status),
	}
	for _, v := range r.RoleIDs {
		userProto.RoleIds = append(userProto.RoleIds, &pb_user_v1.Index{Id:int64(v)})
	}
	if r.ID != 0 {
		userProto.Id = &pb_user_v1.Index{Id: int64(r.ID)}
	}
	return userProto, nil
}

func DecodeUserModel(ctx context.Context, res interface{}) (interface{}, error) {
	userProto, ok := res.(*pb_user_v1.UserProto)
	if !ok {
		return nil, errors.New("error type")
	}
	user :=  models.UserRolesDTO{
		//BaseModel: models.BaseModel{
		//	ID: int(userProto.Id.Id),
		//},
		//UserName:  user.UserName,
		//LoginName: user.LoginName,
		//Password:  user.Password,
		//Mobile:    int(user.Mobile),
		//Status:    int(user.Ststus),
		//GroupID:   int(user.GroupId),
		//UserType:  int(user.UserType),
		User : models.User{
			BaseModel : models.BaseModel{
				ID: int(userProto.Id.Id),
			},
			UserName: userProto.UserName,
			LoginName: userProto.LoginName,
			Password: userProto.Password,
			GroupID: int(userProto.GroupId),
			UserType: int(userProto.UserType),
			Mobile: int(userProto.Mobile),
			Status: int(userProto.Ststus),
		},
	}
	if userProto.Id != nil {
		user.ID = int(userProto.Id.Id)
	}
	if userProto.RoleIds != nil {
		for _, v := range userProto.RoleIds {
			user.RoleIDs = append(user.RoleIDs, int(v.Id))
		}
	}
	return user, nil
}

func EncodeAddUsersRequest(ctx context.Context, req interface{}) (interface{}, error) {
	if users, ok := req.(*pb_user_v1.AddUsersRequest); ok {
		return users, nil
	}
	return nil, fmt.Errorf(clientEncodeErr, reflect.TypeOf(req))
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

func EncodeUserCondition(ctx context.Context, i interface{}) (request interface{}, err error) {
	if userCondition, ok := i.(*pb_user_v1.UserQueryCondition); ok{
		return userCondition, nil
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

func DecodeUserResponse(ctx context.Context, i interface{}) (response interface{}, err error) {
	if userResponse, ok := i.(*pb_user_v1.UserQueryResponse); ok {
		return userResponse, nil
	}
	err= fmt.Errorf(clientEncodeErr, reflect.TypeOf(i))
	return
}

func EncodeUserIDs(ctx context.Context, i interface{}) (request interface{}, err error) {
	if userIds, ok := i.([]int64); ok {
		return &pb_user_v1.UserIDs{
			Ids: userIds,
		}, nil
	}
	err = fmt.Errorf(clientEncodeErr, reflect.TypeOf(i))
	return

}

func EncodeGroupAndUserId(ctx context.Context, i interface{}) (request interface{}, err error) {
	r, ok := i.(models.GroupAndUserId)
	if !ok {
		return nil, errors.New("error type")
	}
	groupUserId := &pb_user_v1.GroupUserIds{
		GroupId: int64(r.GroupId),
	}
	for _, v := range r.UserIds {
		groupUserId.UserIds = append(groupUserId.UserIds, int64(v))
	}
	return groupUserId,nil
}
