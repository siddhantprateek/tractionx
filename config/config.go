package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVar(Variable string) string {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("unable to load .env file")
	}
	variable := os.Getenv(Variable)
	return variable
}
