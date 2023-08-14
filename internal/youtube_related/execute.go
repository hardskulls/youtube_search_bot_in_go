package youtube_related

import (
	"youtube_search_go_bot/internal/dialogue"
	"youtube_search_go_bot/internal/logging"

	"golang.org/x/oauth2"
	"gopkg.in/telebot.v3"
)

func ExecuteSearchCmd(
	token *oauth2.Token,
	dData dialogue.DialogueData,
) (
	[]interface{ SearchedItem },
	error,
) {
	if dData.ResultLimit < 1 ||
		dData.Target == "" ||
		dData.TextToSearch == "" ||
		dData.SearchIn == "" {
		err := telebot.Err("ExecuteSearchCmd : some parameters are not set")
		logging.LogError(err)
		return nil, err
	}
	var searchedItems []interface{ SearchedItem }
	switch dData.Target {
	case dialogue.TargetSubscription:
		results, err := SearchSubscriptions(
			token,
			dData.TextToSearch,
			dData.SearchIn,
			dData.ResultLimit,
		)
		if err != nil {
			logging.LogError(err)
			return searchedItems, nil
		}
		for _, s := range results {
			ns := SubscriptionAlias(*s)
			searchedItems = append(searchedItems, &ns)
		}
		logging.LogFuncEnd("ExecuteSearchCmd")
		return searchedItems, nil
	case dialogue.TargetPlaylist:
		results, err := SearchPlaylists(
			token,
			dData.TextToSearch,
			dData.SearchIn,
			dData.ResultLimit,
		)
		if err != nil {
			logging.LogError(err)
			return searchedItems, nil
		}
		for _, p := range results {
			ns := PlaylistAlias(*p)
			searchedItems = append(searchedItems, &ns)
		}
		logging.LogFuncEnd("ExecuteSearchCmd")
		return searchedItems, nil
	default:
		return nil, telebot.Err("Target parameter is invalid")
	}
}

func ExecuteListCmd(
	token *oauth2.Token,
	dData dialogue.DialogueData,
) (
	[]interface{ SearchedItem },
	error,
) {
	logging.LogFuncStart("ExecuteListCmd")
	if dData.ResultLimit < 1 ||
		dData.Target == "" ||
		dData.Sorting == "" {
		err := telebot.Err("ExecuteListCmd : some parameters are not set")
		logging.LogError(err)
		return nil, err
	}
	var searchedItems []interface{ SearchedItem }
	switch dData.Target {
	case dialogue.TargetSubscription:
		results, err := ListSubscriptions(token, dData.Sorting, dData.ResultLimit)
		if err != nil {
			logging.LogError(err)
			return searchedItems, nil
		}
		for _, s := range results {
			ns := SubscriptionAlias(*s)
			searchedItems = append(searchedItems, &ns)
		}
		logging.LogFuncEnd("ExecuteListCmd")
		return searchedItems, nil
	case dialogue.TargetPlaylist:
		results, err := ListPlaylists(token, dData.Sorting, dData.ResultLimit)
		if err != nil {
			logging.LogError(err)
			return searchedItems, nil
		}
		for _, p := range results {
			ns := PlaylistAlias(*p)
			searchedItems = append(searchedItems, &ns)
		}
		logging.LogFuncEnd("ExecuteListCmd")
		return searchedItems, nil
	default:
		return nil, telebot.Err("Target parameter is invalid")
	}
}
