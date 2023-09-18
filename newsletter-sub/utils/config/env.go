package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	App struct {
		Port       string
		Mode       string
		AppName    string
		AppVersion string
		Timezone   string
	}

	GoMail struct {
		SenderEmail    string
		SenderPassword string
	}

	PubSub struct {
		ProjectId        string
		SubscriptionName string
	}
}

var AppCfg AppConfig

func init() {

	var config AppConfig

	config.App.Port = GetEnv("PORT", "8000")
	config.App.Mode = GetEnv("MODE", "Local")
	config.App.AppName = GetEnv("APP_NAME", "Newsletter Publisher")
	config.App.AppVersion = GetEnv("APP_VERSION", "1.0.0")
	config.App.Timezone = GetEnv("TIMEZONE", "Asia/Jakarta")

	config.GoMail.SenderEmail = GetEnv("SENDER_MAIL", "")
	config.GoMail.SenderPassword = GetEnv("SENDER_PASSWORD", "")

	config.PubSub.ProjectId = GetEnv("PROJECT_ID", "")
	config.PubSub.SubscriptionName = GetEnv("SUBSCRIPTION_NAME", "")

	AppCfg = config
}

func GetEnv(key, fallback string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
