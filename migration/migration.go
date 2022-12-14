package main

import (
	"fmt"
	"github.com/cnugroho11/movie_api/initializers"
	"github.com/cnugroho11/movie_api/models"
	"log"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load env")
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.Movie{})

	fmt.Println("Migration complete")
}
