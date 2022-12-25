package response

import "github.com/ChocolateAceCream/blog/model/dbTable"

type Login struct {
	User *dbTable.User `json:"user"`
}
