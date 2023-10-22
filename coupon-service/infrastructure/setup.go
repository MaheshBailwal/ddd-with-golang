package infrastructure

import (
	"os"

	"invafresh.com/coupon-service/infrastructure/config"

	"github.com/ehsandavari/go-logger"
	"github.com/ehsandavari/go-tracer"
)

type Infrastructure struct {
	SConfig *config.SConfig
	ILogger logger.ILogger

	ITracer tracer.ITracer
}

func NewInfrastructure() *Infrastructure {
	sConfig := config.NewConfig()

	value, found := os.LookupEnv("SQL_SERVER")
	if found {
		sConfig.MsSql.Server = value
		sConfig.MsSql.User = os.Getenv("SQL_USER")
		sConfig.MsSql.Password = os.Getenv("SQL_PWD")
		sConfig.MsSql.Database = os.Getenv("SQL_DB")
		sConfig.RemoteService.DeparmentServiceIp = os.Getenv("DEPT_SERVICE")
	}

	return &Infrastructure{
		SConfig: sConfig,
		ILogger: logger.NewLogger(sConfig.Logger.Level, sConfig.Logger.Mode, sConfig.Logger.Encoder),

		ITracer: tracer.NewTracer(
			sConfig.Tracer.IsEnabled,
			sConfig.Tracer.Host,
			sConfig.Tracer.Port,
			sConfig.Service.Id,
			sConfig.Service.Name,
			sConfig.Service.Version,
			"production",
			sConfig.Tracer.Sampler,
			sConfig.Tracer.UseStdout,
		),
	}
}

func (r *Infrastructure) Close() {
	if err := r.ILogger.Sync(); err != nil {
		r.ILogger.Error("error in sync logger : ", err)
	}

	// if err := r.SPostgres.Close(); err != nil {
	// 	r.ILogger.Error("error in close postgres : ", err)
	// }

	if err := r.ITracer.Shutdown(); err != nil {
		r.ILogger.Error("error in sync logger : ", err)
	}
}
