package controllers

import (
	"user_service/config"
	"user_service/models"
	"user_service/utils/functions"
)

// Auto migrate
func Migrate() {
	config.DB.AutoMigrate(
		models.UserModel{},
	// models.SocialInfoModel{},
	)
	functions.ShowLog("MigrateModel", "Success")
}
