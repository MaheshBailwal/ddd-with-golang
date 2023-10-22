package dtos

// import (
// 	"database/sql"
// )

type CouponProfileAdd struct {
	DeptNo      int `binding:"required"`
	CouponPluNo NullInt64
	CouponUpc   string `binding:"required"`
	Value       int    `binding:"required"`
} //@name CouponProfileAdd

type CouponProfileUpdate struct {
	Id          int
	DeptNo      int
	CouponPluNo NullInt64
	CouponUpc   string
	Value       int
} //@name CouponProfileUpdate

type CouponProfileDetail struct {
	Id          int
	DeptNo      int
	DeptName    string
	CouponPluNo NullInt64
	CouponUpc   string
	Value       int
	ModifyUser  string
	ModifyDate  string
}

type CouponProfileId struct {
	Id int64
}
