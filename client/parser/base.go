package parser

import (
	"context"
	"fmt"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
	"reflect"
)

func DecodeNullProto(ctx context.Context, res interface{}) (interface{}, error) {
	if resp, ok := res.(*pb_user_v1.NullResponse); ok {
		return *resp, nil
	}
	return nil, fmt.Errorf(clientEncodeErr, reflect.TypeOf(res))
}

func EncodeIndexProto(_ context.Context, request interface{}) (interface{}, error) {
	if index, ok := request.(int); ok {
		return &pb_user_v1.Index{
			Id: int64(index),
		}, nil
	}
	if index, ok := request.(int64); ok {
		return &pb_user_v1.Index{
			Id: int64(index),
		}, nil
	}
	if index, ok := request.(int8); ok {
		return &pb_user_v1.Index{
			Id: int64(index),
		}, nil
	}
	if index, ok := request.(int32); ok {
		return &pb_user_v1.Index{
			Id: int64(index),
		}, nil
	}
	return nil, fmt.Errorf(clientEncodeErr, reflect.TypeOf(request))
}
