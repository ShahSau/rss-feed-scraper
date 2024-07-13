package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello, World!")
	godotenv.Load(".env")
	port := os.Getenv(("PORT"))
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	fmt.Println("Server is running on port: ", port)
}
