package controllers

import (
	"sytem_service/config"
	"sytem_service/models"
	"sytem_service/utils/functions"
)

// Auto migrate
func Migrate() {
	config.DB.AutoMigrate(
		models.CommonCodeModel{},
	)
	functions.ShowLog("MigrateModel", "Success")
}
