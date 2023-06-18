package response

import "github.com/ChocolateAceCream/blog/model/dbTable"

type ArticleInfo struct {
	dbTable.Article
	File []byte `json:"file"`
}

/* for nested struct with field from other package, such as dbTable.Article from above struct, the struct usage would be like this:
article = response.ArticleInfo{
	File:    fileBytes,
	Article: a,
}

go will omit package name and use the field name for the field
*/
