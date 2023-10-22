package controllers

import (
	"fmt"

	"github.com/MaheshBailwal/mscore/api"
	"github.com/ehsandavari/go-logger"
	"github.com/gin-gonic/gin"
	emsgp "invafresh.com/coupon-service/application/api/errorMessageProvider"
	"invafresh.com/coupon-service/application/factories"
	"invafresh.com/coupon-service/domain/dtos"
	"invafresh.com/coupon-service/domain/proxies"
	"invafresh.com/coupon-service/domain/services"
	"invafresh.com/coupon-service/infrastructure/config"
)

type CouponController struct {
	base    api.BaseController
	sconfig config.SConfig
	logger  logger.ILogger
}

func NewCouponController(routerGroup *gin.RouterGroup, logger logger.ILogger, sconfig config.SConfig) {
	couponController := CouponController{
		base:    api.NewBaseController(),
		sconfig: sconfig,
		logger:  logger,
	}

	errMsgProvider := emsgp.NewErrorMessageProvider()

	routerGroup = routerGroup.Group("coupon")
	{
		routerGroup.PUT("", api.RequestHandler[*dtos.CouponProfileUpdate, int64](couponController.update).Handle(couponController.logger, errMsgProvider))
		routerGroup.GET("", api.RequestHandler[dtos.EmptyDto, []dtos.CouponProfileDetail](couponController.getall).Handle(couponController.logger, errMsgProvider))
		routerGroup.DELETE("", api.RequestHandler[[]dtos.CouponProfileId, bool](couponController.delete).Handle(couponController.logger, errMsgProvider))
	}
}

// Update coupon
//
//	@Tags		coupon
//	@Summary	insert/update coupon
//	@Accept		json
//	@Produce	json
//	@Param		params			body		CouponProfileUpdate				false	"Query Params"
//	@Success	200				{object}	ApiResponse[int]
//	@Failure	400				{object}	ApiResponse[ApiError]
//	@Failure	500				{object}	ApiResponse[ApiError]
//	@Router		/coupon [PUT]
//
// @Security Bearer
func (c *CouponController) update(ctx *gin.Context, dto *dtos.CouponProfileUpdate) (int64, error) {
	coupanService := services.CouponService{
		UnitOfWork: factories.CreateUnitOfWork(c.sconfig),
	}

	entityId, err := coupanService.Update(factories.CreateServiceContext(ctx), dto)
	if err != nil {
		return -1, err
	}
	coupanService.UnitOfWork.Commit()
	return entityId, nil
}

// Get coupon
//
//	@Tags		coupon
//	@Summary	get all coupons
//	@Accept		json
//	@Produce	json
//	@Param      filter    		query     	string  false  "filters for data"
//	@Success	200				{object}	ApiResponse[[]dtos.CouponProfileDetail]
//	@Failure	400				{object}	ApiResponse[ApiError]
//	@Failure	500				{object}	ApiResponse[ApiError]
//	@Router		/coupon [GET]
//
// @Security Bearer
func (c *CouponController) getall(ctx *gin.Context, emptyDto dtos.EmptyDto) ([]dtos.CouponProfileDetail, error) {
	coupanService := services.CouponService{
		UnitOfWork:      factories.CreateUnitOfWork(c.sconfig),
		DepartmentProxy: *proxies.NewDepartmentProxy(fmt.Sprintf("%s:%d", c.sconfig.RemoteService.DeparmentServiceIp, c.sconfig.RemoteService.DeparmentServicePort)),
	}

	filters := c.base.ParseFilter(ctx.Request.URL.Query().Get("filter"))
	couponsDetail, err := coupanService.GetAll(factories.CreateServiceContext(ctx), filters)
	if err != nil {
		return nil, err
	}

	return couponsDetail, nil
}

// Delete
//
//	@Tags		coupon
//	@Summary	delete coupons by key
//	@Accept		json
//	@Produce	json
//	@Param      ids   		     body     	[]CouponProfileId  false  "coupan ids"
//	@Success	200				{object}    bool
//	@Failure	400				{object}	ApiResponse[ApiError]
//	@Failure	500				{object}	ApiResponse[ApiError]
//	@Router		/coupon [DELETE]
//
// @Security Bearer
func (c *CouponController) delete(ctx *gin.Context, keys []dtos.CouponProfileId) (bool, error) {
	coupanService := services.CouponService{
		UnitOfWork: factories.CreateUnitOfWork(c.sconfig),
	}

	return coupanService.Delete(factories.CreateServiceContext(ctx), keys)
}
