package interfaces

import (
	"github.com/MaheshBailwal/mscore/persistence"
	"invafresh.com/department-service/domain/entities"
)

type (
	IDepartmentRepository interface {
		persistence.IGenericRepository[entities.Departments]
	}
	IUnitOfWork interface {
		DepartmentRepository() IDepartmentRepository
		Commit()
		Rollback()
		Conclude(error)
	}
)
