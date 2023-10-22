package telegram

import (
	"github.com/krau/manyacg/storage/config"
	"github.com/krau/manyacg/storage/logger"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

type StorageTelegram struct{}

var bot *telego.Bot
var chatID telego.ChatID

func init() {
	b, err := telego.NewBot(config.Cfg.Storages.Telegram.Token, telego.WithDefaultLogger(false, true))
	if err != nil {
		logger.L.Fatalf("Error creating bot: %v", err)
		return
	}
	bot = b
	if config.Cfg.Storages.Telegram.Username != "" {
		chatID = tu.Username(config.Cfg.Storages.Telegram.Username)
	} else {
		chatID = tu.ID(config.Cfg.Storages.Telegram.ChatId)
	}
}
