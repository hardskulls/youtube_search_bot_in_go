package handlers

import (
	"fmt"
	"os"
	"strconv"
	"youtube_search_go_bot/internal/commands"
	"youtube_search_go_bot/internal/db"
	"youtube_search_go_bot/internal/dialogue"
	"youtube_search_go_bot/internal/keyboards"
	"youtube_search_go_bot/internal/logging"

	_ "github.com/go-telegram/bot"
	"gopkg.in/telebot.v3"
)

func RegisterCommandHandlers(b *telebot.Bot) {
	logging.LogFuncStart("RegisterCommandHandlers")
	b.Handle(string(commands.Start), func(c telebot.Context) error {
		return c.Send("This bot lets you search stuff on your YouTube channel")
	})
	b.Handle(string(commands.Info), func(c telebot.Context) error {
		err := handleInfoCmd(b, c)
		logging.LogError(err)
		return err
	})
	b.Handle(string(commands.Search), func(c telebot.Context) error {
		err := handleSearchCmd(b, c)
		logging.LogError(err)
		return err
	})
	b.Handle(string(commands.List), func(c telebot.Context) error {
		err := handleListCmd(b, c)
		logging.LogError(err)
		return err
	})
	b.Handle(string(commands.LogOut), func(c telebot.Context) error {
		err := handleLogOutCmd(c)
		logging.LogError(err)
		return err
	})
	logging.LogFuncEnd("RegisterCommandHandlers")
}

func handleSearchCmd(b *telebot.Bot, c telebot.Context) error {
	logging.LogFuncStart("handleSearchCmd")
	kb := keyboards.SearchSettings.CreateKB()
	msg, err := b.Send(c.Chat(), "Search parameters ⚙", &kb)
	if err != nil {
		logging.LogVar("b.Send error", "b.Send error")
		return err
	}

	dialogueData := dialogue.DialogueData{ActiveCmd: dialogue.SearchCommand, MsgWithCallbackId: msg.ID}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		logging.LogVar("db error", "db error")
		return telebot.Err("wrong db url")
	}
	err = db.SaveDialogueData(strconv.FormatInt(c.Update().Message.Sender.ID, 10), dialogueData, dbUrl)
	if err != nil {
		logging.LogVar("SaveDialogueData error", "SaveDialogueData error")
		return err
	}
	logging.LogFuncEnd("handleSearchCmd")

	return nil
}

func handleListCmd(b *telebot.Bot, c telebot.Context) error {
	logging.LogFuncStart("RegisterCommandHandlers")
	kb := keyboards.ListSettings.CreateKB()
	msg, err := b.Send(c.Chat(), "List parameters ⚙", &kb)
	if err != nil {
		return err
	}

	dialogueData := dialogue.DialogueData{ActiveCmd: dialogue.ListCommand, MsgWithCallbackId: msg.ID}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		return telebot.Err("wrong db url")
	}
	err = db.SaveDialogueData(strconv.FormatInt(c.Update().Message.Sender.ID, 10), dialogueData, dbUrl)
	if err != nil {
		return err
	}
	logging.LogFuncEnd("RegisterCommandHandlers")

	return nil
}

func handleInfoCmd(b *telebot.Bot, c telebot.Context) error {
	logging.LogFuncStart("handleInfoCmd")
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		logging.LogVar("dbUrl error", "dbUrl error")
		return telebot.Err("wrong db url")
	}
	dData, err := db.GetDialogueData(strconv.FormatInt(c.Message().Sender.ID, 10), dbUrl)
	if err != nil {
		logging.LogVar(dData, "dData")
		logging.LogError(err)
	}

	text := fmt.Sprintf(
		"Cureent state: \n⌨ Active command : %v \n🎯 Target : %v \n🧮 Result limit : %v \n🏷 Search in : %v "+
			"\n🗂 Sorting : %v \n💬 Text to search : %v",
		dData.ActiveCmd, dData.Target, dData.ResultLimit, dData.SearchIn, dData.Sorting, dData.TextToSearch,
	)

	_, err = b.Send(c.Chat(), text)
	if err != nil {
		logging.LogVar(err, "err is")
		return err
	}
	logging.LogFuncEnd("handleInfoCmd")

	return nil
}

func handleLogOutCmd(c telebot.Context) error {
	logging.LogFuncStart("RegisterCommandHandlers")
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		return telebot.Err("wrong db url")
	}
	err := db.LogOut(strconv.FormatInt(c.Sender().ID, 10), dbUrl)
	if err != nil {
		_ = c.Send("Log out command failed ❌")
	} else {
		_ = c.Send("Logged out successfully ✅")
	}
	logging.LogFuncEnd("RegisterCommandHandlers")
	return nil
}
