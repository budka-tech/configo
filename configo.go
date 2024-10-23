package configo

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Database struct {
	Type          string        `yaml:"type" env-required:"true"`
	Host          string        `yaml:"host" env-required:"true"`
	Port          int           `yaml:"port" env-required:"true"`
	Name          string        `yaml:"name" env-required:"true"`
	User          string        `yaml:"user" env-required:"true"`
	Password      string        `yaml:"password" env-required:"true"`
	Schema        string        `yaml:"schema" env-default:"public"`
	MigrationPath string        `yaml:"migrationPath" env-required:"true"`
	MaxAttempts   int           `yaml:"maxAttempts" env-required:"true"`
	AttemptDelay  time.Duration `yaml:"attemptDelay" env-required:"true"`
}

type Redis struct {
	Host string `yaml:"host" env-required:"true"`
	Port int    `yaml:"port" env-required:"true"`
	Db   int    `yaml:"db" `
}

type Sentry struct {
	Host string `yaml:"host" env-required:"true"`
	Key  string `yaml:"key" env-required:"true"`
}

type Service struct {
	Port uint16 `yaml:"port" env-required:"true"`
}

type Ws struct {
	Port               int     `yaml:"port" env-required:"true"`
	MaxOneIpConnection int     `yaml:"maxOneIpConnection" env-required:"true"`
	Session            session `yaml:"session"`
}

type session struct {
	MinPingDuration       time.Duration `yaml:"minPingDuration" env-required:"true"`
	MaxPingDuration       time.Duration `yaml:"maxPingDuration" env-required:"true"`
	MaxInactivityDuration time.Duration `yaml:"maxInactivityDuration" env-required:"true"`
}

func MustLoad[TConfig any]() *TConfig {
	path := fetchConfigPath()

	if path == "" {
		panic("Путь конфига не найден")
	}

	if _, err := os.Stat(path); err != nil {
		panic("Файл конфига не найден")
	}

	var cfg TConfig

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("Ошибка загрузки конфига: " + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var result string
	flag.StringVar(&result, "config", "", "Путь до файла конфига")
	flag.Parse()

	if result == "" {
		result = os.Getenv("CONFIG_PATH")
	}

	return result
}
