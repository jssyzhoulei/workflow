package parser

import (
	"context"
	"gitee.com/grandeep/org-svc/src/models"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
)

func DecodeRoleProto(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(*pb_user_v1.RoleProto)
	role := models.Role{
		Name:       r.Name,
		DataPermit: int(r.DataPermit),
		Status:     int(r.Status),
	}
	role.ID = int(r.Id)
	role.Remark = r.Remark
	return role, nil
}

func EncodeRoleProto(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(models.Role)
	return &pb_user_v1.RoleProto{
		Id:         int64(r.ID),
		Name:       r.Name,
		DataPermit: int32(r.DataPermit),
		Status:     int32(r.Status),
		Remark:     r.Remark,
	}, nil
}

func DecodeRoleMenuPermissionProto(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(*pb_user_v1.RoleMenuPermissionProto)
	return models.RoleMenuPermission{
		RoleID:       int(r.RoleId),
		MenuID:       int(r.MenuId),
		PermissionID: int(r.PermissionId),
	}, nil
}

func EncodeRoleMenuPermission(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(models.RoleMenuPermission)
	return &pb_user_v1.RoleMenuPermissionProto{
		RoleId:       int64(r.RoleID),
		MenuId:       int64(r.MenuID),
		PermissionId: int64(r.PermissionID),
	}, nil
}

func buildRoleMenuPermission(r *[]*pb_user_v1.RoleMenuPermissionProto, ctx context.Context) *[]*models.RoleMenuPermission {
	var resp []*models.RoleMenuPermission
	for i := range *r {
		decodeMP, _ := DecodeRoleMenuPermissionProto(ctx, (*r)[i])
		rmp := decodeMP.(models.RoleMenuPermission)
		resp = append(resp, &rmp)
	}
	return &resp
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

func DecodeCreateMenuPermRequestProto(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(*pb_user_v1.CreateMenuPermRequestProto)
	mpr := models.CreateMenuPermRequest{
		Role: models.Role{
			Name:       r.Name,
			DataPermit: int(r.DataPermit),
			Status:     int(r.Status),
		},
		MenuPerms: *buildRoleMenuPermission(&r.RoleMenuPermissions, ctx),
	}
	mpr.Remark = r.Remark
	return mpr, nil
}

func EncodeCreateMenuPermRequest(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(*models.CreateMenuPermRequest)
	return &pb_user_v1.CreateMenuPermRequestProto{
		Name:                r.Name,
		DataPermit:          int32(r.DataPermit),
		Status:              int32(r.Status),
		Remark:              r.Remark,
		RoleMenuPermissions: *buildRoleMenuPermissionProto(&r.MenuPerms, ctx),
	}, nil
}

func DecodeRolePageProto(ctx context.Context, req interface{}) (interface{}, error) {
	r := req.(*pb_user_v1.RolePageRequestProto)
	return r, nil
}

func EncodeRolePageProto(ctx context.Context, req interface{}) (interface{}, error) {
	return req.(*pb_user_v1.RolePageRequestProto), nil
}