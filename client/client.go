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

type GrpcUserConnFunc func(conn *grpc.ClientConn) services.UserServiceI

type MakeUserEndpointFunc func(userService services.UserServiceI) endpoint.Endpoint

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
		instance: instance,
		retryMax: retry,
		retryTimeout: timeOut,
	}

	return o
}

func (o *OrgServiceClient) GetUserService() services.UserServiceI {
	endpoints := &org_endpoints.UserServiceEndpoint{}
	endpoints.AddUserEndpoint = o.getRetryEndpoint(org_endpoints.MakeAddUserEndpoint, addUserGrpcConn)
	return endpoints
}

func (o *OrgServiceClient) factoryFor(makeEndpoint MakeUserEndpointFunc, conn GrpcUserConnFunc) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		con, err := grpc.Dial(instance, grpc.WithInsecure())
		if err != nil {
			return nil, nil, err
		}

		endpoints := makeEndpoint(conn(con))
		return endpoints, con, err
	}
}

func (o *OrgServiceClient) getRetryEndpoint(ept MakeUserEndpointFunc, conn GrpcUserConnFunc) endpoint.Endpoint {
	factory := o.factoryFor(ept, conn)
	endpointer := sd.NewEndpointer(o.instance, factory, log.NewNopLogger())
	balance := lb.NewRandom(endpointer, time.Now().UnixNano())
	return lb.Retry(o.retryMax, o.retryTimeout, balance)
}