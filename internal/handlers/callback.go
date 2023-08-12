package handlers

import (
	"errors"
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

		m, err := b.Edit(c.Callback(), &markup)
		logging.LogError(err)
		logging.LogVar(markup, "markup")
		logging.LogVar(m, "m")
		logging.LogFuncEnd("RegisterCallbackHandlers")
		return err
	})
}

// Execute 'Search' or 'List' command.
func Execute(b *telebot.Bot, c telebot.Context) error {
	logging.LogFuncStart("Execute")
	dbUrl := os.Getenv("DB_URL")
	dialogueData, err := db.GetDialogueData(strconv.FormatInt(c.Sender().ID, 10), dbUrl)
	if err != nil {
		return err
	}
	token, err := db.GetOAuthToken(strconv.FormatInt(c.Sender().ID, 10), dbUrl)
	if err != nil {
		sendLoginMsg(c, b)
		return err
	}

	var searchableItems []interface{ youtube_related2.SearchedItem }

	switch dialogueData.ActiveCmd {
	case dialogue.SearchCommand:
		searchableItems, err = youtube_related2.ExecuteSearchCmd(token, dialogueData)
		if err != nil {
			_, e := b.Send(c.Chat(), sendMissingSearchParams(dialogueData))
			logging.LogError(e)
			return err
		}
	case dialogue.ListCommand:
		searchableItems, err = youtube_related2.ExecuteListCmd(token, dialogueData)
		if err != nil {
			_, e := b.Send(c.Chat(), sendMissingListParams(dialogueData))
			logging.LogError(e)
			return err
		}
	default:
		err := errors.New("hit default case in switch statement")
		logging.LogError(err)
		return err
	}

	for _, item := range searchableItems {
		text := formatResult(item)
		_, _ = b.Send(c.Chat(), text, telebot.ModeHTML)
	}

	_, _ = b.Send(c.Chat(), fmt.Sprintf("Found %v results", len(searchableItems)))
	logging.LogFuncEnd("Execute")

	return nil
}

func formatResult(item interface{youtube_related2.SearchedItem}) string {
	text := fmt.Sprintf("<b>%v</b>, \n\n%v, \n\n%v, \n\n%v", item.Title(), item.Description(), item.Date(), item.Link())
	return text
}

func sendMissingListParams(dialogueData dialogue.DialogueData) string {
	text := fmt.Sprintf("You need to specify missing parameters for the "+
		"command: \nResult limit - %v \nTarget - %v \nSorting - %v",
		dialogueData.ResultLimit, dialogueData.Target, dialogueData.Sorting,
	)
	return text
}

func sendMissingSearchParams(dialogueData dialogue.DialogueData) string {
	text := fmt.Sprintf("You need to specify missing parameters for the "+
		"command: \nResult limit - %v \nTarget - %v \nText to search - %v \nSearch in - %v",
		dialogueData.ResultLimit, dialogueData.Target, dialogueData.TextToSearch, dialogueData.SearchIn,
	)
	return text
}

func sendLoginMsg(c telebot.Context, b *telebot.Bot) {
	htmlizer := utils.HTMLizer{}

	logIn := htmlizer.StrongBold("log in")
	link := htmlizer.InlineURL(htmlizer.Bold("Log In"), net.LogInUrl(c.Sender().ID))

	text := fmt.Sprintf("To %v with your account, follow this link: \r%v", logIn, link)
	_, err := b.Send(c.Chat(), text, telebot.ModeHTML)
	logging.LogError(err)
}
