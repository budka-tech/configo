package configloader

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

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
