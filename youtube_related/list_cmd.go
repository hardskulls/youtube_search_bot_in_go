package youtube_related

import (
	"context"
	"errors"
	"golang.org/x/oauth2"
	"google.golang.org/api/youtube/v3"
	"sort"
	"youtube_search_go_bot/keyboards"
	"youtube_search_go_bot/logging"
)

func ListSubscriptions(token *oauth2.Token, sorting keyboards.Sorting, resultLim uint16) ([]*youtube.Subscription, error) {
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
			if len(buf) <= int(resultLim) {
				buf = append(buf, sub)
			} else {
				break
			}
		}
		return nil
	})
	if err != nil {
		logging.LogFuncEnd("SearchSubsByTitle")
		return buf, errors.New("request.Pages() returned an error")
	}
	if len(buf) < 1 {
		logging.LogFuncEnd("SearchSubsByTitle")
		return buf, errors.New("no channel title matches the specified string")
	}

	switch sorting {
	case keyboards.SortingDate:
		sort.Slice(buf, func(i, j int) bool {
			return buf[i].Snippet.PublishedAt < buf[j].Snippet.PublishedAt
		})
	case keyboards.SortingAlphabetical:
		sort.Slice(buf, func(i, j int) bool {
			return buf[i].Snippet.Title < buf[j].Snippet.Title
		})
	}
	logging.LogFuncEnd("SearchSubsByTitle")
	return buf, nil
}

func ListPlaylists(token *oauth2.Token, sorting keyboards.Sorting, resultLim uint16) ([]*youtube.Playlist, error) {
	ctx := context.Background()
	youtubeService, err := newYouTubeService(token)
	if err != nil {
		return nil, err
	}

	part := []string{"contentDetails", "id", "snippet"}
	request := youtubeService.Playlists.List(part).Mine(true).MaxResults(50)

	buf := make([]*youtube.Playlist, 0)

	err = request.Pages(ctx, func(plr *youtube.PlaylistListResponse) error {
		for _, sub := range plr.Items {
			if len(buf) <= int(resultLim) {
				buf = append(buf, sub)
			} else {
				break
			}
		}
		return nil
	})
	if err != nil {
		logging.LogFuncEnd("SearchPlaylistsBy")
		return buf, errors.New("request.Pages() returned an error")
	}
	if len(buf) < 1 {
		logging.LogFuncEnd("SearchPlaylistsBy")
		return buf, errors.New("no channel title matches the specified string")
	}

	switch sorting {
	case keyboards.SortingDate:
		sort.Slice(buf, func(i, j int) bool {
			return buf[i].Snippet.PublishedAt < buf[j].Snippet.PublishedAt
		})
	case keyboards.SortingAlphabetical:
		sort.Slice(buf, func(i, j int) bool {
			return buf[i].Snippet.Title < buf[j].Snippet.Title
		})
	}
	logging.LogFuncEnd("SearchPlaylistsBy")
	return buf, nil
}
