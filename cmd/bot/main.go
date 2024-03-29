package main

import (
	"log"
	"net/http"
	"os"
	"time"
	"youtube_search_go_bot/internal/commands"
	"youtube_search_go_bot/internal/db"
	"youtube_search_go_bot/internal/errors"
	handlers2 "youtube_search_go_bot/internal/handlers"
	"youtube_search_go_bot/internal/logging"

	//_ "github.com/gin-gonic/gin"
	//tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gopkg.in/telebot.v3"
)

func main() {
	log.SetFlags(log.LstdFlags)
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

	handlers2.RegisterCommandHandlers(bot)
	handlers2.RegisterTextHandlers(bot)
	handlers2.RegisterCallbackHandlers(bot)

	port, ok := os.LookupEnv("PORT")
	if !ok {
		log.Panicf("%v not set", port)
	}

	http.HandleFunc("/google_callback", handlers2.GoogleCallbackHandler)
	go func() {
		log.Println("server started")
		err := http.ListenAndServe(":"+port, nil)
		errors.ExitOnError(err)
	}()

	println("bot starting")
	bot.Start()
}
