package bot

import (
	"tgBotHakaton/configs"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var botPointer *tgbotapi.BotAPI = nil

func InitBot() error {
	config := configs.NewConfig()

	bot, err := tgbotapi.NewBotAPI(config.Telegram.BotApiToken)
	if err != nil {
		return err
	}
	bot.Debug = config.Telegram.Debug
	botPointer = bot

	u := tgbotapi.NewUpdate(0)
	u.Timeout = timeout

	updates := bot.GetUpdatesChan(u)
	handleUpdate(&updates)
	return nil
}
