package initializers

import "github.com/prateek69/go-auth/models"

func SyncDB() {
	DB.AutoMigrate(&models.User{})
}
