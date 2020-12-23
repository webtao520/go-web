package config

import _"github.com/nicksnyder/go-i18n/v2/i18n"

//  系统配置
type App struct {
	Address  string
	Static   string
	Log      string
	Locale   string
	Language string
}

type Database struct {
	Driver   string
	Address  string
	Database string
	User     string
	Password string
}

type Configuration struct {
	App          App
	Db           Database
}

