package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Value struct {
	Database Database
	Auth     Auth
	Log      Log
	Server   Server
	Xendit   Xendit
	Mailer   Mailer
}

type Database struct {
	DbUrl      string
	DbPort     string
	DbName     string
	DbUser     string
	DbPassword string
}

type Auth struct {
	SecretKey string
}

type Server struct {
	Base string
	Port int
}

type Log struct {
	Level string
}

type Xendit struct {
	ApiKey       string
	PublicKey    string
	WebhookToken string
}

type Mailer struct {
	MailerHost     string
	MailerPort     int
	MailerUser     string
	MailerPassword string
}

func InitEnv() (*Value, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	port, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		return nil, err
	}

	mailerPort, err := strconv.Atoi(os.Getenv("MAILER_PORT"))
	if err != nil {
		return nil, err
	}

	return &Value{
		Database: Database{
			DbUrl:      os.Getenv("DB_URL"),
			DbPort:     os.Getenv("DB_PORT"),
			DbName:     os.Getenv("DB_NAME"),
			DbUser:     os.Getenv("DB_USER"),
			DbPassword: os.Getenv("DB_PASSWORD"),
		},
		Auth: Auth{
			SecretKey: os.Getenv("AUTH_SECRETKEY"),
		},
		Log: Log{
			Level: os.Getenv("LOG_LEVEL"),
		},
		Server: Server{
			Base: os.Getenv("SERVER_BASE"),
			Port: port,
		},
		Xendit: Xendit{
			ApiKey:       os.Getenv("XENDIT_API_KEY"),
			PublicKey:    os.Getenv("XENDIT_PUBLIC_KEY"),
			WebhookToken: os.Getenv("XENDIT_WEBHOOK_TOKEN"),
		},
		Mailer: Mailer{
			MailerHost:     os.Getenv("MAILER_HOST"),
			MailerPort:     mailerPort,
			MailerUser:     os.Getenv("MAILER_USER"),
			MailerPassword: os.Getenv("MAILER_PASSWORD"),
		},
	}, nil
}
