package migration

import (
	"log"
	"rest/api/configs"
	userModel "rest/api/models/User"
)

func RunMigration() {
	err := configs.DB.AutoMigrate(&userModel.User{})
	if err != nil {
		log.Println(err)
	}
}
