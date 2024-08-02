package config

import (
	"log"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	configPath = "./.env"
)

type Config struct {
	TgBots   *BotsConfig
	Postgres *PostgresConfig
	Logger   *LoggerConfig
	Consumer *ConsumerConfig
}

type BotsConfig struct {
	MainToken    string `env:"MAIN_TOKEN"    env-required:"true"`
	LoggerToken  string `env:"LOGGER_TOKEN"  env-required:"true"`
	TelegramHost string `env:"TELEGRAM_HOST" env-required:"true"`
}

type PostgresConfig struct {
	User     string `env:"POSTGRES_USER"     env-required:"true"`
	Password string `env:"POSTGRES_PASSWORD" env-required:"true"`
	DB       string `env:"POSTGRES_DB"       env-required:"true"`
	Host     string `env:"POSTGRES_HOST"     env-required:"true"`
	Port     string `env:"POSTGRES_PORT"     env-required:"true"`
	SSLMode  string `env:"POSTGRES_SSL_MODE" env-required:"true"`
}

type LoggerConfig struct {
	AdminID int64  `env:"ADMIN_ID" env-required:"true"`
	Level   string `default:"info" env:"LOG_LEVEL"`
}

type ConsumerConfig struct {
	BatchSize int `env:"BATCH_SIZE" env-required:"true"`
}

type RedisConfig struct {
	Host       string        `env:"REDIS_HOST"       env-required:"true"`
	Port       string        `env:"REDIS_PORT"       env-required:"true"`
	Password   string        `env:"REDIS_PASSWORD"   env-required:"true"`
	DB         int           `env:"REDIS_DB"         env-required:"true"`
	Expiration time.Duration `env:"REDIS_EXPIRATION" env-required:"true"`
	Jitter     time.Duration `env:"REDIS_JITTER"     env-required:"true"`
}

type DictConfig struct {
	Path string `env:"DICTIONARY_PATH" env-required:"true"`
}

func MustLoad() *Config {
	config := &Config{}

	err := cleanenv.ReadConfig(configPath, config)
	if err != nil {
		log.Fatalf("Error while loading config: %s", err)
	}

	return config
}
