package config

import (
	"gopkg.in/simversity/gottp.v3/conf"
	"os"
)

type config struct {
	Gottp  conf.GottpSettings
	Khabar struct {
		DbUrl		     string
		DbName               string
		TranslationDirectory string
		Debug                bool
	}
}

func (self *config) MakeConfig(configPath string) {
	self.Gottp.Listen = "127.0.0.1:8911"

	if DbUrl := os.Getenv("MONGODB_URL"); DbUrl != "" {
		self.Khabar.DbUrl = DbUrl
	} else {
		self.Khabar.DbUrl = "mongodb://localhost/notifications_testing"
	}

	self.Khabar.DbName = "notifications_testing"

	if configPath != "" {
		conf.MakeConfig(configPath, self)
	}
}

func (self *config) GetGottpConfig() *conf.GottpSettings {
	return &self.Gottp
}

var Settings config
