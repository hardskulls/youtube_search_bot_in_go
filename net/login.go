package net

import (
	"golang.org/x/oauth2"
	"strconv"
	config2 "youtube_search_go_bot/config"
	"youtube_search_go_bot/dialogue"
	"youtube_search_go_bot/errors"
)

// Creates url for YouTube authentification.
func LogInUrl(userId int64) string {
	state := dialogue.BuildKeyValueString(
		[]dialogue.KVStruct{
			{Key: "state_code", Value: "kut987987_576fg78d5687lojfvkzr_85y6_435sgred_vnhgx_gdut"},
			{Key: "for_user", Value: strconv.FormatInt(userId, 10)},
		},
	)
	config, err := config2.LoadConfig()
	errors.ExitOnError(err)
	return config.AuthCodeURL(state, oauth2.AccessTypeOffline)
}
