package main

import (
	"carrmod/backend/config"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Backend is ready")
	//load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//logging
	config.Logging()
	//database
	config.Database()
	//routes
	config.Web()
	log.Println("App closing")
}
