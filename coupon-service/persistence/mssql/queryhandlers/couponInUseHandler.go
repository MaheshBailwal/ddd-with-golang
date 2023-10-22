package queryhandlers

import (
	"fmt"

	"github.com/MaheshBailwal/mscore/persistence"
	"invafresh.com/coupon-service/domain/queries"
)

type IQueryHandler interface {
	HandlesQuery() string
	Handle(persistence.IDomainQuery, *persistence.DBContext) any
}

type CouponInUseHandler struct {
}

func (h CouponInUseHandler) HandlesQuery() string {
	return "CouponInUseHandler"
}

func (h CouponInUseHandler) Handle(q persistence.IDomainQuery, dBContext *persistence.DBContext) any {

	couponQuery := q.(queries.CouponInUse)

	query := fmt.Sprintf("SELECT count(*) as c FROM [periscope].[MarkdownProfiles_V4] where Method='Coupon' and Value=%d", couponQuery.CouponKey)
	rows, _ := dBContext.ExecuteQuery(query)

	var count int64
	rows.Next()
	err := rows.Scan(&count)

	if err != nil {
		return false
	}

	return count > 0
}
