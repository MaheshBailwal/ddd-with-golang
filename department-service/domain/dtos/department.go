package dtos

type DepartmentSet struct {
	Number        int64  `binding:"required"`
	Name          string `binding:"required"`
	ReferenceCode string
} //@name DepartmentSet

type DepartmentDetail struct {
	Number        int64
	Name          string
	ReferenceCode string
	ModifyUser    string
	ModifyDate    string
}

type DepartmentNumber struct {
	Number int64
} //@name DepartmentId
