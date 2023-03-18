package main

import (
	"log"
	"net/http"
	"os"
	"time"
	"youtube_search_go_bot/handlers"
	"youtube_search_go_bot/logging"

	"youtube_search_go_bot/commands"
	"youtube_search_go_bot/errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gopkg.in/telebot.v3"
)

func main() {
	logging.LogFuncStart("main")
	println("main function started")
	oldBot, e := tgbotapi.NewBotAPI(os.Getenv(""))
	errors.ExitOnError(e)

	setCommands := tgbotapi.NewSetMyCommands(
		tgbotapi.BotCommand{Command: string(commands.Start), Description: "Start the bot"},
		tgbotapi.BotCommand{Command: string(commands.Info), Description: "Show current status"},
		tgbotapi.BotCommand{Command: string(commands.Search), Description: "Search items"},
		tgbotapi.BotCommand{Command: string(commands.List), Description: "List items"},
		tgbotapi.BotCommand{Command: string(commands.LogOut), Description: "Log out"},
	)
	_, err := oldBot.Request(setCommands)
	errors.ExitOnError(err)

	pref := telebot.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	newBot, err := telebot.NewBot(pref)
	errors.ExitOnError(err)

	handlers.RegisterCommandHandlers(newBot)
	handlers.RegisterTextHandlers(newBot)
	handlers.RegisterCallbackHandlers(newBot)

	port, ok := os.LookupEnv("PORT")
	if !ok {
		log.Panicf("%v not set", port)
	}

	http.HandleFunc("/google_callback", handlers.GoogleCallbackHandler)
	go func() {
		err := http.ListenAndServe(":"+port, nil)
		if err != nil {
			errors.ExitOnError(err)
		}
	}()

	newBot.Start()
}
