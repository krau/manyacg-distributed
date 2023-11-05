package storages

import (
	"github.com/krau/manyacg/storage/config"
	"github.com/krau/manyacg/storage/storages/local"
	"github.com/krau/manyacg/storage/storages/lskypro"
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
	if config.Cfg.Storages.LskyPro.Enable {
		lskypro.InitLskyPro()
		Storages["lsky_pro"] = new(lskypro.StorageLskyPro)
	}
}
