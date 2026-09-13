package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Codrul/go-rss-aggregator/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		UserName string `json:"user_name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err !=nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %s", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserName: params.UserName,
		
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user:%s", err))
		return
	}


	respondWithJSON(w, 201, user)

}


func (apiCfg *apiConfig) handlerGetAllUsers(w http.ResponseWriter, r *http.Request) {
	type User struct {
		Id string `json:"id"`
		UserName string `json:"user_name"`
	}

	users, err := apiCfg.DB.GetAllUsers(r.Context())
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error getting users: %s", err))
	}

	response := []User{}

	for _, user := range users {
		response = append(response, User{
			Id: user.ID.String(),
			UserName: user.UserName,
		})
	}

	respondWithJSON(w, 200, response)

}
