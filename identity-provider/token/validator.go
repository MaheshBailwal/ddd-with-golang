package token

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/microsoft/go-mssqldb"
)

type OwnerPasswordValidator struct {
	DBConnectionString string
	isConOpen          bool
	db                 sqlx.DB
}

func NewOwnerPasswordValidator(connString string) OwnerPasswordValidator {
	return OwnerPasswordValidator{
		DBConnectionString: connString,
	}
}

func (v *OwnerPasswordValidator) Validate(request TokenRequest) (User, bool) {

	query := fmt.Sprintf("Select user_id from [Periscope].[Users] where User_ID='%s'", request.UserName)
	fmt.Println(query)
	v.openConnection()
	row := v.db.QueryRowx(query)
	user := User{}
	err := row.StructScan(&user)

	if err != nil {
		fmt.Println("Error->", err)
		return user, false
	}

	return user, true
}

func (v *OwnerPasswordValidator) openConnection() {
	if v.isConOpen {
		return
	}
	db, err := sqlx.Connect("mssql", v.DBConnectionString)
	if err != nil {
		log.Fatalln(err)
	}
	v.db = *db
	v.isConOpen = true
}
