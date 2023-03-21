package youtube_related

import (
	"golang.org/x/oauth2"
	"gopkg.in/telebot.v3"
	"youtube_search_go_bot/internal/dialogue"
)

func ExecuteSearchCmd(token *oauth2.Token, dialogueData dialogue.DialogueData) ([]interface{ SearchedItem }, error) {
	if dialogueData.ResultLimit < 1 || dialogueData.Target == "" || dialogueData.TextToSearch == "" || dialogueData.SearchIn == "" {
		return nil, telebot.Err("ExecuteSearchCmd : some parameters are not set")
	}
	var searchedItems []interface{ SearchedItem }
	switch dialogueData.Target {
	case dialogue.TargetSubscription:
		r, err := SearchSubscriptions(token, dialogueData.TextToSearch, dialogueData.SearchIn, dialogueData.ResultLimit)
		if err != nil {
			return searchedItems, err
		}
		for _, s := range r {
			ns := SubscriptionAlias(*s)
			searchedItems = append(searchedItems, &ns)
		}
		return searchedItems, nil
	case dialogue.TargetPlaylist:
		r, err := SearchPlaylists(token, dialogueData.TextToSearch, dialogueData.SearchIn, dialogueData.ResultLimit)
		if err != nil {
			return searchedItems, err
		}
		for _, p := range r {
			ns := PlaylistAlias(*p)
			searchedItems = append(searchedItems, &ns)
		}
		return searchedItems, nil
	default:
		return nil, telebot.Err("Target parameter is invalid")
	}
}

func ExecuteListCmd(token *oauth2.Token, dialogueData dialogue.DialogueData) ([]interface{ SearchedItem }, error) {
	if dialogueData.ResultLimit < 1 || dialogueData.Target == "" || dialogueData.Sorting == "" {
		return nil, telebot.Err("ExecuteListCmd : some parameters are not set")
	}
	var searchedItems []interface{ SearchedItem }
	switch dialogueData.Target {
	case dialogue.TargetSubscription:
		r, err := ListSubscriptions(token, dialogueData.Sorting, dialogueData.ResultLimit)
		if err != nil {
			return searchedItems, err
		}
		for _, s := range r {
			ns := SubscriptionAlias(*s)
			searchedItems = append(searchedItems, &ns)
		}
		return searchedItems, nil
	case dialogue.TargetPlaylist:
		r, err := ListPlaylists(token, dialogueData.Sorting, dialogueData.ResultLimit)
		if err != nil {
			return searchedItems, err
		}
		for _, p := range r {
			ns := PlaylistAlias(*p)
			searchedItems = append(searchedItems, &ns)
		}
		return searchedItems, nil
	default:
		return nil, telebot.Err("Target parameter is invalid")
	}
}
