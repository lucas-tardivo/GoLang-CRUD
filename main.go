package main

import (
	"log"
	"net/http"
	"simple-crud/router"
	"simple-crud/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/wpcodevo/golang-gorm-postgres/initializers"
)

var (
	server              *gin.Engine

	PersonController      controller.PersonController
	PersonRouteController router.PersonRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	PersonController = controller.NewPersonController(initializers.DB)
	PersonRouteController = router.NewRoutePersonController(PersonController)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8080", config.ClientOrigin}

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	PersonRouteController.StartRouter(router)
	log.Fatal(server.Run(":" + config.ServerPort))
}