package client

import (
	"context"
	org_endpoints "gitee.com/grandeep/org-svc/src/endpoints"
	"gitee.com/grandeep/org-svc/src/services"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/etcdv3"
	"github.com/go-kit/kit/sd/lb"
	"google.golang.org/grpc"
	"io"
	"time"
)

type OrgServiceClient struct {
	instance     *etcdv3.Instancer
	retryMax     int
	retryTimeout time.Duration
}

type GrpcUserConnFunc func(conn *grpc.ClientConn) services.UserServiceInterface

type MakeUserEndpointFunc func(userService services.UserServiceInterface) endpoint.Endpoint

type MakeGroupEndpointFunc func(userService services.GroupServiceInterface) endpoint.Endpoint

type GrpcGroupConnFunc func(conn *grpc.ClientConn) services.GroupServiceInterface

type GrpcRoleConnFunc func(conn *grpc.ClientConn) services.RoleServiceI

type MakeRoleEndpointFunc func(roleService services.RoleServiceI) endpoint.Endpoint

type MakePermissionEndpointFunc func(services.PermissionServiceInterface) endpoint.Endpoint

type GrpcPermissionConnFunc func(conn *grpc.ClientConn) services.PermissionServiceInterface

func NewOrgServiceClient(addr []string, retry int, timeOut time.Duration) *OrgServiceClient {
	var (
		etcdAddrs = addr
		serName   = "svc.org"
		ttl       = 5 * time.Second
	)
	options := etcdv3.ClientOptions{
		DialTimeout:   ttl,
		DialKeepAlive: ttl,
	}
	etcdClient, err := etcdv3.NewClient(context.Background(), etcdAddrs, options)
	if err != nil {
		return nil
	}
	instance, err := etcdv3.NewInstancer(etcdClient, serName, log.NewNopLogger())
	if err != nil {
		return nil
	}
	o := &OrgServiceClient{
		instance:     instance,
		retryMax:     retry,
		retryTimeout: timeOut,
	}

	return o
}

func (o *OrgServiceClient) GetUserService() services.UserServiceInterface {
	endpoints := &org_endpoints.UserServiceEndpoint{}
	endpoints.AddUserEndpoint = o.getRetryUserEndpoint(org_endpoints.MakeAddUserEndpoint, userGrpcConn)
	endpoints.GetUserByIDEndpoint = o.getRetryUserEndpoint(org_endpoints.MakeGetUserByIDEndpoint, userGrpcConn)
	endpoints.UpdateUserByIDEndpoint = o.getRetryUserEndpoint(org_endpoints.MakeUpdataUserByIDEndpoint, userGrpcConn)
	endpoints.DeleteUserByIDEndpoint = o.getRetryUserEndpoint(org_endpoints.MakeDeleteUserByIDEndpoint, userGrpcConn)
	endpoints.AddUsersEndpoint = o.getRetryUserEndpoint(org_endpoints.MakeAddUsersEndpoint, userGrpcConn)
	endpoints.GetUserListEndpoint = o.getRetryUserEndpoint(org_endpoints.MakeGetUserListEndpoint, userGrpcConn)
	endpoints.BatchDeleteUsersEndpoint = o.getRetryUserEndpoint(org_endpoints.MakeBatchDeleteUsersEndpoint, userGrpcConn)
	endpoints.ImportUsersByGroupIdEndpoint = o.getRetryUserEndpoint(org_endpoints.MakeImportUsersByGroupIdEndpoint, userGrpcConn)
	endpoints.GetUsersEndpoint = o.getRetryUserEndpoint(org_endpoints.MakeGetUsersEndpoint, userGrpcConn)
	return endpoints
}

func (o *OrgServiceClient) factoryForUser(makeEndpoint MakeUserEndpointFunc, conn GrpcUserConnFunc) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		con, err := grpc.Dial(instance, grpc.WithInsecure())
		if err != nil {
			return nil, nil, err
		}

		endpoints := makeEndpoint(conn(con))
		return endpoints, con, err
	}
}

func (o *OrgServiceClient) getRetryUserEndpoint(ept MakeUserEndpointFunc, conn GrpcUserConnFunc) endpoint.Endpoint {
	factory := o.factoryForUser(ept, conn)
	endpointer := sd.NewEndpointer(o.instance, factory, log.NewNopLogger())
	balance := lb.NewRandom(endpointer, time.Now().UnixNano())
	return lb.Retry(o.retryMax, o.retryTimeout, balance)
}

func (o *OrgServiceClient) GetRoleService() services.RoleServiceI {
	endpoints := &org_endpoints.RoleServiceEndpoint{}
	endpoints.AddRoleEndpoint = o.getRetryRoleEndpoint(org_endpoints.MakeAddRoleEndpoint, RoleGrpcConn)
	endpoints.UpdateRoleEndpoint = o.getRetryRoleEndpoint(org_endpoints.MakeUpdateRoleEndpoint, RoleGrpcConn)
	endpoints.DeleteRoleEndpoint = o.getRetryRoleEndpoint(org_endpoints.MakeDeleteRoleEndpoint, RoleGrpcConn)
	endpoints.QueryRoleEndpoint = o.getRetryRoleEndpoint(org_endpoints.MakeQueryRoleEndpoint, RoleGrpcConn)
	endpoints.QueryRolesEndpoint = o.getRetryRoleEndpoint(org_endpoints.MakeQueryRolesEndpoint, RoleGrpcConn)
	endpoints.MenuTreeEndpoint = o.getRetryRoleEndpoint(org_endpoints.MakeMenuTreeEndpoint, RoleGrpcConn)
	return endpoints
}

func (o *OrgServiceClient) factoryForRole(makeEndpoint MakeRoleEndpointFunc, conn GrpcRoleConnFunc) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		con, err := grpc.Dial(instance, grpc.WithInsecure())
		if err != nil {
			return nil, nil, err
		}

		endpoints := makeEndpoint(conn(con))
		return endpoints, con, err
	}
}

func (o *OrgServiceClient) getRetryRoleEndpoint(ept MakeRoleEndpointFunc, conn GrpcRoleConnFunc) endpoint.Endpoint {
	factory := o.factoryForRole(ept, conn)
	endpointer := sd.NewEndpointer(o.instance, factory, log.NewNopLogger())
	balance := lb.NewRandom(endpointer, time.Now().UnixNano())
	return lb.Retry(o.retryMax, o.retryTimeout, balance)
}

// #####################  Group #################################3

// GetGroupService  ...
func (o *OrgServiceClient) GetGroupService() services.GroupServiceInterface {
	return &org_endpoints.GroupServiceEndpoint{
		GroupAddEndpoint:                       o.getGroupRetryEndpoint(org_endpoints.MakeGroupAddEndpoint, groupGrpcConn),
		GroupQueryWithQuotaByConditionEndpoint: o.getGroupRetryEndpoint(org_endpoints.MakeGroupQueryWithQuotaByConditionEndpoint, groupGrpcConn),
		GroupUpdateEndpoint:                    o.getGroupRetryEndpoint(org_endpoints.MakeGroupUpdateEndpoint, groupGrpcConn),
		QuotaUpdateEndpoint:                    o.getGroupRetryEndpoint(org_endpoints.MakeQuotaUpdateEndpoint, groupGrpcConn),
		GroupTreeQueryEndpoint:                 o.getGroupRetryEndpoint(org_endpoints.MakeGroupTreeQueryEndpoint, groupGrpcConn),
		GroupDeleteEndpoint:                    o.getGroupRetryEndpoint(org_endpoints.MakeGroupDeleteEndpoint, groupGrpcConn),
		QueryGroupAndSubGroupsUsersEndpoint:    o.getGroupRetryEndpoint(org_endpoints.MakeQueryGroupAndSubGroupsUsersEndpoint, groupGrpcConn),
		SetGroupQuotaUsedEndpoint:              o.getGroupRetryEndpoint(org_endpoints.MakeSetGroupQuotaUsedEndpoint, groupGrpcConn),
		QueryGroupIDAndSubGroupsIDEndpoint:     o.getGroupRetryEndpoint(org_endpoints.MakeQueryGroupIDAndSubGroupsIDEndpoint, groupGrpcConn),
	}
}

func (o *OrgServiceClient) getGroupRetryEndpoint(ept MakeGroupEndpointFunc, conn GrpcGroupConnFunc) endpoint.Endpoint {
	factory := o.factoryGroupFor(ept, conn)
	endpointer := sd.NewEndpointer(o.instance, factory, log.NewNopLogger())
	balance := lb.NewRandom(endpointer, time.Now().UnixNano())
	return lb.Retry(o.retryMax, o.retryTimeout, balance)
}

func (o *OrgServiceClient) factoryGroupFor(makeEndpoint MakeGroupEndpointFunc, conn GrpcGroupConnFunc) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		con, err := grpc.Dial(instance, grpc.WithInsecure())
		if err != nil {
			return nil, nil, err
		}

		endpoints := makeEndpoint(conn(con))
		return endpoints, con, err
	}
}

func (o *OrgServiceClient) GetPermissionService() services.PermissionServiceInterface {
	endpoints := &org_endpoints.PermissionServiceEndpoint{}
	endpoints.AddPermissionEndpoint = o.getPermissionRetryEndpoint(org_endpoints.MakeAddPermissionEndpoint, permissionGrpcConn)
	endpoints.AddMenuEndpoint = o.getPermissionRetryEndpoint(org_endpoints.MakeAddMenuEndpoint, permissionGrpcConn)
	endpoints.GetMenuCascadeByModuleEndpoint = o.getPermissionRetryEndpoint(org_endpoints.MakeGetMenuCascadeByModuleEndpoint, permissionGrpcConn)
	endpoints.GetPermissionByIDEndpoint = o.getPermissionRetryEndpoint(org_endpoints.MakeGetPermissionByID, permissionGrpcConn)
	endpoints.DeletePermissionByIDEndpoint = o.getPermissionRetryEndpoint(org_endpoints.MakeDeletePermissionByID, permissionGrpcConn)
	endpoints.UpdatePermissionByIDEndpoint = o.getPermissionRetryEndpoint(org_endpoints.MakeUpdatePermissionByIDEndpoint, permissionGrpcConn)
	return endpoints
}

func (o *OrgServiceClient) getPermissionRetryEndpoint(ept MakePermissionEndpointFunc, conn GrpcPermissionConnFunc) endpoint.Endpoint {
	factory := o.factoryPermissionFor(ept, conn)
	endpointer := sd.NewEndpointer(o.instance, factory, log.NewNopLogger())
	balance := lb.NewRandom(endpointer, time.Now().UnixNano())
	return lb.Retry(o.retryMax, o.retryTimeout, balance)
}

func (o *OrgServiceClient) factoryPermissionFor(makeEndpoint MakePermissionEndpointFunc, conn GrpcPermissionConnFunc) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		con, err := grpc.Dial(instance, grpc.WithInsecure())
		if err != nil {
			return nil, nil, err
		}

		endpoints := makeEndpoint(conn(con))
		return endpoints, con, err
	}
}
