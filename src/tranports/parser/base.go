package parser

import (
	"context"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
)

func EncodeNullProto(_ context.Context, res interface{}) (interface{}, error) {
	response := res.(pb_user_v1.NullResponse)
	return &response, nil
}
