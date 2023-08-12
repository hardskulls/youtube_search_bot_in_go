package handlers

import (
	"fmt"
	"os"
	"strconv"
	"youtube_search_go_bot/internal/db"
	"youtube_search_go_bot/internal/dialogue"
	"youtube_search_go_bot/internal/keyboards"
	"youtube_search_go_bot/internal/logging"
	"youtube_search_go_bot/internal/net"
	"youtube_search_go_bot/internal/utils"
	youtube_related2 "youtube_search_go_bot/internal/youtube_related"

	"gopkg.in/telebot.v3"
)

func RegisterCallbackHandlers(b *telebot.Bot) {
	b.Handle(telebot.OnCallback, func(c telebot.Context) error {
		logging.LogFuncStart("RegisterCallbackHandlers")
		err := telebot.Err("Something went wrong!")

		userId := c.Sender().ID
		callback := c.Callback().Data

		dbUrl := os.Getenv("DB_URL")
		if dbUrl == "" {
			return err
		}

		markup := *b.NewMarkup()

		switch callback {
		case string(keyboards.ListCancel), string(keyboards.ListSettings), string(keyboards.ListResultLimit),
			string(keyboards.ListTargetOptions), string(keyboards.ListSortingOptions):
			markup = keyboards.ListButton(callback).CreateKB()
		case string(keyboards.ListExecute):
			err := Execute(b, c)
			return err

		case string(keyboards.SearchCancel), string(keyboards.SearchSettings), string(keyboards.SearchResultLimit),
			string(keyboards.SearchTargetOptions), string(keyboards.SearchSearchInOptions):
			markup = keyboards.SearchButton(callback).CreateKB()
			logging.LogVar(markup, "markup")
		case string(keyboards.SearchExecute):
			err := Execute(b, c)
			return err

		case string(keyboards.SearchTargetSubscription), string(keyboards.SearchTargetPlaylist),
			string(keyboards.ListTargetSubscription), string(keyboards.ListTargetPlaylist):
			err := db.SaveTarget(strconv.FormatInt(userId, 10), callback, dbUrl)
			if err != nil {
				return err
			}
			markup = keyboards.SearchSettings.CreateKB()

		case string(keyboards.SearchInTitle), string(keyboards.SearchInDescription):
			err := db.SaveSearchIn(strconv.FormatInt(userId, 10), keyboards.SearchIn(callback), dbUrl)
			if err != nil {
				return err
			}
			markup = keyboards.SearchSettings.CreateKB()

		case string(keyboards.SortingDate), string(keyboards.SortingAlphabetical):
			err := db.SaveSorting(strconv.FormatInt(userId, 10), keyboards.Sorting(callback), dbUrl)
			if err != nil {
				return err
			}
			markup = keyboards.SearchSettings.CreateKB()
		}

		_, err = b.Edit(c.Callback(), &markup)
		if err != nil {
			logging.LogError(err)
		}
		logging.LogFuncEnd("RegisterCallbackHandlers")
		return err
	})
}

// Execute 'Search' or 'List' command.
func Execute(b *telebot.Bot, c telebot.Context) error {
	dbUrl := os.Getenv("DB_URL")
	dialogueData, err := db.GetDialogueData(strconv.FormatInt(c.Sender().ID, 10), dbUrl)
	if err != nil {
		return err
	}
	token, err := db.GetOAuthToken(strconv.FormatInt(c.Sender().ID, 10), dbUrl)
	if err != nil {
		htmlizer := utils.HTMLizer{}
		text := fmt.Sprintf(
			"To %v with your account, follow this link: \r%v",
			htmlizer.StrongBold("log in"),
			htmlizer.InlineURL(htmlizer.Bold("Log In"), net.LogInUrl(c.Sender().ID)),
		)
		_, _ = b.Send(c.Chat(), text, telebot.ModeHTML)
		return err
	}

	var searchableItems []interface{ youtube_related2.SearchedItem }

	switch dialogueData.ActiveCmd {
	case dialogue.SearchCommand:
		r, err := youtube_related2.ExecuteSearchCmd(token, dialogueData)
		if err != nil {
			text := fmt.Sprintf("You need to specify missing parameters for the "+
				"command: \nResult limit - %v \nTarget - %v \nText to search - %v \nSearch in - %v",
				dialogueData.ResultLimit, dialogueData.Target, dialogueData.TextToSearch, dialogueData.SearchIn,
			)
			_, err := b.Send(c.Chat(), text)
			return err
		} else {
			searchableItems = r
		}
	case dialogue.ListCommand:
		r, err := youtube_related2.ExecuteListCmd(token, dialogueData)
		if err != nil {
			text := fmt.Sprintf("You need to specify missing parameters for the "+
				"command: \nResult limit - %v \nTarget - %v \nSorting - %v",
				dialogueData.ResultLimit, dialogueData.Target, dialogueData.Sorting,
			)
			_, err := b.Send(c.Chat(), text)
			return err
		} else {
			searchableItems = r
		}
	}

	for _, item := range searchableItems {
		text := fmt.Sprintf("<b>%v</b>, \n\n%v, \n\n%v, \n\n%v", item.Title(), item.Description(), item.Date(), item.Link())
		_, _ = b.Send(c.Chat(), text, telebot.ModeHTML)
	}

	_, _ = b.Send(c.Chat(), fmt.Sprintf("Found %v results", len(searchableItems)))

	return nil
}
