package gapi

import (
	"fmt"
	"log"
	"net"

	"invafresh.com/coupon-service/application/grpc/pb"
	"invafresh.com/coupon-service/infrastructure/config"
	"google.golang.org/grpc"

	"github.com/ehsandavari/go-logger"
)

type GApi struct {
	sConfig *config.SConfig
	iLogger logger.ILogger
}

func NewGApi(sConfig *config.SConfig, iLogger logger.ILogger) *GApi {
	var gApi GApi
	gApi.sConfig = sConfig
	return &gApi
}

func (r *GApi) Start() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", r.sConfig.Service.Grpc.Port))
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", r.sConfig.Service.Grpc.Port, err)
	}

	s := pb.NewCouponService(*r.sConfig)

	grpcServer := grpc.NewServer()

	pb.RegisterCouponServiceServer(grpcServer, &s)

	fmt.Println("gRPC server running over port ", r.sConfig.Service.Grpc.Port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port %s: %v", r.sConfig.Service.Grpc.Port, err)
	}
}

func (r *GApi) Stop() {
	if r.sConfig.Service.Api.IsEnabled {
		// err := r.server.Shutdown(context.Background())
		// if err != nil {
		// 	r.iLogger.Error("error in shutdown api server", err)
		// }
	}
}
