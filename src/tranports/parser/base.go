package parser

import (
	"context"
	"fmt"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
	"reflect"
)

func EncodeNullProto(_ context.Context, res interface{}) (interface{}, error) {
	response := res.(pb_user_v1.NullResponse)
	return &response, nil
}

func DecodeIndexProto(ctx context.Context, i interface{}) (request interface{}, err error) {
	if index, ok := i.(*pb_user_v1.Index); ok {
		return int(index.Id), nil
	}
	return nil, fmt.Errorf(transportDecodeError, reflect.TypeOf(i))
}
