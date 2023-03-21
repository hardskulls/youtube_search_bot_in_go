package handlers

import (
	"gopkg.in/telebot.v3"
	"os"
	"strconv"
	"strings"
	"youtube_search_go_bot/db"
)

func RegisterTextHandlers(b *telebot.Bot) {
	b.Handle(telebot.OnText, func(c telebot.Context) error {
		return handleText(b, c)
	})
}

func handleText(b *telebot.Bot, c telebot.Context) error {
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		return telebot.Err("wrong db url")
	}
	dialogueData, err := db.GetDialogueData(strconv.FormatInt(c.Message().Sender.ID, 10), dbUrl)
	if err != nil {
		return err
	}

	text := c.Text()

	if strings.Contains(dialogueData.LastCallback, "ResultLimit") {
		num, err := strconv.ParseUint(text, 10, 16)
		if err != nil || num < 1 {
			b.Send(c.Chat(), "You must send some number greater than 1!")
			return err
		} else {
			dialogueData.ResultLimit = uint16(num)
			err := db.SaveDialogueData(strconv.FormatInt(c.Sender().ID, 10), dialogueData, dbUrl)
			if err != nil {
				return err
			}
			_, err = b.Send(c.Chat(), "Saved!")
			return err
		}
	} else if strings.Contains(dialogueData.LastCallback, "TextToSearch") {
		dialogueData.TextToSearch = c.Text()
		err := db.SaveDialogueData(strconv.FormatInt(c.Sender().ID, 10), dialogueData, dbUrl)
		if err != nil {
			return err
		}
		_, err = b.Send(c.Chat(), "Saved!")
		return err
	} else {
		_, err := b.Send(c.Chat(), "Try sending some commands first ⌨")
		return err
	}
}
