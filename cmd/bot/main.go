package main

import (
	"log"
	"github.com/Alena02cactus/Telegram-bot_for_meal-planing/internal/bot"
	"github.com/Alena02cactus/Telegram-bot_for_meal-planing/internal/config"
	"github.com/Alena02cactus/Telegram-bot_for_meal-planing/internal/database"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки конфига:", err)
	}

	db, err := database.NewPostgresDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}
	defer db.Close()

	nutritionBot, err := bot.New(cfg.TelegramToken, db)
	if err != nil {
		log.Fatal("Ошибка создания бота:", err)
	}

	log.Println("Бот запущен!")
	nutritionBot.Start()
}