package config

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func ReadToken() string {
	dir, _ := os.Getwd()

	err := godotenv.Load(filepath.Join(dir, "./.env"))
	if err != nil {
		panic(err)
	}

	TOKEN := os.Getenv("TOKEN")

	return TOKEN
}
