package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleUpdate(updates *tgbotapi.UpdatesChannel) {
	for update := range *updates {
		go func(update tgbotapi.Update) {
			if update.Message != nil {
				HandleMessage(update.Message)
			}
		}(update)
	}
}

func HandleMessage(msg *tgbotapi.Message) {
	switch msg.Text {
	default:
		go handleText(msg)
	case "/start":
		go handleStart(msg)
	}
}
