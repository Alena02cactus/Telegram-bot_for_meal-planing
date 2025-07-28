package bot

import (
	"log"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type NutritionBot struct {
	bot *tgbotapi.BotAPI
	db  *sql.DB
}

func New(token string, db *sql.DB) (*NutritionBot, error) {
	botAPI, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &NutritionBot{
		bot: botAPI,
		db:  db,
	}, nil
}

func (b *NutritionBot) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		switch update.Message.Text {
		case "/start":
			b.handleStart(update)
		default:
			b.handleText(update)
		}
	}
}