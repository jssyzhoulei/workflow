package parser

import (
	"context"
	"fmt"
	"gitee.com/grandeep/org-svc/src/models"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
	"reflect"
)

var (
	transportDecodeError = "transport decode error, type is %v"
)

func DecodePermissionProto(ctx context.Context, i interface{}) (request interface{}, err error) {
	per,ok := i.(*pb_user_v1.PermissionProto)
	if !ok {
		return models.Permission{}, fmt.Errorf(transportDecodeError, reflect.TypeOf(i))
	}
	permission := models.Permission{
		UriName:    per.UriName,
		Method:     models.RequestMethod(per.Method),
		Uri:       	per.Uri,
		Relation:   int(per.Relation),
		ButtonName: per.ButtonName,
		ButtonKey:  per.ButtonKey,
		MenuID:     int(per.MenuId),
	}
	if per.Index != nil {
		permission.ID = int(per.Index.Id)
	}
	return permission, nil
}


func DecodeMenuProto(ctx context.Context, i interface{}) (request interface{}, err error) {
	menu, ok := i.(*pb_user_v1.MenuProto)
	if !ok {
		return models.Menu{}, fmt.Errorf(transportDecodeError, reflect.TypeOf(i))
	}
	return models.Menu{
		BaseModel:models.BaseModel{
			CreateUserID:  int(menu.UserId),
		},
		Name:         menu.Name,
		ParentID:     int(menu.ParenId),
		Module:       models.MenuModule(menu.Module),
		Order:        int(menu.Order),
		TemplatePath: menu.TemplatePath,
	}, nil
}

func DecodeModuleProto(ctx context.Context, i interface{}) (request interface{}, err error) {
	module, ok := i.(*pb_user_v1.MenuModule)
	if !ok {
		return nil, fmt.Errorf(transportDecodeError, reflect.TypeOf(i))
	}
	return models.MenuModule(module.Module), nil
}

func EncodeCascadeProto(ctx context.Context, i interface{}) (response interface{}, err error) {
	if cascades, ok := i.(*pb_user_v1.Cascades); ok {
		return cascades, nil
	}
	err = fmt.Errorf(transportDecodeError, reflect.TypeOf(i))
	return
}

func EncodePermissionProto(ctx context.Context, i interface{}) (response interface{}, err error) {
	if permission, ok := i.(models.Permission); ok {
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
			Module: int32(permission.Module),
		}, nil
	}
	return nil, fmt.Errorf(transportDecodeError, reflect.TypeOf(i))
}
