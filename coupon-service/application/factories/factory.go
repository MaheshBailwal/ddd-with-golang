package factories

import (
	"context"
	"fmt"

	"github.com/MaheshBailwal/mscore/core"
	"github.com/MaheshBailwal/mscore/factories"
	corep "github.com/MaheshBailwal/mscore/persistence"
	"github.com/gin-gonic/gin"

	"invafresh.com/coupon-service/domain/interfaces"
	"invafresh.com/coupon-service/infrastructure/config"
	"invafresh.com/coupon-service/persistence"
	"invafresh.com/coupon-service/persistence/mssql"
)

func CreateUnitOfWork(sconfig config.SConfig) interfaces.IUnitOfWork {

	con := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;",
		sconfig.MsSql.Server, sconfig.MsSql.User, sconfig.MsSql.Password, sconfig.MsSql.Port, sconfig.MsSql.Database)
	dbc := corep.DBContext{
		ConnectionString: con,
		DBType:           "mssql",
		Schema:           "periscope",
	}

	deptRepo := mssql.NewCouponRepository(&dbc)
	return persistence.NewUnitOfWork(&dbc, &deptRepo)
}

func CreateServiceContext(ctx *gin.Context) core.ServiceContext {
	return factories.CreateServiceContext(ctx)
}

func CreateServiceContextGrpc(ctx context.Context) core.ServiceContext {
	return factories.CreateServiceContextGrpc(ctx)
}
