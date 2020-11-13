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
		return nil, errors.New("DecodeGroupUpdateProto 失败")
	}
	return r, nil
}

// DecodeQuotaUpdateProto ...
func DecodeQuotaUpdateProto(ctx context.Context, request interface{}) (interface{}, error) {
	r, ok := request.(*pb_user_v1.QuotaUpdateRequest)
	if !ok {
		return nil, errors.New("DecodeQuotaUpdateProto 失败")
	}
	return r, nil
}

// EncodeGroupTreeQueryProto ...
func EncodeGroupTreeQueryProto(ctx context.Context, request interface{}) (interface{}, error) {
	r, ok := request.(*pb_user_v1.GroupTreeResponse)
	if !ok {
		return nil, errors.New("EncodeGroupTreeQueryProto 失败")
	}
	return r, nil
}

// DecodeGroupIDProto ...
func DecodeGroupIDProto(ctx context.Context, request interface{}) (interface{}, error) {
	r, ok := request.(*pb_user_v1.GroupID)
	if !ok {
		return nil, errors.New("DecodeGroupIDProto 失败")
	}
	return r, nil
}

// DecodeSetGroupQuotaUsedProto ...
func DecodeSetGroupQuotaUsedProto(ctx context.Context, request interface{}) (interface{}, error) {
	r, ok := request.(*pb_user_v1.SetGroupQuotaUsedRequest)
	if !ok {
		return nil, errors.New("DecodeSetGroupQuotaUsedProto 失败")
	}
	return r, nil
}

// EncodeGroupIDsResponse ...
func EncodeGroupIDsResponse(ctx context.Context, request interface{}) (interface{}, error) {
	r, ok := request.(*pb_user_v1.GroupIDsResponse)
	if !ok {
		return nil, errors.New("EncodeGroupIDsResponse 失败")
	}
	return r, nil
}

// DecodeQueryQuotaByCondition ...
func DecodeQueryQuotaByCondition(ctx context.Context, request interface{}) (interface{}, error) {
	r, ok := request.(*pb_user_v1.QueryQuotaByCondition)
	if !ok {
		return nil, errors.New("DecodeQueryQuotaByCondition 失败")
	}
	return r, nil
}

// EncodeQueryQuotaByConditionResponse ...
func EncodeQueryQuotaByConditionResponse(ctx context.Context, request interface{}) (interface{}, error) {
	r, ok := request.(*pb_user_v1.QueryQuotaByConditionResponse)
	if !ok {
		return nil, errors.New("EncodeQueryQuotaByConditionResponse 失败")
	}
	return r, nil
}

func EncodeGroupsProto(_ context.Context, res interface{}) (interface{}, error) {
	if _,ok := res.(*pb_user_v1.Groups);ok {
		return res, nil
	}
	return nil, errors.New("EncodeGroupsProto 失败")
}

// EncodeQueryQuotaResponse ...
func EncodeQueryQuotaResponse(_ context.Context, res interface{}) (interface{}, error) {
	if _,ok := res.(*pb_user_v1.QueryQuotaResponse);ok {
		return res, nil
	}
	return nil, errors.New("EncodeGroupQuotaResponse 失败")
}

// DecodeGroupIDWithPage ...
func DecodeGroupIDWithPage(_ context.Context, res interface{}) (interface{}, error) {
	if _,ok := res.(*pb_user_v1.GroupIDWithPage);ok {
		return res, nil
	}
	return nil, errors.New("DecodeGroupIDWithPage 失败")
}


// EncodeGroupUsersWithPage ...
func EncodeGroupUsersWithPage(_ context.Context, res interface{}) (interface{}, error) {
	if _,ok := res.(*pb_user_v1.GroupUsersWithPage);ok {
		return res, nil
	}
	return nil, errors.New("EncodeGroupUsersWithPage 失败")
}

