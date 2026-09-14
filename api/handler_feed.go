package api

import (
	"net/http"
	"fmt"
	"time"
	"encoding/json"

	"github.com/Codrul/go-rss-aggregator/internal/database"
	"github.com/google/uuid"
)

func (ApiCfg *ApiConfig) HandlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameters struct {
		Name string `json:"user_name"`
		URL string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err !=nil {
		RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %s", err))
		return
	}

	feed, err := ApiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
		Url: params.URL,
		UserID: user.ID,
		
	})
	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("Couldn't create feed: %s", err))
		return
	}


	RespondWithJSON(w, 201, databaseFeedToFeed(feed))

}


func (ApiCfg *ApiConfig) HandlerGetFeeds(w http.ResponseWriter, r *http.Request){

	feeds, err := ApiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("Couldn't get feeds: %s", err))
		return
	}

	RespondWithJSON(w, 200, databaseFeedsToFeeds(feeds))
}



