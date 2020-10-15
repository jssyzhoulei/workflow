package parser

import (
	"context"
	"errors"
	"gitee.com/grandeep/org-svc/src/models"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
)

func EncodeRoleMenuPermission(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(models.RoleMenuPermission)
	return &pb_user_v1.RoleMenuPermissionProto{
		RoleId:       int64(r.RoleID),
		MenuId:       int64(r.MenuID),
		PermissionId: int64(r.PermissionID),
	}, nil
}

func buildRoleMenuPermissionProto(r *[]*models.RoleMenuPermission, ctx context.Context) *[]*pb_user_v1.RoleMenuPermissionProto {
	var resp []*pb_user_v1.RoleMenuPermissionProto
	for i := range *r {
		encodeData := (*r)[i]
		encodeMP, _ := EncodeRoleMenuPermission(ctx, *encodeData)
		rmp := encodeMP.(*pb_user_v1.RoleMenuPermissionProto)
		resp = append(resp, rmp)
	}
	return &resp
}

func EncodeCreateMenuPermRequestModel(ctx context.Context, req interface{}) (interface{}, error) {
	if role, ok := req.(models.CreateMenuPermRequest); ok {
		return &pb_user_v1.CreateMenuPermRequestProto{
			Name:                role.Name,
			Remark:              role.Remark,
			DataPermit:          int32(role.DataPermit),
			Status:              int32(role.Status),
			RoleMenuPermissions: *buildRoleMenuPermissionProto(&(role.MenuPerms), ctx),
		}, nil
	}
	return nil, errors.New("error type")
}
