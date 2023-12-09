package seeder

import (
	"fmt"
	articleModels "ginblog/internal/modules/article/models"
	userModels "ginblog/internal/modules/user/models"
	"ginblog/pkg/database"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func Seed() {
	db := database.Connection()
	password, err := bcrypt.GenerateFromPassword([]byte("123"), 1)
	if err != nil {
		log.Fatal("cant hash the password")
		return
	}
	user := userModels.User{Name: "Jinzhu", Email: "jiz@gmail.com", Password: string(password)}
	db.Create(&user)
	log.Printf("user created successfully with email address %s \n", user.Email)
	for i := 0; i < 10; i++ {
		article := articleModels.Article{
			Title:   fmt.Sprintf("Random Title %d", i),
			Content: fmt.Sprintf("Random Contenct %d", i),
			UserID:  1,
		}
		db.Create(&article)
		log.Printf("article created successfully with title %s \n", article.Title)
	}
}
