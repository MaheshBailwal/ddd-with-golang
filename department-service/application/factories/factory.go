package factories

import (
	"context"
	"fmt"

	"github.com/MaheshBailwal/mscore/core"
	"github.com/MaheshBailwal/mscore/factories"
	corep "github.com/MaheshBailwal/mscore/persistence"
	"github.com/gin-gonic/gin"

	"invafresh.com/department-service/domain/interfaces"
	"invafresh.com/department-service/infrastructure/config"
	"invafresh.com/department-service/persistence"
	"invafresh.com/department-service/persistence/mssql"
)

func CreateUnitOfWork(sconfig config.SConfig) interfaces.IUnitOfWork {

	con := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;",
		sconfig.MsSql.Server, sconfig.MsSql.User, sconfig.MsSql.Password, sconfig.MsSql.Port, sconfig.MsSql.Database)
	fmt.Println("con->", con)
	dbc := corep.DBContext{
		ConnectionString: con,
		DBType:           "mssql",
		Schema:           "plum",
	}

	deptRepo := mssql.NewDepartmentRepository(&dbc)
	return persistence.NewUnitOfWork(&dbc, &deptRepo)
}

func CreateServiceContext(ctx *gin.Context) core.ServiceContext {
	return factories.CreateServiceContext(ctx)
}

func CreateServiceContextGrpc(ctx context.Context) core.ServiceContext {
	return factories.CreateServiceContextGrpc(ctx)
}
