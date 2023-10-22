package entities

import (
	"database/sql"
)

type CouponProfiles struct {
	Id          int64 `dbfield:"CpProfile_Key" identity:"true"`
	DeptNo      int
	CouponPluNo sql.NullInt64 `dbfield:"Coupon_PluNo"`
	CouponUpc   string        `dbfield:"Coupon_Upc"`
	Value       int
	ModifyUser  string `dbfield:"UserId"`
	ModifyDate  int
}
