package main

import (
	"SocialServiceAincrad/internal/database"
	getfromenv "SocialServiceAincrad/internal/get_from_env"
	"SocialServiceAincrad/service/handlers"
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("../../configs/.env"); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	if err := run(); err != nil {
		log.Fatal(err)
	}

}

func run() error {
	err := connectToDatabase()
	if err != nil {
		return err
	}
	defer database.DB.Close(context.Background())

	err = database.CreateBaseTables()
	if err != nil {
		return err
	}

	router := startRouter()

	handlers.AllRequests(router)

	router.Run(":8080")
	return nil
}

func connectToDatabase() error {
	dbUrl := getfromenv.GetDatabaseConData()

	err := database.DB_Init(dbUrl)
	if err != nil {
		return err
	}

	err = database.DB.Ping(context.Background())
	if err != nil {
		return fmt.Errorf("error while Ping db connection: %v", err)
	}

	return nil
}

func startRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(corsMiddleware())

	return router
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, Range, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
