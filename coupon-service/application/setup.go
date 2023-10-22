package application

import (
	"invafresh.com/coupon-service/application/api"
	gapi "invafresh.com/coupon-service/application/grpc"
	"invafresh.com/coupon-service/infrastructure/config"
	"github.com/ehsandavari/go-logger"
)

type Application struct {
	SConfig *config.SConfig
	ILogger logger.ILogger
	sApi    *api.SApi
	gApi    *gapi.GApi
}

func NewApplication(sConfig *config.SConfig, iLogger logger.ILogger) *Application {
	return &Application{
		SConfig: sConfig,
		ILogger: iLogger,
		sApi:    api.NewSApi(sConfig, iLogger),
		gApi:    gapi.NewGApi(sConfig, iLogger),
	}
}

var Handlers []func(application *Application)

func (r *Application) Setup() {
	for _, handler := range Handlers {
		handler(r)
	}
	r.sApi.Start()
	r.gApi.Start()
}

func (r *Application) Close() {
	r.sApi.Stop()
}
