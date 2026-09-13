package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main(){

	godotenv.Load()

	Port := os.Getenv("PORT")
	if Port == "" {
		log.Fatal("PORT env variable is not set")
	}

	fmt.Printf("")
}
