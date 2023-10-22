package application

import (
	"github.com/ehsandavari/go-logger"
	"invafresh.com/department-service/application/api"
	"invafresh.com/department-service/application/grpc"
	"invafresh.com/department-service/infrastructure/config"
)

type Application struct {
	SConfig *config.SConfig
	ILogger logger.ILogger
	sApi    *api.SApi
	gApi    *grpc.GApi
}

func NewApplication(sConfig *config.SConfig, iLogger logger.ILogger) *Application {
	return &Application{
		SConfig: sConfig,
		ILogger: iLogger,
		sApi:    api.NewSApi(sConfig, iLogger),
		gApi:    grpc.NewGApi(sConfig, iLogger),
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
