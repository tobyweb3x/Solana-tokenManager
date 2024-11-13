package http

import (
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type EnvConfig struct {
	LoginProvider struct {
		Google struct {
			ClientKey string `env:"CLIENT_PROVIDER_KEY_GOOGLE"`
			Secret    string `env:"CLIENT_PROVIDER_SECRET_GOOGLE"`
		}
		Facebook struct {
			ClientKey string `env:"CLIENT_PROVIDER_KEY_FACEBOOK"`
			Secret    string `env:"CLIENT_PROVIDER_SECRET_FACEBOOK"`
		}
	}

	Database struct {
		Url string `env:"DATABASE_URL"`
	}

	Port string `env:"PORT"`

	CookieSecret  string `env:"COOKIE_SECRET"`
	SessionSecret string `env:"SESSION_SECRET"`
	Build         string `env:"BUILD"`
}

func LoadEnv(cfg *EnvConfig) error {

	if err := godotenv.Load("dev.env"); err != nil {
		log.Printf("error loading dev.env file: %v", err)
	}

	return env.Parse(cfg)
}
