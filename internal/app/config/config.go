package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/mcuadros/go-defaults"
)

type Config struct {
	Timeout  int `env:"TIMEOUT" default:"60"`
	Postgres Postgres
	Redis    Redis
	SMTP     SMTP
	GPT      GPT
	Port     string `env:"PORT" default:"8080"`
}

type Postgres struct {
	URL string `env:"DATABASE_URL" default:"postgres://postgres:postgres@localhost:5432/postgres"`
}

type SMTP struct {
	// APIKey string `env:"SENDGRID_API_KEY"`
	FromEmail string `env:"FROM_EMAIL"`
	FromName  string `env:"FROM_NAME"`
	SMTPHost  string `env:"SMTP_HOST"`
	SMTPPort  string `env:"SMTP_PORT"`
	SMTPUser  string `env:"SMTP_USER"`
	SMTPPass  string `env:"SMTP_PASS"`
}

type Redis struct {
	Addr string `env:"REDIS_ADDR"`
	Pass string `env:"REDIS_PASS"`
	DB   int    `env:"REDIS_DB"`
}

type GPT struct {
	Token string `env:"OPENAI_TOKEN"`
}

func New(filenames ...string) (*Config, error) {
	cfg := new(Config)

	defaults.SetDefaults(cfg)

	if len(filenames) > 0 {
		if err := godotenv.Load(filenames...); err != nil {
			return nil, err
		}
	}

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
