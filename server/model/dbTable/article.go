package dbTable

import "github.com/ChocolateAceCream/blog/global"

type Article struct {
	global.MODEL
	Title     string `json:"title" gorm:"comment:article title" binding:"required"`
	Author    User   `json:"author" gorm:"foreignKey:AuthorID"`
	AuthorID  string `json:"authorID" gorm:"comment:foreignKey"`
	Path      string `json:"path" gorm:"comment:path to .md file"`
	Abstract  string `json:"abstract" gorm:"comment:first line of articles"`
	Published int    `json:"published" gorm:"comment:1-published, 2-private"`
	// Tags      []Tag m2m
	// Comments []Comment
}
