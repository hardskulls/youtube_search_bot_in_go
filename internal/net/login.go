package net

import (
	"golang.org/x/oauth2"
	"strconv"
	config2 "youtube_search_go_bot/internal/config"
	dialogue2 "youtube_search_go_bot/internal/dialogue"
	"youtube_search_go_bot/internal/errors"
)

// Creates url for YouTube authentification.
func LogInUrl(userId int64) string {
	state := dialogue2.BuildKeyValueString(
		[]dialogue2.KVStruct{
			{Key: "state_code", Value: "kut987987_576fg78d5687lojfvkzr_85y6_435sgred_vnhgx_gdut"},
			{Key: "for_user", Value: strconv.FormatInt(userId, 10)},
		},
	)
	config, err := config2.LoadConfig()
	errors.ExitOnError(err)
	return config.AuthCodeURL(state, oauth2.AccessTypeOffline)
}
