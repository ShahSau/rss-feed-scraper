package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	port := os.Getenv(("PORT"))
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	fmt.Println("Server is running on port: ", port)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		MaxAge:           300,
		AllowCredentials: false,
		ExposedHeaders:   []string{"Link"},
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/health", handlerReadiness)
	v1Router.Get("/error", handlerError)

	router.Mount("/v1", v1Router) // Mount the v1Router under /v1

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Server is running on port: %s", port)
	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}

}
