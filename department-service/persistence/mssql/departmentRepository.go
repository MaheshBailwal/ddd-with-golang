package mssql

import (
	"invafresh.com/department-service/domain/entities"
	//"invafresh.com/department-service/persistence"
	"github.com/MaheshBailwal/mscore/persistence"
)

type DepartmentRepository struct {
	persistence.GenericRepository[entities.Departments]
}

func NewDepartmentRepository(db *persistence.DBContext) DepartmentRepository {
	return DepartmentRepository{
		GenericRepository: persistence.NewGenericRepository[entities.Departments](db),
	}
}
