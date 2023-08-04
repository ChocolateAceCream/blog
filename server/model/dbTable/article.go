package dbTable

import "github.com/ChocolateAceCream/blog/global"

type Article struct {
	global.MODEL
	Title       string `json:"title" gorm:"comment:article title" binding:"required_if=Published 1"`
	Author      User   `json:"author" gorm:"foreignKey:AuthorID"`
	AuthorID    uint   `json:"authorId" gorm:"comment:foreignKey" binding:"required_if=Published 1"`
	Abstract    string `json:"abstract" gorm:"comment:first line of articles" binding:"required_if=Published 1"`
	Content     string `json:"content" gorm:"type:text;comment:article content in md" binding:"required_if=Published 1"`
	Published   int    `json:"published" gorm:"default:2;comment:1-published, 2-private"`
	ViewedTimes int    `json:"viewedTimes" gorm:"default:0;comment: count of viewed times"`
	// Tags      []Tag m2m
	Comments []Comment `json:"comments"`
}
