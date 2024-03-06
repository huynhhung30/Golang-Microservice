package controllers

import (
	"product_service/config"
	"product_service/models"
	"product_service/utils/functions"
)

// Auto migrate
func Migrate() {
	config.DB.AutoMigrate(
		models.ProductModel{},
	)
	functions.ShowLog("MigrateModel", "Success")
}
