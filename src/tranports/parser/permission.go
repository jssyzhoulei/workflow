package parser

import (
	"context"
	"errors"
	"gitee.com/grandeep/org-svc/src/models"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
)

func EncodePermissionProto(ctx context.Context, i interface{}) (response interface{}, err error) {
	return
}

func DecodePermissionProto(ctx context.Context, i interface{}) (request interface{}, err error) {
	per,ok := i.(*pb_user_v1.PermissionProto)
	if !ok {
		return models.Permission{}, errors.New("error type")
	}
	return models.Permission{
		UriName:    per.UriName,
		Method:     models.RequestMethod(per.Method),
		Uri:       	per.Uri,
		Relation:   int(per.Relation),
		ButtonName: per.ButtonName,
		ButtonKey:  per.ButtonKey,
		MenuID:     int(per.MenuId),
	}, nil
}
