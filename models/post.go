package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Post struct {
	ID 			uint 		`json:"id" gorm:"primaryKey"`
	Title 		string		`json:"title"`
	Content 	string		`json: "content"`
	Author   	string      `json:"author`
	CreatedAT 	time.Time   `json:created_at`
}

type User struct {
	ID			uint 		`json:"id" gorm:"primaryKey"`
	Username	string 		`json:username  gorm:"not null; unique"` 
	Password	string 		`json:password`
}

func (user *User) HashPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err!= nil{
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}
func (user *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}