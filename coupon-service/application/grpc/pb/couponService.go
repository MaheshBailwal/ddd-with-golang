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
	"invafresh.com/coupon-service/application/factories"
	"invafresh.com/coupon-service/domain/interfaces"
	"invafresh.com/coupon-service/domain/services"
	"invafresh.com/coupon-service/infrastructure/config"
)

type CouponService struct {
	UnimplementedCouponServiceServer
	sconfig config.SConfig
}

func NewCouponService(config config.SConfig) CouponService {
	return CouponService{
		sconfig: config,
	}
}

func (c *CouponService) GetAll(ctx context.Context, request *EmptyParams) (*CouponProfileDetails, error) {
	log.Println("Received message GetAll Coupon")

	coupanService := services.CouponService{
		UnitOfWork: c.CreateUnitOfWork(),
	}

	filters := persistence.QueryFilter{}
	couponsDetails, err := coupanService.GetAll(createServiceContext(ctx), filters)
	if err != nil {
		return nil, err
	}

	var couponProfilesDetail []*CouponProfileDetail
	var couponProfileDetail CouponProfileDetail

	for _, entity := range couponsDetails {
		copier.Copy(&couponProfileDetail, &entity)
		couponProfilesDetail = append(couponProfilesDetail, &couponProfileDetail)
	}

	return &CouponProfileDetails{
		CouponProfile: couponProfilesDetail,
	}, nil

}

func createServiceContext(ctx context.Context) core.ServiceContext {
	return factories.CreateServiceContextGrpc(ctx)
}

func (r *CouponService) CreateUnitOfWork() interfaces.IUnitOfWork {
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
