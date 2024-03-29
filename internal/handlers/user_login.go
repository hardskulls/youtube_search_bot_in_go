package handlers

import (
	"context"
	"errors"
	"net/http"
	"os"
	"youtube_search_go_bot/internal/config"
	"youtube_search_go_bot/internal/db"
	"youtube_search_go_bot/internal/dialogue"
	"youtube_search_go_bot/internal/logging"
)

// Handle OAuth2 code.
func GoogleCallbackHandler(w http.ResponseWriter, r *http.Request) {
	// All manually added parameters: user id, state code, etc.
	state := r.FormValue("state")
	defer r.Body.Close()

	// Additional value that to be checked for safety.
	stateCode, err := dialogue.ExtractValue(state, "state_code")
	if err != nil {
		logging.LogError(err)
		w.Write([]byte("requires value missing"))
		return
	}
	if stateCode != os.Getenv("STATE_CODE") {
		err = errors.New("invalid state code")
		logging.LogError(err)
		w.Write([]byte("invalid value"))
		return
	}

	userId, err := dialogue.ExtractValue(state, "for_user")
	if err != nil {
		logging.LogError(err)
		w.Write([]byte("internal errro"))
		return
	}
	if userId == "" {
		err = errors.New("for_user is empty")
		logging.LogError(err)
		w.Write([]byte("missing required value"))
		return
	}

	exchangeableCode := r.FormValue("code")
	if exchangeableCode == "" {
		err = errors.New("exchangeable code is empty")
		logging.LogError(err)
		w.Write([]byte("internal error"))
		return
	}

	oauthConfig, err := config.LoadConfig()
	if err != nil {
		logging.LogError(err)
		w.Write([]byte("internal error"))
		return
	}
	token, err := oauthConfig.Exchange(context.Background(), exchangeableCode)
	if err != nil {
		logging.LogError(err)
		w.Write([]byte("internal error"))
		return
	}
	if token == nil || token.AccessToken == "" {
		err = errors.New("token is empty")
		logging.LogError(err)
		w.Write([]byte("internal error"))
		return
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		logging.LogError(err)
		w.Write([]byte("internal error"))
		return
	}
	if err = db.SaveOAuthToken(userId, token, dbUrl); err != nil {
		logging.LogError(err)
		w.Write([]byte("internal error"))
		return
	}

	botRedirectUrl := os.Getenv("BOT_REDIRECT_URL")
	if botRedirectUrl == "" {
		logging.LogError(err)
		w.Write([]byte("internal error"))
		return
	}
	w.Header().Add("Location", botRedirectUrl)
	w.WriteHeader(http.StatusMovedPermanently)

	logging.LogFuncEnd("googleCallbackHandler")
}
