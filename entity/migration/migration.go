package migration

import (
	"fmt"

	"github.com/fahmiammryhi/tutor-go-fiber/database"
	"github.com/fahmiammryhi/tutor-go-fiber/entity"
)

func RunMigrate() { //migrate user
	err := database.DB.AutoMigrate(&entity.User{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Success to Migrate")
}
