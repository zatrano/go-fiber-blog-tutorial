package models

type Post struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CategoryId  uint
	Category    Category
}
