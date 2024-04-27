package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func handleStart(msg *tgbotapi.Message) {
	if msg == nil {
		return // Тут нужно не забыть добавить обработку ошибок или хотя бы какое-то логгирование (когда я напишу пакет)
	}

	if msg.Chat.Type == "private" {
		sendMessage(int(msg.Chat.ID), startMessage)
		return
	}
}
