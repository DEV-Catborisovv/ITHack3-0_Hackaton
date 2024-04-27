package app

import (
	"fmt"
	"log"
	"tgBotHakaton/internal/app/bot"
)

func InitApp() {
	err := bot.InitBot()
	if err != nil {
		log.Panic(fmt.Sprintf("Error with listen telegramBot :\n%s\n", err))
		return
	}
}
