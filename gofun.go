package gofun

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func gofun() string {
	cwd, _ := os.Getwd()
	envPath := filepath.Join(cwd, ".env")
	_, err := os.Stat(envPath)
	if err == nil {
		err := godotenv.Load()
		if err != nil {
			os.Exit(1)
		}
		brick_config_name := os.Getenv("brick_config_name")
		return brick_config_name
	}
	return "brick"
}
