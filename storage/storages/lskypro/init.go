package lskypro

import (
	"os"

	"github.com/imroc/req/v3"
	"github.com/krau/manyacg/storage/config"
	"github.com/krau/manyacg/storage/logger"
)

type StorageLskyPro struct{}

var apiURL string
var token string
var email string
var password string
var lskyProClient *req.Client

func InitLskyPro() {
	if isConfigEmpty() {
		logger.L.Fatalf("LskyPro config is empty")
		return
	}
	apiURL = config.Cfg.Storages.LskyPro.URL
	token = config.Cfg.Storages.LskyPro.Token
	email = config.Cfg.Storages.LskyPro.Email
	password = config.Cfg.Storages.LskyPro.Password
	lskyProClient = req.C().SetCommonHeader("Accept", "application/json")
	if token == "" {
		token = login()
		if token == "" {
			os.Exit(1)
		}
	}
	lskyProClient.SetCommonBearerAuthToken(token)
}
