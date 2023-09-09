package lib

import (
	"github.com/spf13/viper"
	"log"
)

type Env struct {
	ServerPort  string `mapstructure:"SERVER_PORT"`
	Environment string `mapstructure:"ENV"`

	DBUsername string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`
}

func NewEnv() Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Не удалось прочитать конфиг!")
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Не удалось прочитать окружение: ", err)
	}
	log.Println("Конфиг успешно прочитан!")

	return env
}
