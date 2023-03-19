package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"youtube_search_go_bot/commands"
	"youtube_search_go_bot/db"
	"youtube_search_go_bot/errors"
	"youtube_search_go_bot/handlers"
	"youtube_search_go_bot/logging"

	//_ "github.com/gin-gonic/gin"
	//tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gopkg.in/telebot.v3"
)

func main() {
	logging.LogFuncStart("main")
	println("main function started")

	dbUrl, ok := os.LookupEnv("DB_URL")
	if !ok {
		log.Panicf("%v not set", dbUrl)
	}
	err := db.CreateTables(dbUrl)
	errors.ExitOnError(err)

	botToken, ok := os.LookupEnv("TELEGRAM_BOT_TOKEN")
	if !ok {
		log.Panicf("%v not set", botToken)
	}

	pref := telebot.Settings{
		Token:  botToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := telebot.NewBot(pref)
	errors.ExitOnError(err)

	err = bot.RemoveWebhook(true)
	errors.ExitOnError(err)

	botCommands := []telebot.Command{
		{Text: string(commands.Start), Description: "Start the bot"},
		{Text: string(commands.Info), Description: "Show current status"},
		{Text: string(commands.Search), Description: "Search items"},
		{Text: string(commands.List), Description: "List items"},
		{Text: string(commands.LogOut), Description: "Log out"},
	}

	err = bot.SetCommands(botCommands)
	errors.ExitOnError(err)

	handlers.RegisterCommandHandlers(bot)
	handlers.RegisterTextHandlers(bot)
	handlers.RegisterCallbackHandlers(bot)

	port, ok := os.LookupEnv("PORT")
	if !ok {
		log.Panicf("%v not set", port)
	}

	http.HandleFunc("/google_callback", handlers.GoogleCallbackHandler)
	go func() {
		println("server started")
		err := http.ListenAndServe(":"+port, nil)
		if err != nil {
			errors.ExitOnError(err)
		}
	}()

	println("bot starting")
	bot.Start()
}
