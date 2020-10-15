package parser

import (
	"context"
	"errors"
	"gitee.com/grandeep/org-svc/src/models"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
)

func EncodePermissionModel(ctx context.Context, req interface{}) (interface{}, error) {
	if permission, ok := req.(models.Permission); ok {
		return &pb_user_v1.PermissionProto{
			Index:                nil,
			UserId:               0,
			UriName:              permission.UriName,
			Method:               int32(permission.Method),
			Uri:                  permission.Uri,
			Relation:             int32(permission.Relation),
			ButtonName:           permission.ButtonName,
			ButtonKey:            permission.ButtonKey,
			MenuId:               int64(permission.MenuID),
		},nil
	}
	return nil, errors.New("error type")
}
