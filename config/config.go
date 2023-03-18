package config

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func LoadConfig() (*oauth2.Config, error) {
	bytes, err := os.ReadFile("client_secret_web_app.json")
	if err != nil {
		return nil, err
	}

	config, err := google.ConfigFromJSON(bytes, "https://www.googleapis.com/auth/youtube.readonly")
	if err != nil {
		return nil, err
	}

	return config, nil
}
