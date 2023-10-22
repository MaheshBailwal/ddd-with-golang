package entities

import "database/sql"

//TODO: make name singular department
type Departments struct {
	Number        int64  `dbfield:"DeptNo" pk:"true"`
	Name          string `dbfield:"DeptName"`
	ReferenceCode sql.NullString
	ModifyUser    string `dbfield:"UserId"`
	ModifyDate    int
}
