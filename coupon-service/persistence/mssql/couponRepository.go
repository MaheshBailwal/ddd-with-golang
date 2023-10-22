package mssql

import (
	"github.com/MaheshBailwal/mscore/persistence"
	"invafresh.com/coupon-service/domain/entities"
	"invafresh.com/coupon-service/domain/queries"
	"invafresh.com/coupon-service/persistence/mssql/queryhandlers"
)

type CouponRepository struct {
	persistence.GenericRepository[entities.CouponProfiles]
}

func NewCouponRepository(db *persistence.DBContext) CouponRepository {
	queryhandlersMap := make(map[string]persistence.IQueryHandler)
	queryhandlersMap[queries.CouponInUse{}.GetQueryName()] = queryhandlers.CouponInUseHandler{}

	return CouponRepository{
		GenericRepository: persistence.NewGenericRepository[entities.CouponProfiles](db, queryhandlersMap),
	}
}
