package main

import (
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	flag "github.com/spf13/pflag"
)

func parseTelegramUsers(usersLoginsList string) map[string]struct{} {
	allowedUsers := make(map[string]struct{})
	for _, login := range strings.Split(usersLoginsList, ";") {
		if login == "" {
			continue
		}
		allowedUsers[login] = struct{}{}
	}
	return allowedUsers
}

func main() {
	var token string
	var parseMode string
	var allowedUsersList string
	var workdirPath string
	var enableDebug bool

	flag.StringVar(&token, "token", "", "telegramm token")
	flag.StringVar(&parseMode, "mode", "HTML", "telegram parse mode")
	flag.StringVar(&allowedUsersList, "users", "", "users who are allowed to use this bot separated by semicolumns")
	flag.StringVar(&workdirPath, "path", ".", "working folder for the bot")
	flag.BoolVar(&enableDebug, "debug", false, "enable debug")

	flag.Parse()

	err := os.Chdir(workdirPath)
	if err != nil {
		log.Fatalf("Failed to use %s as working dir %s", workdirPath, err)
	}

	if currentdir, err := os.Getwd(); err == nil {
		log.Printf("Starting bot in %s directory", currentdir)
	} else {
		log.Fatalf("Failed to get working directory %s", err)
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("Failed to creat bot %s", err)
	}

	allowedUsers := parseTelegramUsers(allowedUsersList)

	if len(allowedUsers) == 0 {
		log.Fatalf("There are no users allowed to use this bot. Halting.")
	}

	bot.Debug = enableDebug

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

UPDATES_LOOP:
	for update := range updates {
		login := update.SentFrom().UserName
		log.Printf("Got request from %s", login)
		if _, ok := allowedUsers[update.SentFrom().UserName]; !ok {
			log.Printf("Login %s was not found in white list, continue", login)
			continue UPDATES_LOOP
		}

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

			if len(dt) == 0 {
				log.Printf("No command found, continue")
				continue UPDATES_LOOP
			}

			switch dt[0] {
			case 'f':
				text, markup = getFileMessage(dt)
			case 'c':
				text, markup = getFilesListMessage()
			default:
				log.Printf("Unknown command")
				continue UPDATES_LOOP
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
