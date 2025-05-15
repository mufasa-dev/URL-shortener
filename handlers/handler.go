package handlers

import (
	"url-shortener/config"

	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("Handler")
	db = config.GetSQLite()
}
