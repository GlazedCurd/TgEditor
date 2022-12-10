package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	flag "github.com/spf13/pflag"
)

func main() {
	var token string
	var parseMode string
	var enableDebug bool

	flag.StringVar(&token, "token", "", "telegramm token")
	flag.StringVar(&parseMode, "mode", "HTML", "telegram parse mode")
	flag.BoolVar(&enableDebug, "debug", false, "enable debug")

	flag.Parse()

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = enableDebug

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {

		// For any simple message - render files list
		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ParseMode = parseMode

			text, markup := getFilesListMessage()
			msg.ReplyMarkup = markup

			msg.Text = text

			if _, err = bot.Send(msg); err != nil {
				log.Printf("Error while sending message %s", err)
			}
		} else if update.CallbackQuery != nil {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				log.Printf("Error while processing callback %s", err)
			}

			dt := update.CallbackQuery.Data

			var text string
			var markup tgbotapi.InlineKeyboardMarkup

			switch dt[0] {
			case 'f':
				text, markup = getFileMessage(dt)
			default:
				text, markup = getFilesListMessage()
			}

			upd := tgbotapi.NewEditMessageTextAndMarkup(
				update.CallbackQuery.Message.Chat.ID,
				update.CallbackQuery.Message.MessageID,
				text,
				markup,
			)
			upd.ParseMode = parseMode
			if _, err := bot.Request(upd); err != nil {
				// TODO: add handling of "no changes error"
				log.Printf("Error while processing text %s", err)
			}

		}
	}
}
