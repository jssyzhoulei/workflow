package parser

import (
	"context"
	"errors"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
)

func DecodeGroupAddProto(ctx context.Context, request interface{}) (interface{}, error) {
	r, ok := request.(*pb_user_v1.GroupAddRequest)
	if !ok {
		return nil, errors.New("解码 添加组 请求失败")
	}
	return r, nil
}

func EncodeGroupProto(ctx context.Context, request interface{}) (interface{}, error) {
	r, ok := request.(*pb_user_v1.GroupResponse)
	if !ok {
		return nil, errors.New("编码 添加组 结果失败")
	}
	return r, nil
}
