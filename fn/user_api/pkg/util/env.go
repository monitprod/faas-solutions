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
		log.Println("INFO: user_api dot env not initialized:", err)
	}

	log.Println("User API Env Started!")
}
