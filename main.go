package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Lgdev07/superapi/api/controllers"
	"github.com/joho/godotenv"
)

func main() {
	if _, err := os.Stat(".env"); !os.IsNotExist(err) {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error getting env, %v", err)
		} else {
			fmt.Println("We are getting values")
		}
	}

	server := controllers.Server{}
	server.Initialize(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	)
	server.Run()

}
