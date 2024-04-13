package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/ilhomsoliev/pocket-telegram-bot/pkg/config"
	"github.com/ilhomsoliev/pocket-telegram-bot/pkg/storage"
	"github.com/zhashkevych/go-pocket-sdk"
)

type Bot struct {
	bot         *tgbotapi.BotAPI
	client      *pocket.Client
	redirectURL string

	storage storage.TokenStorage

	messages config.Messages
}

func NewBot(bot *tgbotapi.BotAPI, client *pocket.Client, redirectURL string, storage storage.TokenStorage, messages config.Messages) *Bot {
	return &Bot{
		bot:         bot,
		client:      client,
		redirectURL: redirectURL,
		storage:     storage,
		messages:    messages,
	}
}

func (b *Bot) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		// Handle commands
		if update.Message.IsCommand() {
			if err := b.handleCommand(update.Message); err != nil {
				b.handleError(update.Message.Chat.ID, err)
			}

			continue
		}

		// Handle regular messages
		if err := b.handleMessage(update.Message); err != nil {
			b.handleError(update.Message.Chat.ID, err)
		}
	}

	return nil
}
