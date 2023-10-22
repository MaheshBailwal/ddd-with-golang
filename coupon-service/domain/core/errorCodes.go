package core

type ErrCode int

const (
	CouponNotFound ErrCode = 1001
	CouponInUse    ErrCode = 1002
)

type ErrorMsg struct {
	Code    ErrCode
	Message string
}

func GetErrorMsg() func(ErrCode) string {
	// innerMap is captured in the closure returned below
	innerMap := map[ErrCode]string{
		CouponNotFound: "Coupon Not Found",
		CouponInUse:    "Coupon In Use. Can not be deleted",
	}

	return func(key ErrCode) string {
		return innerMap[key]
	}
}
