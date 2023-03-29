package main

import (
	"fmt"
	"log"
	"simple-crud/model/entity"

	"github.com/wpcodevo/golang-gorm-postgres/initializers"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&model.Person{})
	fmt.Println("? Migration complete")
}
