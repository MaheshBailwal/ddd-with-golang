package main

import (
	"os"

	"invafresh.com/identity-provider/api"
	"invafresh.com/identity-provider/config"
)

func main() {
	sConfig := config.NewConfig()

	value, found := os.LookupEnv("SQL_SERVER")
	if found {
		sConfig.MsSql.Server = value
		sConfig.MsSql.User = os.Getenv("SQL_USER")
		sConfig.MsSql.Password = os.Getenv("SQL_PWD")
		sConfig.MsSql.Database = os.Getenv("SQL_DB")
	}

	apis := api.NewApi(sConfig)
	apis.Start()
}
