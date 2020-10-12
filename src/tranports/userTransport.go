package tranports

import (
	"context"
	"gitee.com/grandeep/org-svc/src/endpoints"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
	"gitee.com/grandeep/org-svc/src/tranports/parser"
	transport "github.com/go-kit/kit/transport/grpc"
)

type userGrpcTransport struct {
	addUser     transport.Handler
	getUserById transport.Handler
}

func NewUserGrpcTransport(endpoint *endpoints.UserServiceEndpoint) *userGrpcTransport {
	var (
		addUserServer = transport.NewServer(endpoint.AddUserEndpoint, parser.DecodeUserProto, parser.EncodeUserProto)
	)
	return &userGrpcTransport{
		addUser:     addUserServer,
	}
}


func (u *userGrpcTransport) RpcAddUser(ctx context.Context, proto *pb_user_v1.UserProto) (*pb_user_v1.UserProto, error) {
	_, user, err := u.addUser.ServeGRPC(ctx, proto)
	if err != nil {
		return nil, err
	}
	return user.(*pb_user_v1.UserProto), err
}

func (u *userGrpcTransport) RpcGetUserById(ctx context.Context, proto *pb_user_v1.UserProto) (*pb_user_v1.UserProto, error) {
	panic("implement me")
}