package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func sendMessage(chatid int, text string) {
	msg := tgbotapi.NewMessage(int64(chatid), text)
	msg.ParseMode = "html"
	botPointer.Send(msg)
}
