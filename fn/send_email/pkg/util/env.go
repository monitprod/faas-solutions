package util

import (
	"log"

	"github.com/joho/godotenv"
)

func StartEnv() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("Starting Env. . .")

	envPath := GetRootPath() + "/.env"
	err := godotenv.Load(envPath)

	if err != nil {
		log.Println("INFO: send_email dot env not initialized:", err)
	}

	log.Println("Send Email Env Started!")
}
