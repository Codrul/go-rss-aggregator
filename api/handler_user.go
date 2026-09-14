package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/Codrul/go-rss-aggregator/internal/database"
)

func (ApiCfg *ApiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		UserName string `json:"user_name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err !=nil {
		RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %s", err))
		return
	}

	user, err := ApiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserName: params.UserName,
		
	})
	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("Couldn't create user:%s", err))
		return
	}


	RespondWithJSON(w, 201, user)

}


func (ApiCfg *ApiConfig) HandlerGetAllUsers(w http.ResponseWriter, r *http.Request) {
	type User struct {
		Id string `json:"id"`
		UserName string `json:"user_name"`
	}

	users, err := ApiCfg.DB.GetAllUsers(r.Context())
	if err != nil {
		RespondWithError(w, 500, fmt.Sprintf("Error getting users: %s", err))
	}

	response := []User{}

	for _, user := range users {
		response = append(response, User{
			Id: user.ID.String(),
			UserName: user.UserName,
		})
	}

	RespondWithJSON(w, 200, response)

}


func (ApiCfg *ApiConfig) HandlerGetUser(w http.ResponseWriter, r *http.Request) {
}












