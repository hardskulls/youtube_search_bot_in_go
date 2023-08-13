package net

import (
	"os"
	"strconv"
	config2 "youtube_search_go_bot/internal/config"
	dialogue2 "youtube_search_go_bot/internal/dialogue"
	"youtube_search_go_bot/internal/errors"

	"golang.org/x/oauth2"
)

// Creates url for YouTube authentification.
func LogInUrl(userId int64) string {
	state := dialogue2.BuildKeyValueString(
		[]dialogue2.KVStruct{
			{Key: "state_code", Value: os.Getenv("STATE_CODE")},
			{Key: "for_user", Value: strconv.FormatInt(userId, 10)},
		},
	)
	config, err := config2.LoadConfig()
	errors.ExitOnError(err)
	return config.AuthCodeURL(state, oauth2.AccessTypeOffline)
}
