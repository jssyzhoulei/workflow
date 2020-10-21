package parser

import (
	"context"
	"fmt"
	"gitee.com/grandeep/org-svc/src/models"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
	"reflect"
)

func DecodeUserProto(ctx context.Context, req interface{}) (interface{}, error) {

	r, ok := req.(*pb_user_v1.UserProto)
	if !ok {
		return models.User{}, fmt.Errorf(transportDecodeError, reflect.TypeOf(req))
	}
	userRoleDTO := models.UserRolesDTO{
		User: models.User{UserName: r.UserName,
		LoginName: r.LoginName,
		Password: r.Password,
		Mobile: int(r.Mobile),
		GroupID: int(r.GroupId),
		UserType: int(r.UserType),
		Status: int(r.Ststus),},
	}
	if r.Id != nil {
		userRoleDTO.ID = int(r.Id.Id)
	}
	if r.RoleIds != nil {
		for _, roleId := range r.RoleIds {
			if roleId != nil {
				userRoleDTO.RoleIDs = append(userRoleDTO.RoleIDs, int(roleId.Id))
			}
		}
	}
	return userRoleDTO, nil
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
		Ststus: int64(r.Status),
		GroupId: int64(r.GroupID),
		UserType: int64(r.UserType),
	}, nil
}

func DecodeAddUsersRequest(ctx context.Context, req interface{}) (interface{}, error) {
	if users, ok := req.(*pb_user_v1.AddUsersRequest); ok {
		return users, nil
	}
	return nil, fmt.Errorf(transportDecodeError, reflect.TypeOf(req))

}

func EncodeUsers(ctx context.Context, req interface{}) (response interface{}, err error) {
	if users, ok := req.(*pb_user_v1.Users); ok {
		return users, nil
	}
	err = fmt.Errorf(transportDecodeError, reflect.TypeOf(req))
	return
}

func DecodeUserPageProto(c context.Context, req interface{}) (interface{}, error) {
	if userPage,ok := req.(*pb_user_v1.UserPage); ok {
		return userPage, nil
	}
	fmt.Println(reflect.TypeOf(req))
	return nil, fmt.Errorf(transportDecodeError, reflect.TypeOf(req))
}

func EncodeUsersPage(c context.Context, res interface{}) (interface{}, error) {
	if usersPage, ok := res.(*pb_user_v1.UsersPage); ok {
		return usersPage, nil
	}
	return nil, fmt.Errorf(transportDecodeError, reflect.TypeOf(res))
}

func DecodeUserIDsProto(c context.Context, req interface{}) (interface{}, error) {
	if userIds,ok := req.(*pb_user_v1.UserIDs); ok {
		return userIds.Ids, nil
	}
	return nil, fmt.Errorf(transportDecodeError, reflect.TypeOf(req))
}

func DecodeGroupUserId(c context.Context, req interface{}) (interface{}, error) {
	r, ok := req.(*pb_user_v1.GroupUserIds)
	if !ok {
		return models.User{}, fmt.Errorf(transportDecodeError, reflect.TypeOf(req))
	}
	groupUserIds := models.GroupAndUserId{
		GroupId: int(r.GroupId),
	}

	if r.UserIds != nil {
		for _, userId := range r.UserIds {
			groupUserIds.UserIds = append(groupUserIds.UserIds, int(userId))
		}
	}
	return groupUserIds, nil
}

