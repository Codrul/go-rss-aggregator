package api

import (
	"net/http"
	"fmt"

	"github.com/Codrul/go-rss-aggregator/internal/database"
	"github.com/Codrul/go-rss-aggregator/internal/database/auth"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (ApiCfg *ApiConfig) MiddlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		apiKey, err := auth.GetAPIKey(r.Header)

		if err != nil{
			RespondWithError(w, 403, fmt.Sprintf("auth error: %v", err))
			return
		}

		user, err := ApiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil{
			RespondWithError(w, 404, fmt.Sprintf("user not found: %v", err))
			return
		}

		handler(w, r, user)
	}
}

