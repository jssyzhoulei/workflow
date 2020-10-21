package tranports

import (
	"context"
	"gitee.com/grandeep/org-svc/src/endpoints"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
	"gitee.com/grandeep/org-svc/src/tranports/parser"
	transport "github.com/go-kit/kit/transport/grpc"
)

type userGrpcTransport struct {
	addUser        transport.Handler
	getUserByID    transport.Handler
	updateUserByID transport.Handler
	deleteUserByID transport.Handler
	getUserList    transport.Handler
	batchDeleteUsers transport.Handler
	addUsers       transport.Handler
	importUsersByGroupId transport.Handler
}

func NewUserGrpcTransport(userEndpoint *endpoints.UserServiceEndpoint) *userGrpcTransport {
	var (
		addUserServer = transport.NewServer(userEndpoint.AddUserEndpoint, parser.DecodeUserProto, parser.EncodeNullProto)
		getUserByIDServer = transport.NewServer(userEndpoint.GetUserByIDEndpoint, parser.DecodeIndexProto, parser.EncodeUserProto)
		updateUserByIDServer = transport.NewServer(userEndpoint.UpdateUserByIDEndpoint, parser.DecodeUserProto, parser.EncodeNullProto)
		addUsersServer = transport.NewServer(userEndpoint.AddUsersEndpoint, parser.DecodeAddUsersRequest, parser.EncodeNullProto)
		deleteUserByIDServer = transport.NewServer(userEndpoint.DeleteUserByIDEndpoint, parser.DecodeIndexProto, parser.EncodeNullProto)
		getUserListServer = transport.NewServer(userEndpoint.GetUserListEndpoint, parser.DecodeUserPageProto, parser.EncodeUsersPage)
		batchDeleteUsersServer = transport.NewServer(userEndpoint.BatchDeleteUsersEndpoint, parser.DecodeUserIDsProto, parser.EncodeNullProto)
		importUsersByGroupIdServer = transport.NewServer(userEndpoint.ImportUsersByGroupIdEndpoint, parser.DecodeIndexProto, parser.EncodeNullProto)
	)
	return &userGrpcTransport{
		addUser:     addUserServer,
		getUserByID: getUserByIDServer,
		updateUserByID: updateUserByIDServer,
		deleteUserByID: deleteUserByIDServer,
		getUserList: getUserListServer,
		batchDeleteUsers: batchDeleteUsersServer,
		addUsers: addUsersServer,
		importUsersByGroupId: importUsersByGroupIdServer,
	}
}


func (u *userGrpcTransport) RpcAddUser(ctx context.Context, proto *pb_user_v1.UserProto) (*pb_user_v1.NullResponse, error) {
	_, resp, err := u.addUser.ServeGRPC(ctx, proto)
	if err != nil {
		return nil, err
	}
	return resp.(*pb_user_v1.NullResponse), nil
}

func (u *userGrpcTransport) RpcGetUserById(c context.Context, index *pb_user_v1.Index) (*pb_user_v1.UserProto, error) {
	_, resp, err := u.getUserByID.ServeGRPC(c, index)
	if err != nil {
		return nil, err
	}
	return resp.(*pb_user_v1.UserProto), nil
}

func (u *userGrpcTransport) RpcUpdateUserByID(ctx context.Context, proto *pb_user_v1.UserProto) (*pb_user_v1.NullResponse, error) {
	_, resp, err := u.updateUserByID.ServeGRPC(ctx, proto)
	if err != nil {
		return nil, err
	}
	return resp.(*pb_user_v1.NullResponse), nil
}

func (u *userGrpcTransport) RpcImportUsersByGroupId(ctx context.Context, index *pb_user_v1.Index) (*pb_user_v1.NullResponse, error) {
	_, resp, err := u.importUsersByGroupId.ServeGRPC(ctx, index)
	if err != nil {
		return nil, err
	}
	return resp.(*pb_user_v1.NullResponse), nil
}

func (u *userGrpcTransport) RpcDeleteUserByID(ctx context.Context, index *pb_user_v1.Index) (*pb_user_v1.NullResponse, error){
	_, resp, err := u.deleteUserByID.ServeGRPC(ctx, index)
	if err != nil {
		return nil, err
	}
	return resp.(*pb_user_v1.NullResponse), nil
}

//func (u *userGrpcTransport) RpcGetUserList(ctx context.Context, proto *pb_user_v1.UserProto) ()


func (u *userGrpcTransport) RpcAddUsers(ctx context.Context, request *pb_user_v1.AddUsersRequest) (*pb_user_v1.NullResponse, error) {
	_, resp, err := u.addUsers.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp.(*pb_user_v1.NullResponse), nil
}

func (u *userGrpcTransport) RpcGetUserList(c context.Context, proto *pb_user_v1.UserPage) (*pb_user_v1.UsersPage, error) {
	_, resp, err := u.getUserList.ServeGRPC(c, proto)
	if err != nil {
		return nil, err
	}
	return resp.(*pb_user_v1.UsersPage), nil
}

func (u *userGrpcTransport) RpcBatchDeleteUsers(c context.Context, ids *pb_user_v1.UserIDs) (*pb_user_v1.NullResponse, error) {
	_, resp, err := u.batchDeleteUsers.ServeGRPC(c, ids)
	if err != nil {
		return nil, err
	}
	return resp.(*pb_user_v1.NullResponse), nil
}