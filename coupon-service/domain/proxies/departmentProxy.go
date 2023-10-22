package proxies

import (
	"context"
	"fmt"
	"log"

	"github.com/MaheshBailwal/mscore/core"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	department_service "invafresh.com/coupon-service/infrastructure/remoteServices/department"
)

type DepartmentProxy struct {
	grpcPort string
}

func NewDepartmentProxy(portNumber string) *DepartmentProxy {
	return &DepartmentProxy{grpcPort: portNumber}
}

func (d *DepartmentProxy) GetDepartmets(sc core.ServiceContext) *department_service.DepartmentDetails {
	conn := d.makeConnection()
	defer conn.Close()

	c := department_service.NewDepartmentServiceClient(conn)

	header := metadata.New(map[string]string{"user_id": sc.CurrentUserId, "space": "", "org": "", "limit": "", "offset": ""})

	ctx := metadata.NewOutgoingContext(context.Background(), header)
	response, err := c.GetAll(ctx, &department_service.EmptyParams_Dep{})
	if err != nil {

		//log.Fatalf("Error when calling SayHello: %s", err)

		st := status.Convert(err)
		for _, detail := range st.Details() {
			switch t := detail.(type) {
			case *errdetails.ErrorInfo:
				fmt.Println("Oops! Your request was rejected by the server.")
				fmt.Println("Reson Code", t.GetReason())
				fmt.Println("Meatadata", t.GetMetadata())

			}
		}
	}

	return response
}

func (d *DepartmentProxy) makeConnection() *grpc.ClientConn {
	var conn *grpc.ClientConn
	port := fmt.Sprintf("%s", d.grpcPort)
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	return conn
}
