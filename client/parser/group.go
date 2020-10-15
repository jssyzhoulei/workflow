/*
Package parser ...
这里和 src/transports/parser 中的解析参数函数相反
这里是客户端的角度进行编解码
Encode 发送的数据 -> proto message
Decode proto message -> 返回的消息
*/
package parser

import (
	"context"
	"errors"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
)

// EncodeGroupAddProto ...
func EncodeGroupAddProto(ctx context.Context, request interface{}) (interface{}, error) {
	r, ok := request.(*pb_user_v1.GroupAddRequest)
	if !ok {
		return nil, errors.New("解码 添加组 请求失败")
	}
	return r, nil
}

// DecodeGroupProto ...
func DecodeGroupProto(ctx context.Context, request interface{}) (interface{}, error) {
	r, ok := request.(*pb_user_v1.GroupResponse)
	if !ok {
		return nil, errors.New("编码 添加组 结果失败")
	}
	return r, nil
}

// EncodeGroupQueryWithQuotaByConditionProto ...
func EncodeGroupQueryWithQuotaByConditionProto(ctx context.Context, request interface{}) (interface{}, error) {
	r, ok := request.(*pb_user_v1.GroupQueryWithQuotaByConditionRequest)
	if !ok {
		return nil, errors.New("EncodeGroupQueryWithQuotaByConditionProto 失败")
	}
	return r, nil
}

// DecodeGroupQueryWithQuotaByConditionProto ...
func DecodeGroupQueryWithQuotaByConditionProto(ctx context.Context, response interface{}) (interface{}, error) {
	r, ok := response.(*pb_user_v1.GroupQueryWithQuotaByConditionResponse)
	if !ok {
		return nil, errors.New("DecodeGroupQueryWithQuotaByConditionProto 失败")
	}
	return r, nil
}