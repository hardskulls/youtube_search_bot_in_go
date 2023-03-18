package youtube_related

import (
	"context"
	config2 "youtube_search_go_bot/config"

	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

// TODO: Any memory safety issues with passing & returning refs?
func newYouTubeService(token *oauth2.Token) (*youtube.Service, error) {
	ctx := context.Background()
	config, err := config2.LoadConfig()
	if err != nil {
		return nil, err
	}

	tokSource := config.TokenSource(ctx, token)

	return youtube.NewService(ctx, option.WithTokenSource(tokSource))
}
