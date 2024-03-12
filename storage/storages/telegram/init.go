package telegram

import (
	"github.com/krau/manyacg/storage/config"
	"github.com/krau/manyacg/storage/logger"
	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegoapi"
	tu "github.com/mymmrac/telego/telegoutil"
)

type StorageTelegram struct{}

var bot *telego.Bot
var chatID telego.ChatID

func InitTelegram() {
	b, err := telego.NewBot(config.Cfg.Storages.Telegram.Token, telego.WithDefaultLogger(false, true),
		telego.WithAPICaller(&telegoapi.RetryCaller{
			Caller:       telegoapi.DefaultFastHTTPCaller,
			MaxAttempts:  4,
			ExponentBase: 2,
			StartDelay:   10,
		}))
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
