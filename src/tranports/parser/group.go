/*
Package parser
这里和 client/parser 中的解析参数函数相反
这里是服务端端的角度进行编解码
Decode 发送过来的数据 -> proto message
Encode proto message -> 返回给客户端的消息
*/

package parser

import (
	"context"
	"errors"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
)

// DecodeGroupAddProto ...
func DecodeGroupAddProto(ctx context.Context, request interface{}) (interface{}, error) {
	r, ok := request.(*pb_user_v1.GroupAddRequest)
	if !ok {
		return nil, errors.New("解码 添加组 请求失败")
	}
	return r, nil
}
// EncodeGroupProto ...
func EncodeGroupProto(ctx context.Context, request interface{}) (interface{}, error) {
	r, ok := request.(*pb_user_v1.GroupResponse)
	if !ok {
		return nil, errors.New("编码 添加组 结果失败")
	}
	return r, nil
}

// DecodeGroupQueryByConditionProto ...
func DecodeGroupQueryByConditionProto(ctx context.Context, request interface{}) (interface{}, error) {
	r, ok := request.(*pb_user_v1.GroupQueryWithQuotaByConditionRequest)
	if !ok {
		return nil, errors.New("transport DecodeGroupQueryByConditionProto 失败")
	}
	return r, nil
}

// EncodeGroupQueryByConditionProto ...
func EncodeGroupQueryByConditionProto(ctx context.Context, response interface{}) (interface{}, error) {
	r, ok := response.(*pb_user_v1.GroupQueryWithQuotaByConditionResponse)
	if !ok {
		return nil, errors.New("transport EncodeGroupQueryByConditionProto 失败")
	}
	return r, nil
}

// DecodeGroupUpdateProto ...
func DecodeGroupUpdateProto(ctx context.Context, request interface{}) (interface{}, error) {
	r, ok := request.(*pb_user_v1.GroupUpdateRequest)
	if !ok {
		return nil, errors.New("解码 添加组 请求失败")
	}
	return r, nil
}
