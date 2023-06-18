package initializer

import (
	"context"
	"errors"
	"fmt"

	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/model/dbTable"
	"gorm.io/gorm"
)

const InitArticleOrder = InitUserOrder + 1

type articleInitializer struct{}

func init() {
	Register(InitArticleOrder, &articleInitializer{})
}

func (ai *articleInitializer) Name() string {
	return "article"
}

func (ai *articleInitializer) Initialize(ctx context.Context) (next context.Context, err error) {
	ru := userInitializer{}
	user, ok := ctx.Value(ru.Name() + "superadmin").(dbTable.User)
	if !ok {
		return next, fmt.Errorf("fail to find %s user in user-role initializer", "superadmin")
	}
	entities := []dbTable.Article{
		{
			Title:     "Hello World",
			Path:      "hello.md",
			Abstract:  "hello world first blog",
			Published: 1,
			Author:    user,
		},
	}
	if err = global.DB.Create(&entities).Error; err != nil {
		return ctx, fmt.Errorf("fail to init article data, err: %w", err)
	}

	next = ctx
	for _, e := range entities {
		next = context.WithValue(next, ai.Name()+e.Title, e)
	}
	return next, nil
}

func (ai *articleInitializer) InitDataVerify(ctx context.Context) bool {
	var record dbTable.Article
	err := global.DB.Preload("User").First(&record).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return record.Author.Username == "superadmin"
}
