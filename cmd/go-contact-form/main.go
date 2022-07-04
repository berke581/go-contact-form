package main

import (
	"log"
	// "os"

	"github.com/joho/godotenv"

	"github.com/berke581/go-contact-form/api/v1/router"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found.")
		// os.Exit(-1)
	}
}

func main() {
	router.ServeRouter()
}
