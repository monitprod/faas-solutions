package util

import (
	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
)

func StartEnv() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})

	log.Println("Starting Env. . .")

	envPath := GetRootPath() + "/.env"
	err := godotenv.Load(envPath)

	if err != nil {
		log.Println("INFO: user_api dot env not initialized:", err)
	}

	log.Println("User API Env Started!")
}
