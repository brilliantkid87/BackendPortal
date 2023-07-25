package database

import (
	"fmt"
	"test/models"
	"test/pkg/mysql"
)

// automatic migrate
func RunMigration() {
	err := mysql.DB.AutoMigrate(&models.User{})

	if err != nil {
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
