package config

import (
	"github.com/caarlos0/env/v11"

	"pkg/database/postgresql"
	"pkg/errors"
	"pkg/trace"
)

// Config - общая структура конфига
type Config struct {

	// Адрес для http-сервера
	HTTP string `env:"LISTEN_HTTP"`

	// Данные базы данных
	Repository postgresql.PostgreSQLConfig
	DBName     string `env:"DB_NAME"`

	Tracer trace.TracerConfig

	// Доступы к телеграм-боту
	Telegram struct {
		Enabled bool   `env:"TG_BOT_ENABLED" envDefault:"false"`
		Token   string `env:"TG_BOT_TOKEN" envDefault:"secret"`
	}

	ServiceName string `env:"SERVICE_NAME"`
}

// GetConfig возвращает конфигурацию из .env файла
func GetConfig() (config Config, err error) {
	if err = env.Parse(&config); err != nil {
		return config, errors.InternalServer.Wrap(err)
	}
	return config, nil
}
