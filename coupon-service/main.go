package main

import (
	"log"
	"time"

	"invafresh.com/coupon-service/application"
	"invafresh.com/coupon-service/infrastructure"

	"github.com/ehsandavari/go-graceful-shutdown"
	"github.com/joho/godotenv"
)

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	loadEnv()
	run()
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("error in loading .env file", err)
	}
}

func run() {
	newInfrastructure := infrastructure.NewInfrastructure()
	//newPersistence := persistence.NewPersistence(newInfrastructure.ILogger, newInfrastructure.SPostgres)
	newApplication := application.NewApplication(newInfrastructure.SConfig, newInfrastructure.ILogger)
	newApplication.Setup()

	shutdownFunc := func() {
		newApplication.Close()
	}
	cleanupFunc := func() {
		newInfrastructure.Close()
	}
	graceful.Shutdown(shutdownFunc, cleanupFunc, time.Duration(newInfrastructure.SConfig.Service.GracefulShutdownSecond)*time.Second)
}

/*
// Go connection Sample Code:
package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/microsoft/go-mssqldb"
)

var db *sql.DB
var server = "azuresql-mb.database.windows.net"
var port = 1433
var user = "mahesh"
var password = "password@123"
var database = "inva-poc"

func main() {
	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)
	var err error
	fmt.Println(connString)
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!")
}
*/
