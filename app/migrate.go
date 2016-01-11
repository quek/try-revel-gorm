package app

import (
	"github.com/myaccount/my-app/app/models"
)

func RunMigrate() {
	DB.AutoMigrate(&models.User{})
}
