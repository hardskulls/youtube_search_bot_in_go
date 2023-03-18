package handlers

import (
	"fmt"
	_ "github.com/go-telegram/bot"
	"gopkg.in/telebot.v3"
	"os"
	"strconv"
	"youtube_search_go_bot/commands"
	"youtube_search_go_bot/db"
	"youtube_search_go_bot/dialogue"
	"youtube_search_go_bot/keyboards"
)

func RegisterCommandHandlers(b *telebot.Bot) {
	b.Handle(string(commands.Start), func(c telebot.Context) error {
		return c.Send("This bot lets you search stuff on your YouTube channel")
	})
	b.Handle(string(commands.Info), func(c telebot.Context) error {
		return handleInfoCmd(b, c)
	})
	b.Handle(string(commands.Search), func(c telebot.Context) error {
		return handleSearchCmd(b, c)
	})
	b.Handle(string(commands.List), func(c telebot.Context) error {
		return handleListCmd(b, c)
	})
	b.Handle(string(commands.LogOut), func(c telebot.Context) error {
		return handleLogOutCmd(c)
	})
}

func handleSearchCmd(b *telebot.Bot, c telebot.Context) error {
	msg, err := b.Send(c.Chat(), "Search parameters ⚙", keyboards.SearchSettings.CreateKB())
	if err != nil {
		return err
	}

	dialogueData := dialogue.DialogueData{ActiveCmd: dialogue.SearchCommand, MsgWithCallbackId: msg.ID}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		return telebot.Err("wrong db url")
	}
	err = db.SaveDialogueData(strconv.FormatInt(c.Update().Message.Sender.ID, 10), dialogueData, dbUrl)
	if err != nil {
		return err
	}

	return nil
}

func handleListCmd(b *telebot.Bot, c telebot.Context) error {
	msg, err := b.Send(c.Chat(), "List parameters ⚙", keyboards.ListSettings.CreateKB())
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

	return nil
}

func handleInfoCmd(b *telebot.Bot, c telebot.Context) error {
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		return telebot.Err("wrong db url")
	}
	dialogueData, err := db.GetDialogueData(strconv.FormatInt(c.Message().Sender.ID, 10), dbUrl)
	if err != nil {
		return err
	}

	text := fmt.Sprintf(
		"Cureent state: \n Active command - %v \nTarget - %v \nResult limit - %v \nSearch in - %v \nSorting - %v",
		dialogueData.ActiveCmd, dialogueData.Target, dialogueData.ResultLimit, dialogueData.SearchIn, dialogueData.Sorting,
	)

	_, err = b.Send(c.Chat(), text)
	if err != nil {
		return err
	}

	return nil
}

func handleLogOutCmd(c telebot.Context) error {
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		return telebot.Err("wrong db url")
	}
	err := db.LogOut(strconv.FormatInt(c.Sender().ID, 10), dbUrl)
	if err != nil {
		_ = c.Send("Log out command failed ")
	} else {
		_ = c.Send("Logged out successfully!")
	}
	return nil
}
