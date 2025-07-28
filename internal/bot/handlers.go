package bot  
import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"  
)

func (b *NutritionBot) handleStart(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(
		update.Message.Chat.ID, 
		"Привет! Я помогу тебе планировать рацион. Введи свой вес и рост в формате: 70 175",  
	)

	_, err := b.bot.Send(msg)
	if err != nil {
		b.logger.Error("Ошибка отправки сообщения", zap.Error(err))
	}
}

func (b *NutritionBot) handleText(update tgbotapi.Update) {
	if update.Message.Text == "" {
		return
	}

	_, err := b.db.Exec(
		"INSERT INTO users (chat_id, weight, height) VALUES (?, ?, ?)",
		update.Message.Chat.ID,
		70,  
		175, 
	)
	if err != nil {
		b.logger.Error("Ошибка сохранения в БД", zap.Error(err))
	}

	reply := tgbotapi.NewMessage(
		update.Message.Chat.ID,
		"Данные сохранены! Используй /plan для добавления блюд.",
	)

	reply.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/plan"),
			tgbotapi.NewKeyboardButton("/help"),
		),
	)

	_, err = b.bot.Send(reply)
	if err != nil {
		b.logger.Error("Ошибка отправки сообщения", zap.Error(err))
	}
}

// Формат: /plan день_недели блюдо
// Пример: /plan Monday Овсянка с фруктами
func (b *NutritionBot) handlePlan(update tgbotapi.Update) {
	// Разбиваем сообщение на части: ["/plan", "Monday", "Овсянка..."]
	args := strings.SplitN(update.Message.Text, " ", 3)
	if len(args) < 3 {
		msg := tgbotapi.NewMessage(
			update.Message.Chat.ID,
			"❌ Неверный формат. Используй: /plan день_недели блюдо\nПример: /plan Monday Овсянка с фруктами",
		)
		b.bot.Send(msg)
		return
	}

	day := args[1]    // "Monday"
	meal := args[2]   // "Овсянка с фруктами"

	// Сохраняем в БД
	_, err := b.db.Exec(
		"INSERT INTO meals (user_id, day, meal) VALUES (?, ?, ?)",
		update.Message.Chat.ID,
		day,
		meal,
	)
	if err != nil {
		b.logger.Error("Ошибка сохранения блюда", zap.Error(err))
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "🚫 Ошибка при сохранении. Попробуй снова.")
		b.bot.Send(msg)
		return
	}

	// Отправляем подтверждение
	msg := tgbotapi.NewMessage(
		update.Message.Chat.ID,
		"✅ Блюдо добавлено!\n" +
		"День: " + day + "\n" +
		"Блюдо: " + meal,
	)
	b.bot.Send(msg)
}