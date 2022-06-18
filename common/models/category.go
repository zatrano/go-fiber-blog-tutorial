package models

type Category struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Posts       []Post
}
