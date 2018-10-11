package config

import ()

type Configuration struct {
	Server ServerConfiguration
	Database DatabaseConfiguration
	Telegram TelegramConfiguration 
	SqlDatabase SqlDatabaseConfiguration
}
