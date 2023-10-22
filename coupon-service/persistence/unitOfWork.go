package persistence

import (
	"github.com/MaheshBailwal/mscore/persistence"
	"invafresh.com/coupon-service/domain/interfaces"
)

type UnitOfWork struct {
	dbContext        *persistence.DBContext
	couponRepository interfaces.ICouponProfileRepository
}

func NewUnitOfWork(dbContext *persistence.DBContext,
	couponRepo interfaces.ICouponProfileRepository) interfaces.IUnitOfWork {
	return &UnitOfWork{
		dbContext:        dbContext,
		couponRepository: couponRepo,
	}
}

func (r *UnitOfWork) CouponProfileRepository() interfaces.ICouponProfileRepository {
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
