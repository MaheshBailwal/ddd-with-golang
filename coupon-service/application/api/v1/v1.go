package v1

import (
	_ "invafresh.com/coupon-service/application/api/docs"
	"invafresh.com/coupon-service/application/api/v1/controllers"
	"invafresh.com/coupon-service/infrastructure/config"
	"github.com/ehsandavari/go-logger"
	"github.com/gin-gonic/gin"
)

//	@title			api
//	@version		1.0
//	@description	Example Api

//	@contact.name	Ehsan Davari
//	@contact.url	https://github.com/ehsandavari
//	@contact.email	ehsandavari.ir@gmail.com

//	@BasePath	/api/v1

func Setup(routerGroup *gin.RouterGroup, logger logger.ILogger, sconfig config.SConfig) {

	apiRouterGroup := routerGroup.Group("/v1")
	{
		controllers.NewCouponController(apiRouterGroup, logger, sconfig)
	}

}
