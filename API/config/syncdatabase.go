package initializer

import "github.com/ginniss2022/union/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}