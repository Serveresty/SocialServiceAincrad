package main

import (
	"SocialServiceAincrad/internal/database"
	getfromjson "SocialServiceAincrad/internal/get_from_json"
	"SocialServiceAincrad/service/handlers"
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func main() {

	if err := run(); err != nil {
		log.Fatal(err)
	}

}

func run() error {
	dbConnection, err := connectToDatabase()
	if err != nil {
		return err
	}

	router := startRouter()

	handlers.AllRequests(router, dbConnection)

	router.Run(":8080")
	return nil
}

func connectToDatabase() (*pgx.Conn, error) {
	dbUrl, err := getfromjson.GetDatabaseConData()
	if err != nil {
		log.Fatalf("Error while getting database url: %v", err)
		return nil, err
	}

	conn, err := database.DB_Init(dbUrl)
	if err != nil {
		log.Fatalf("Error while connecting to database: %v", err)
		return nil, err
	}

	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatal("db err")
		return nil, err
	}
	fmt.Println("db works")
	return conn, nil
}

func startRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	return router
}
