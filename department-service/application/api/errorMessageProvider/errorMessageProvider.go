package errorMessageProvider

import (
	"net/http"

	"github.com/MaheshBailwal/mscore/api"
	"github.com/MaheshBailwal/mscore/core"
	"invafresh.com/department-service/domain/constants"
)

type ErrorMessageProvider struct{}

func NewErrorMessageProvider() api.IErrorMessageProvider {
	return ErrorMessageProvider{}
}

func (e ErrorMessageProvider) GetErrorMessage() func(core.ErrCode) string {
	// innerMap is captured in the closure returned below
	innerMap := map[core.ErrCode]string{
		constants.CouponNotFound: "Coupon Not Found",
		constants.CouponInUse:    "Coupon In Use. Can not be deleted",
	}

	return func(key core.ErrCode) string {
		return innerMap[key]
	}
}

func (e ErrorMessageProvider) GetHttpStatus(err core.ServiceError) int {
	switch err.Code {
	case constants.CouponNotFound:
		return http.StatusNotFound
	case constants.CouponInUse:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}

}
