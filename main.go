package main

import (
	"log"
	"os"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main(){

	godotenv.Load()

	Port := os.Getenv("PORT")
	if Port == "" {
		log.Fatal("PORT env variable is not set")
	}

	router := chi.NewRouter()

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
