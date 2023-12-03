package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	DatabaseUrl string
	AppPort     string
}

func GetVariables() *Environment {
	err := godotenv.Load("vault/.env")
	if err != nil {
		fmt.Print("error")
		//log.Fatal("Failed to load variable. [Error: %s]", err)
	}

	return &Environment{
		DatabaseUrl: os.Getenv("DATABASE_URL"),
		AppPort:     os.Getenv("PORT"),
	}
}
