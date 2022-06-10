package config

import (
	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
)

type Conf struct {
	Port   string `env:"PORT,default=8080"`
	Secret string `env:"SECRET,default=SECRET"`
}

var Config Conf

func initConf() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	if err := envdecode.Decode(&Config); err != nil {
		panic(err)
	}
}

func GetConfig() Conf {
	initConf()
	return Config
}
