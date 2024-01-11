package main

import (
	"SocialServiceAincrad/internal/database"
	getfromjson "SocialServiceAincrad/internal/get_from_json"
	"SocialServiceAincrad/service/handlers"
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	dbUrl, err := getfromjson.GetDatabaseConData()
	if err != nil {
		log.Fatalf("Error while getting database url: %v", err)
	}

	conn, err := database.DB_Init(dbUrl)
	if err != nil {
		log.Fatalf("Error while connecting to database: %v", err)
	}

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatal("db err")
	}
	fmt.Println("db works")

	handlers.AllRequests(router)

	router.Run(":8080")
}
