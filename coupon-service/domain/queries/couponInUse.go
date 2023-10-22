package queries

type CouponInUse struct {
	CouponKey int64
}

func (q CouponInUse) GetQueryName() string {
	return "CouponInUse"
}
