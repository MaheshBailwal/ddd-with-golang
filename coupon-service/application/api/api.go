package api

import (
	"context"

	"net/http"

	docs "invafresh.com/coupon-service/application/api/docs"
	"invafresh.com/coupon-service/infrastructure/config"

	"github.com/MaheshBailwal/mscore/api/middlewares"
	"github.com/ehsandavari/go-logger"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	v1 "invafresh.com/coupon-service/application/api/v1"
)

type SApi struct {
	server  *http.Server
	sConfig *config.SConfig
	iLogger logger.ILogger
}

func NewSApi(sConfig *config.SConfig, iLogger logger.ILogger) *SApi {
	var sApi SApi
	sApi.sConfig = sConfig
	if sConfig.Service.Api.IsEnabled {
		sApi.server = &http.Server{
			Addr: sConfig.Service.Api.Host + ":" + sConfig.Service.Api.Port,
		}
		sApi.iLogger = iLogger
	}
	return &sApi
}

func (r *SApi) Start() {
	if r.sConfig.Service.Api.IsEnabled {
		gin.SetMode(r.sConfig.Service.Api.Mode)
		engine := gin.Default()
		docs.SwaggerInfo.BasePath = "/api/v1"
		engine.Use(
			middlewares.Cors(),
			middlewares.RequestId(),
			middlewares.AuthMiddleware(),
		)

		monitoringRouterGroup := engine.Group("/-")
		{
			monitoringRouterGroup.GET("/health", func(ctx *gin.Context) { ctx.Status(http.StatusOK) })
			monitoringRouterGroup.GET("/liveness", func(ctx *gin.Context) { ctx.Status(http.StatusOK) })
			monitoringRouterGroup.GET("/readiness", func(ctx *gin.Context) { ctx.Status(http.StatusOK) })
			monitoringRouterGroup.GET("/metrics", gin.WrapH(promhttp.Handler()))
		}

		engine.GET("/swagger/*any", ginSwagger.WrapHandler(
			swaggerFiles.Handler,
		))

		apiRouterGroup := engine.Group("/api")
		{
			v1.Setup(apiRouterGroup, r.iLogger, *r.sConfig)
		}

		go func() {
			r.server.Handler = engine.Handler()
			if err := r.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				r.iLogger.Fatal("error in serve api server", err)
			}
		}()
		r.iLogger.Info("api server start on port : ", r.sConfig.Service.Api.Port)
	}
}

func (r *SApi) Stop() {
	if r.sConfig.Service.Api.IsEnabled {
		err := r.server.Shutdown(context.Background())
		if err != nil {
			r.iLogger.Error("error in shutdown api server", err)
		}
	}
}
