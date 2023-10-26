package storages

import (
	"github.com/krau/manyacg/storage/config"
	"github.com/krau/manyacg/storage/storages/local"
	"github.com/krau/manyacg/storage/storages/telegram"
)

var Storages = make(map[string]Storage)

func InitStorages() {
	if config.Cfg.Storages.Local.Enable {
		local.InitLocal()
		Storages["local"] = new(local.StorageLocal)
	}
	if config.Cfg.Storages.Telegram.Enable {
		telegram.InitTelegram()
		Storages["telegram"] = new(telegram.StorageTelegram)
	}
}
