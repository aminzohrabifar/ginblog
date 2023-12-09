package migration

import (
	"fmt"
	articleModels "ginblog/internal/modules/article/models"
	userModels "ginblog/internal/modules/user/models"
	"ginblog/pkg/database"
	"log"
)

func Migrate() {
	db := database.Connection()
	err := db.AutoMigrate(&userModels.User{}, &articleModels.Article{})
	if err != nil {
		log.Fatal("cant migrate")
		return
	}
	fmt.Println("Migration done ...")
}
