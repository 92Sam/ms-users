package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func GetEnviroment(typeEnv string) {
	if typeEnv == "env" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}

