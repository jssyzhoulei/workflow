package parser

import (
	"context"
	"fmt"
	"gitee.com/grandeep/org-svc/src/models"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
	"reflect"
)

var (
	clientEncodeErr = "client encode err type is %v"
)

func EncodePermissionModel(ctx context.Context, req interface{}) (interface{}, error) {
	if permission, ok := req.(models.Permission); ok {
		return &pb_user_v1.PermissionProto{
			Index:      &pb_user_v1.Index{
				Id:                   int64(permission.ID),
			},
			UserId:     0,
			UriName:    permission.UriName,
			Method:     int32(permission.Method),
			Uri:        permission.Uri,
			Relation:   int32(permission.Relation),
			ButtonName: permission.ButtonName,
			ButtonKey:  permission.ButtonKey,
			MenuId:     int64(permission.MenuID),
		}, nil
	}
	return nil, fmt.Errorf(clientEncodeErr, reflect.TypeOf(req))
}

func EncodeMenuModel(_ context.Context, req interface{}) (interface{}, error) {
	if menu, ok := req.(models.Menu); ok {
		return &pb_user_v1.MenuProto{
			Index:        nil,
			UserId:       0,
			Name:         menu.Name,
			ParenId:      int64(menu.ParentID),
			Module:       pb_user_v1.Module(menu.Module),
			Order:        int64(menu.Order),
			TemplatePath: menu.TemplatePath,
		}, nil
	}
	return nil, fmt.Errorf(clientEncodeErr, reflect.TypeOf(req))
}

func EncodeMenuModule(_ context.Context, i interface{}) (interface{}, error) {
	if module, ok := i.(models.MenuModule); ok {
		return &pb_user_v1.MenuModule{
			Module: pb_user_v1.Module(module),
		}, nil
	}

	return nil, fmt.Errorf(clientEncodeErr, reflect.TypeOf(i))
}

func DecodeCascadeProto(_ context.Context, i interface{}) (interface{}, error) {
	if cascades, ok := i.(*pb_user_v1.Cascades); ok {
		return cascades, nil
	}
	err := fmt.Errorf(clientEncodeErr, reflect.TypeOf(i))
	return nil, err
}

func DecodePermissionProto(_ context.Context, i interface{}) (interface{}, error) {
	if permission, ok := i.(*pb_user_v1.PermissionProto); ok {
		pm := models.Permission{
			UriName:    permission.UriName,
			Method:     models.RequestMethod(permission.Method),
			Uri:        permission.Uri,
			Relation:   int(permission.Relation),
			ButtonName: permission.ButtonName,
			ButtonKey:  permission.ButtonKey,
			MenuID:     int(permission.MenuId),
			Module:     models.MenuModule(permission.Module),
		}
		if permission.Index != nil {
			pm.ID = int(permission.Index.Id)
		}
		return pm, nil
	}
	return nil, fmt.Errorf(clientEncodeErr, reflect.TypeOf(i))
}
