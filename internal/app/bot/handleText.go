package bot

import (
	"fmt"
	raifapiclient_methods "tgBotHakaton/internal/app/raif-api-client/methods"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleText(msg *tgbotapi.Message) {
	if msg.Chat.Type == "private" {
		err, resp := raifapiclient_methods.GetOrderDataFromUid(msg.Text)
		if err != nil {
			sendMessage(int(msg.Chat.ID), errorWithTryGetData)
			return
		}

		var state string = ""
		switch resp.QrStatus {
		default:
			state = stateError
		case "NEW":
			state = stateNew
		case "IN_PROGRESS":
			state = stateProgress
		case "CANCELLED":
			state = stateCanceled
		case "EXPIRED":
			state = stateExpired
		case "PAID":
			state = statePaid
		}
		sendMessage(int(msg.Chat.ID), fmt.Sprintf("Текущее состояние: %s", state))
		return
	}
}
