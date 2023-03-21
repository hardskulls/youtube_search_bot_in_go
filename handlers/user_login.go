package handlers

import (
	"context"
	"errors"
	"net/http"
	"os"
	"youtube_search_go_bot/config"
	"youtube_search_go_bot/db"
	"youtube_search_go_bot/dialogue"
	"youtube_search_go_bot/logging"
)

// Handle OAuth2 code.
func GoogleCallbackHandler(w http.ResponseWriter, r *http.Request) {
	// All manually added parameters: user id, state code, etc.
	state := r.FormValue("state")

	// Additional value that to be checked for safety.
	stateCode, err := dialogue.ExtractValue(state, "state_code")
	if err != nil {
		logging.LogError(err)
		return
	}
	if stateCode != "kut987987_576fg78d5687lojfvkzr_85y6_435sgred_vnhgx_gdut" {
		err = errors.New("state code is empty")
		logging.LogError(err)
		return
	}

	userId, err := dialogue.ExtractValue(state, "for_user")
	if err != nil {
		logging.LogError(err)
		return
	}
	if userId == "" {
		err = errors.New("for_user is empty")
		logging.LogError(err)
		return
	}

	exchangeableCode := r.FormValue("code")
	if exchangeableCode == "" {
		err = errors.New("exchangeable code is empty")
		logging.LogError(err)
		return
	}

	oauthConfig, err := config.LoadConfig()
	if err != nil {
		logging.LogError(err)
		return
	}
	token, err := oauthConfig.Exchange(context.Background(), exchangeableCode)
	if err != nil {
		logging.LogError(err)
		return
	}
	if token == nil || token.AccessToken == "" {
		err = errors.New("token is empty")
		logging.LogError(err)
		return
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		logging.LogError(err)
		return
	}
	if err = db.SaveOAuthToken(userId, token, dbUrl); err != nil {
		logging.LogError(err)
		return
	}

	botRedirectUrl := os.Getenv("BOT_REDIRECT_URL")
	if botRedirectUrl == "" {
		logging.LogError(err)
		return
	}
	w.Header().Add("Location", botRedirectUrl)
	w.WriteHeader(http.StatusMovedPermanently)

	logging.LogFuncEnd("googleCallbackHandler")
}
