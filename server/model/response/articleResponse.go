package response

import (
	"time"

	"github.com/ChocolateAceCream/blog/model/dbTable"
)

type ArticleInfo struct {
	dbTable.Article
	File []byte `json:"file"`
}

type ArticleBaseInfo struct {
	UpdatedAt time.Time `json:"updatedAt"`
	Title     string    `json:"title"`
	Abstract  string    `json:"abstract"`
	Content   string    `json:"content"`
	AuthorID  uint      `json:"authorId"`
	Author    string    `json:"author"`
}

func ArticleBaseInfoFormatter(a dbTable.Article) ArticleBaseInfo {
	return ArticleBaseInfo{
		UpdatedAt: a.MODEL.UpdatedAt,
		Title:     a.Title,
		Abstract:  a.Abstract,
		Content:   a.Content,
		AuthorID:  a.AuthorID,
		Author:    a.Author.Username,
	}
}

/* for nested struct with field from other package, such as dbTable.Article from above struct, the struct usage would be like this:
article = response.ArticleInfo{
	File:    fileBytes,
	Article: a,
}

go will omit package name and use the field name for the field
*/
