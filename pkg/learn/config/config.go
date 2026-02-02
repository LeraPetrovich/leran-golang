package config

import (
	"log"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	API struct {
		Port int `env:"PORT" envDefault:"7078"`
	}
	JwtSecret string `env:"JWT_SECRET,required"`

	App struct {
		PostgresURI         string `env:"POSTGRES_URI,required"`
		PostgresConnections int    `env:"POSTGRES_CONNECTIONS" envDefault:"1"`
	}
}

func Load() Config {
	var config Config
	if err := env.Parse(&config); err != nil {
		log.Panicf("[config] parsing failed: %+v\n", err)
	}
	return config
}
