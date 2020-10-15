package parser

import (
	"context"
	"errors"
	"fmt"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
	"reflect"
)

func DecodeNullProto(ctx context.Context, res interface{}) (interface{}, error) {
	fmt.Println(reflect.TypeOf(res))
	if resp, ok := res.(*pb_user_v1.NullResponse);ok {
		return *resp, nil
	}
	return nil, errors.New("type error")
}
