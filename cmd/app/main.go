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
	defer dbConnection.Close(context.Background())

	err = database.CreateBaseTables(dbConnection)
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
		return nil, fmt.Errorf("Error while getting database url: %v", err)
	}

	conn, err := database.DB_Init(dbUrl)
	if err != nil {
		return nil, err
	}

	err = conn.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("Error while Ping db connection: %v", err)
	}

	return conn, nil
}

func startRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	return router
}
