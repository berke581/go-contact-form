package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/berke581/go-contact-form/api/v1/router"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found.")
	}
}

func main() {
	router.ServeRouter()
}
