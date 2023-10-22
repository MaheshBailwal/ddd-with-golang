package persistence

import (
	"github.com/MaheshBailwal/mscore/persistence"
	"invafresh.com/department-service/domain/interfaces"
)

type UnitOfWork struct {
	dbContext        *persistence.DBContext
	couponRepository interfaces.IDepartmentRepository
}

func NewUnitOfWork(dbContext *persistence.DBContext,
	couponRepo interfaces.IDepartmentRepository) interfaces.IUnitOfWork {
	return &UnitOfWork{
		dbContext:        dbContext,
		couponRepository: couponRepo,
	}
}

func (r *UnitOfWork) DepartmentRepository() interfaces.IDepartmentRepository {
	return r.couponRepository
}

func (r *UnitOfWork) Commit() {
	if r.dbContext.IsTranOpen {
		r.dbContext.Transcation.Commit()
		r.dbContext.IsTranOpen = false
	}
}

func (r *UnitOfWork) Rollback() {
	if r.dbContext.IsTranOpen {
		r.dbContext.Transcation.Rollback()
		r.dbContext.IsTranOpen = false
	}
}

func (r *UnitOfWork) Conclude(err error) {
	if r.dbContext.IsTranOpen {
		if err == nil {
			r.dbContext.Transcation.Commit()
		} else {
			r.dbContext.Transcation.Rollback()
		}

		r.dbContext.IsTranOpen = false
	}
}
