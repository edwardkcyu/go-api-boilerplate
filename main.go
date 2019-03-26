package main

import (
	"os"

	"github.com/edwardkcyu/go-api-boilerplate/logger"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	env := os.Getenv("ENV")

	logger.Initialize()

	logger.Infow(
		"failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", "test",
		"attempt", 3,
	)

	logger.Infof("this is env %s", env)

}
