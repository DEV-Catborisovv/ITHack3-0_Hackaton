package bot

// Static configs
const timeout = 60

// Static Messages
const startMessage = "👋 Здравствуйте! Этот телеграм бот разработан для проверки оплаты через СБП (только для продавцов).\n\n❗️Чтобы проверить статус оплаты отправьте уникальный идентификатор QR в этот чат"
const errorWithTryGetData = "❌ Возникла ошибка! Нам не удалось получить информацию об этой транзакции. Попробуйте снова..."

// States
const stateError = "Ошибка"
const stateNew = "Новый"
const stateProgress = "Операция выполняется"
const stateCanceled = "Операция отменена"
const stateExpired = "Время выполнения операции истекло"
const statePaid = "Оплачено"
