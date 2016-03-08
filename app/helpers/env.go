package helpers

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/revel/revel"
)

func InitEnv() {
	var err error
	if revel.Config.BoolDefault("mode.test", false) {
		err = godotenv.Load(".env_test")
	} else if revel.Config.BoolDefault("mode.dev", false) {
		err = godotenv.Load(".env_dev")
	}
	if err != nil {
		log.Fatal("Error loading .env_test file")
	}
}
