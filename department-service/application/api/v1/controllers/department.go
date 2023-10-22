package controllers

import (
	"github.com/MaheshBailwal/mscore"
	"github.com/MaheshBailwal/mscore/api"
	"github.com/ehsandavari/go-logger"
	"github.com/gin-gonic/gin"
	emsgp "invafresh.com/department-service/application/api/errorMessageProvider"
	"invafresh.com/department-service/application/factories"
	"invafresh.com/department-service/domain/dtos"
	"invafresh.com/department-service/domain/services"
	"invafresh.com/department-service/infrastructure/config"
)

type DepartmentController struct {
	base    api.BaseController
	sconfig config.SConfig
	logger  logger.ILogger
}

func NewDepartmentController(routerGroup *gin.RouterGroup, logger logger.ILogger, sconfig config.SConfig) {
	departmentController := DepartmentController{
		base:    api.NewBaseController(),
		sconfig: sconfig,
		logger:  logger,
	}
	errMsgProvider := emsgp.NewErrorMessageProvider()
	routerGroup = routerGroup.Group("department")
	{
		routerGroup.PUT("", api.RequestHandler[*dtos.DepartmentSet, int64](departmentController.update).Handle(departmentController.logger, errMsgProvider))
		routerGroup.GET("", api.RequestHandler[dtos.EmptyDto, []dtos.DepartmentDetail](departmentController.getall).Handle(departmentController.logger, errMsgProvider))
		routerGroup.DELETE("", api.RequestHandler[[]dtos.DepartmentNumber, bool](departmentController.delete).Handle(departmentController.logger, errMsgProvider))
	}

	mscore.LogInfo("asas")
}

// Set department
//
//	@Tags		department
//	@Summary	insert/update department
//	@Accept		json
//	@Produce	json
//	@Param		params			body		DepartmentSet				false	"Query Params"
//	@Success	200				{object}	ApiResponse[int]
//	@Failure	400				{object}	ApiResponse[ApiError]
//	@Failure	500				{object}	ApiResponse[ApiError]
//	@Router		/department [PUT]
//
// @Security Bearer
func (d *DepartmentController) update(ctx *gin.Context, dto *dtos.DepartmentSet) (int64, error) {
	departmentService := services.DepartmentService{
		UnitOfWork: factories.CreateUnitOfWork(d.sconfig),
	}

	entityId, err := departmentService.Update(factories.CreateServiceContext(ctx), dto)
	if err != nil {
		return -1, err
	}

	return entityId, nil
}

// Get department
//
//	@Tags		department
//	@Summary	get all departments
//	@Accept		json
//	@Produce	json
//	@Param      filter    		query     	string  false  "filters for data"
//	@Success	200				{object}	ApiResponse[[]dtos.DepartmentDetail]
//	@Failure	400				{object}	ApiResponse[ApiError]
//	@Failure	500				{object}	ApiResponse[ApiError]
//	@Router		/department [GET]
//
// @Security Bearer
func (d *DepartmentController) getall(ctx *gin.Context, emptyDto dtos.EmptyDto) ([]dtos.DepartmentDetail, error) {
	departmentService := services.DepartmentService{
		UnitOfWork: factories.CreateUnitOfWork(d.sconfig),
	}

	filters := d.base.ParseFilter(ctx.Request.URL.Query().Get("filter"))
	departmentsDetail, err := departmentService.GetAll(factories.CreateServiceContext(ctx), filters)
	if err != nil {
		return nil, err
	}

	return departmentsDetail, nil
}

// Delete
//
//	@Tags		department
//	@Summary	delete department by number
//	@Accept		json
//	@Produce	json
//	@Param      ids   		     body     	[]DepartmentId  false  "department ids"
//	@Success	200				{object}    bool
//	@Failure	400				{object}	ApiResponse[ApiError]
//	@Failure	500				{object}	ApiResponse[ApiError]
//	@Router		/department [DELETE]
//
// @Security Bearer
func (d *DepartmentController) delete(ctx *gin.Context, keys []dtos.DepartmentNumber) (bool, error) {
	departmentService := services.DepartmentService{
		UnitOfWork: factories.CreateUnitOfWork(d.sconfig),
	}

	return departmentService.Delete(factories.CreateServiceContext(ctx), keys)
}
