package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_EDITOR_TELEGRAM_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// Loop through each update.
	for update := range updates {
		// Check if we've gotten a message update.
		if update.Message != nil {
			// Construct a new message from the given chat ID and containing
			// the text that we received.
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ParseMode = "MarkdownV2"

			text, markup := getFilesListMessage()
			msg.ReplyMarkup = markup

			msg.Text = text

			// Send the message.
			if _, err = bot.Send(msg); err != nil {
				panic(err)
			}
		} else if update.CallbackQuery != nil {
			// Respond to the callback query, telling Telegram to show the user
			// a message with the data received.
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				panic(err)
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
			upd.ParseMode = "MarkdownV2"
			if _, err := bot.Request(upd); err != nil {
				// TODO: из-за ошибки обновления файла
				log.Printf("Error while processing text %s", err)
			}

		}
	}
}
