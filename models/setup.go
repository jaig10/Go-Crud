package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"fmt"
)

var DB *gorm.DB

func ConnectDatabase () {
	dsn := "host=localhost user=postgres password=your_password dbname=blog port=5432 sslmode=disable timezone=Asia/Shanghai"
    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})  // change the database provider if necessary

    if err != nil {
        panic("Failed to connect to database!")
    } else { 
		fmt.Println("Database Connected Successfully")
	}

    database.AutoMigrate(&Post{})  // register Post model

    DB = database
}