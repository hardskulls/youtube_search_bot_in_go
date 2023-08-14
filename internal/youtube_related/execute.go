package youtube_related

import (
	"youtube_search_go_bot/internal/dialogue"
	"youtube_search_go_bot/internal/logging"

	"golang.org/x/oauth2"
	"gopkg.in/telebot.v3"
)

func ExecuteSearchCmd(
	token *oauth2.Token,
	dialogueData dialogue.DialogueData,
) (
	[]interface{ SearchedItem },
	error,
) {
	if dialogueData.ResultLimit < 1 ||
		dialogueData.Target == "" ||
		dialogueData.TextToSearch == "" ||
		dialogueData.SearchIn == "" {
		err := telebot.Err("ExecuteSearchCmd : some parameters are not set")
		logging.LogError(err)
		return nil, err
	}
	var searchedItems []interface{ SearchedItem }
	switch dialogueData.Target {
	case dialogue.TargetSubscription:
		results, err := SearchSubscriptions(
			token,
			dialogueData.TextToSearch,
			dialogueData.SearchIn,
			dialogueData.ResultLimit,
		)
		if err != nil {
			return searchedItems, err
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
			dialogueData.TextToSearch,
			dialogueData.SearchIn,
			dialogueData.ResultLimit,
		)
		if err != nil {
			return searchedItems, err
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
	dialogueData dialogue.DialogueData,
) (
	[]interface{ SearchedItem },
	error,
) {
	logging.LogFuncStart("ExecuteListCmd")
	if dialogueData.ResultLimit < 1 ||
		dialogueData.Target == "" ||
		dialogueData.Sorting == "" {
		err := telebot.Err("ExecuteListCmd : some parameters are not set")
		logging.LogError(err)
		return nil, err
	}
	var searchedItems []interface{ SearchedItem }
	switch dialogueData.Target {
	case dialogue.TargetSubscription:
		results, err := ListSubscriptions(token, dialogueData.Sorting, dialogueData.ResultLimit)
		if err != nil {
			return searchedItems, err
		}
		for _, s := range results {
			ns := SubscriptionAlias(*s)
			searchedItems = append(searchedItems, &ns)
		}
		logging.LogFuncEnd("ExecuteListCmd")
		return searchedItems, nil
	case dialogue.TargetPlaylist:
		results, err := ListPlaylists(token, dialogueData.Sorting, dialogueData.ResultLimit)
		if err != nil {
			return searchedItems, err
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
