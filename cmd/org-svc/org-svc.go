package main

import (
	"context"
	"flag"
	"fmt"
	"gitee.com/grandeep/org-svc/src/endpoints"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
	"gitee.com/grandeep/org-svc/src/repositories"
	"gitee.com/grandeep/org-svc/src/services"
	"gitee.com/grandeep/org-svc/src/tranports"
	"gitee.com/grandeep/org-svc/utils/src/pkg/engine"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd/etcdv3"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var (
	ip = flag.String("ip", "", "")
)

func main() {
	flag.Parse()
	var (
		serName   = "svc.org"
		ttl       = 5 * time.Second
		quitChan = make(chan error, 1)
		baseServer *grpc.Server
		grpcAddr = ":866"
	)

	e := engine.NewEngine("./resources/config/config.yaml")
	options := etcdv3.ClientOptions{
		DialTimeout:   ttl,
		DialKeepAlive: ttl,
	}
	etcdHost, _ := e.Config.GetString("etcdHost")
	etcdClient, err := etcdv3.NewClient(context.Background(), strings.Split(etcdHost, ";") , options)
	if err != nil {
		e.Logger.Error("etcd host error")
		return
	}
	Register := etcdv3.NewRegistrar(etcdClient, etcdv3.Service{
		Key:   fmt.Sprintf("%s/%s",serName,grpcAddr),
		Value: grpcAddr,
	}, log.NewNopLogger())
	repos := repositories.NewRepoI(e.DB)
	svc := services.NewService(repos, e)
	ept := endpoints.NewEndpoint(svc)
	tpt := tranports.NewTransport(ept)
	go func() {
		grpcListener, err := net.Listen("tcp", grpcAddr)
		if err != nil {
			return
		}
		Register.Register()

		baseServer = grpc.NewServer(grpc.UnaryInterceptor(Intercept))

		e.Logger.Info(fmt.Sprintf("Listening and serving HTTP on %s", grpcAddr))
		pb_user_v1.RegisterRpcOrgServiceServer(baseServer, tpt)
		if err = baseServer.Serve(grpcListener); err != nil {
			quitChan <- err
			return
		}
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		quitChan <- fmt.Errorf("%s", <-c)
	}()
	<- quitChan
	Register.Deregister()
	e.Logger.Info("bye")
	//services.NewService()
}

func Intercept(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	var (
		fn context.CancelFunc
	)
	ctx = context.WithValue(ctx, "grpcContext", info.FullMethod)
	ctx = context.WithValue(ctx, "startTime", time.Now())
	ctx, fn = context.WithTimeout(ctx, time.Second)
	defer fn()
	return handler(ctx, req)
}