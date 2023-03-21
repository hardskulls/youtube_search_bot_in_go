package youtube_related

import (
	"context"
	"errors"
	"golang.org/x/oauth2"
	"google.golang.org/api/youtube/v3"
	"strings"
	"youtube_search_go_bot/internal/keyboards"
	"youtube_search_go_bot/internal/logging"
)

func compareSubscriptionsBy(sub *youtube.Subscription, searchIn keyboards.SearchIn) string {
	res := ""
	switch searchIn {
	case keyboards.SearchInTitle:
		res = sub.Snippet.Title
	case keyboards.SearchInDescription:
		res = sub.Snippet.Description
	}
	return res
}

func comparePlaylistsBy(pl *youtube.Playlist, searchIn keyboards.SearchIn) string {
	res := ""
	switch searchIn {
	case keyboards.SearchInTitle:
		res = pl.Snippet.Title
	case keyboards.SearchInDescription:
		res = pl.Snippet.Description
	}
	return res
}

func SearchSubscriptions(token *oauth2.Token, textToSearch string, searchIn keyboards.SearchIn, resultLim uint16) ([]*youtube.Subscription, error) {
	ctx := context.Background()
	youtubeService, err := newYouTubeService(token)
	if err != nil {
		return nil, err
	}

	part := []string{"contentDetails", "id", "snippet"}
	request := youtubeService.Subscriptions.List(part).Mine(true).MaxResults(50)

	buf := make([]*youtube.Subscription, 0)

	err = request.Pages(ctx, func(slr *youtube.SubscriptionListResponse) error {
		for _, sub := range slr.Items {
			if sub != nil {
				s := strings.ToLower(compareSubscriptionsBy(sub, searchIn))
				substr := strings.ToLower(textToSearch)
				if strings.Contains(s, substr) {
					if len(buf) <= int(resultLim) {
						buf = append(buf, sub)
					} else {
						break
					}
				}
			}
		}
		return nil
	})
	if err != nil {
		logging.LogFuncEnd("SearchSubscriptions")
		return buf, errors.New("request.Pages() returned an error")
	}
	if len(buf) < 1 {
		logging.LogFuncEnd("SearchSubscriptions")
		return buf, errors.New("no channel title matches the specified string")
	}

	logging.LogFuncEnd("SearchSubscriptions")
	return buf, nil
}

func SearchPlaylists(token *oauth2.Token, textToSearch string, searchIn keyboards.SearchIn, resultLim uint16) ([]*youtube.Playlist, error) {
	ctx := context.Background()
	youtubeService, err := newYouTubeService(token)
	if err != nil {
		return nil, err
	}

	part := []string{"contentDetails", "id", "snippet"}
	request := youtubeService.Playlists.List(part).Mine(true).MaxResults(50)

	buf := make([]*youtube.Playlist, 0)

	err = request.Pages(ctx, func(resp *youtube.PlaylistListResponse) error {
		for _, pl := range resp.Items {
			if pl != nil {
				s := strings.ToLower(comparePlaylistsBy(pl, searchIn))
				substr := strings.ToLower(textToSearch)
				if strings.Contains(s, substr) {
					if len(buf) <= int(resultLim) {
						buf = append(buf, pl)
					} else {
						break
					}
				}
			}
		}
		return nil
	})
	if err != nil {
		logging.LogFuncEnd("SearchPlaylists")
		return buf, errors.New("request.Pages() returned an error")
	}
	if len(buf) < 1 {
		logging.LogFuncEnd("SearchPlaylists")
		return buf, errors.New("no channel title matches the specified string")
	}

	logging.LogFuncEnd("SearchPlaylists")
	return buf, nil
}
