package main

import (
	"log"
	"net/http"
	"os"
	"database/sql"
	// internal packages
	"github.com/Codrul/go-rss-aggregator/internal/database"
	"github.com/Codrul/go-rss-aggregator/api"

	// external packages
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // _ here means we don't call this directly in the code
)



func main(){

	godotenv.Load()

	Port := os.Getenv("PORT")
	if Port == "" {
		log.Fatal("PORT env variable is not set")
	}

	dbConn := os.Getenv("DB_URL")
	if dbConn == "" {
		log.Fatal("dbConn env variable is not set")
	}


	conn, err := sql.Open("postgres", dbConn)
	if err != nil {
		log.Fatal("Can't connect to the database:", err)
	}

	apiCfg := api.ApiConfig{
		DB: database.New(conn),
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

	v1Router := chi.NewRouter()
	v1Router.Get("/health", api.HandlerReadiness)
	v1Router.Get("/err", api.HandlerErr)
	v1Router.Post("/users", apiCfg.HandlerCreateUser)
	v1Router.Get("/all-users", apiCfg.HandlerGetAllUsers)
	v1Router.Get("/users", apiCfg.MiddlewareAuth(apiCfg.HandlerGetUser))
	v1Router.Post("/feeds", apiCfg.MiddlewareAuth(apiCfg.HandlerCreateFeed))



	router.Mount("/v1", v1Router)

	log.Printf("Server starting on port: %v", Port)
	error := server.ListenAndServe()
	if error != nil{
		log.Fatal(err)
	}

}
