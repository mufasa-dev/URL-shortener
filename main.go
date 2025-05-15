package main

import (
	"url-shortener/config"
	"url-shortener/router"
)

var logger config.Logger

func main() {
	logger = *config.GetLogger("main")

	if err := config.Init(); err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	router.Initialize()
}
