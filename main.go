package main

import (
	"log"
	"os"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main(){

	godotenv.Load()

	Port := os.Getenv("PORT")
	if Port == "" {
		log.Fatal("PORT env variable is not set")
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 				300,			
	}))

	server := &http.Server{
		Handler: router,
		Addr: ":" + Port,
	}

	log.Printf("Server starting on port: %v", Port)
	err := server.ListenAndServe()
	if err != nil{
		log.Fatal(err)
	}

}
