package models

import "time"

type Post struct {
	ID 			uint 		`json:"id" gorm:"primaryKey"`
	Title 		string		`json:"title"`
	Content 	string		`json: "content"`
	Author   	string      `json:"author`
	CreatedAT 	time.Time   `json:created_at`
}