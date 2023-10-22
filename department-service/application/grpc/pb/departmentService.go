package pb

import (
	"fmt"
	"log"

	"github.com/MaheshBailwal/mscore/core"
	"github.com/MaheshBailwal/mscore/persistence"
	"github.com/jinzhu/copier"
	"golang.org/x/net/context"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"invafresh.com/department-service/application/factories"
	"invafresh.com/department-service/domain/interfaces"
	"invafresh.com/department-service/domain/services"
	"invafresh.com/department-service/infrastructure/config"
)

type DepartmentService struct {
	UnimplementedDepartmentServiceServer
	sconfig config.SConfig
}

func NewDepartmentService(config config.SConfig) DepartmentService {
	return DepartmentService{
		sconfig: config,
	}
}

func (c *DepartmentService) GetAll(ctx context.Context, request *EmptyParams_Dep) (*DepartmentDetails, error) {
	log.Println("Received message GetAll Coupon")

	departmentService := services.DepartmentService{
		UnitOfWork: c.CreateUnitOfWork(),
	}

	filters := persistence.QueryFilter{}
	departmentDetails, err := departmentService.GetAll(createServiceContext(ctx), filters)
	if err != nil {
		return nil, err
	}

	var departmentsDetail []*DepartmentDetail

	for _, entity := range departmentDetails {
		couponProfileDetail := DepartmentDetail{}
		copier.Copy(&couponProfileDetail, &entity)
		departmentsDetail = append(departmentsDetail, &couponProfileDetail)
	}

	return &DepartmentDetails{
		Department: departmentsDetail,
	}, nil

}

func createServiceContext(ctx context.Context) core.ServiceContext {
	return factories.CreateServiceContextGrpc(ctx)
}

func (r *DepartmentService) CreateUnitOfWork() interfaces.IUnitOfWork {
	return factories.CreateUnitOfWork(r.sconfig)
}

func setError() error {
	metadata := make(map[string]string)
	metadata["userId"] = "1"

	st := status.New(codes.InvalidArgument, "invalid username")
	v := &errdetails.ErrorInfo{
		Reason:   "1003",
		Metadata: metadata,
	}
	// br := &errdetails.BadRequest{}
	// br.FieldViolations = append(br.FieldViolations, v)
	st, err := st.WithDetails(v)
	if err != nil {
		// If this errored, it will always error
		// here, so better panic so we can figure
		// out why than have this silently passing.
		panic(fmt.Sprintf("Unexpected error attaching metadata: %v", err))
	}
	return st.Err()
}
