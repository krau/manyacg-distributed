package lskypro

import (
	"github.com/krau/manyacg/storage/config"
	"github.com/krau/manyacg/storage/logger"
)

func isConfigEmpty() bool {
	if config.Cfg.Storages.LskyPro.URL == "" {
		return true
	}
	if config.Cfg.Storages.LskyPro.Token != "" {
		return false
	}
	return config.Cfg.Storages.LskyPro.Email == "" || config.Cfg.Storages.LskyPro.Password == ""
}

func login() string {
	logger.L.Debugf("Logging in...")
	var tokens tokensResp
	_, err := lskyProClient.R().SetBodyJsonMarshal(map[string]string{
		"email":    email,
		"password": password,
	}).
		SetSuccessResult(&tokens).
		Post(apiURL + "/tokens")
	if err != nil {
		logger.L.Fatalf("Error logging in: %v", err)
		return ""
	}
	if !tokens.Status {
		logger.L.Fatalf("Error logging in: %s", tokens.Message)
		return ""
	}
	logger.L.Debug(tokens.Message)
	return tokens.Data.Token
}
