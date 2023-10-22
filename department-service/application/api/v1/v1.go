package v1

import (
	"github.com/ehsandavari/go-logger"
	"github.com/gin-gonic/gin"
	_ "invafresh.com/department-service/application/api/docs"
	"invafresh.com/department-service/application/api/v1/controllers"
	"invafresh.com/department-service/infrastructure/config"
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
		controllers.NewDepartmentController(apiRouterGroup, logger, sconfig)
	}

}
