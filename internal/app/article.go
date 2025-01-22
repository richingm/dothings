package app

import (
	"context"
	"richingm/dothings/internal/entity"
)

type ArticleApp struct {
}

func NewArticleApp() *ArticleApp {
	return &ArticleApp{}
}

// GetArticlesByBookId 根据书籍id获取所有文章
func (*ArticleApp) GetArticlesByBookId(ctx context.Context, bookId int64) ([]entity.ArticleDo, error) {
	return nil, nil
}

// GetArticle 获取文章信息
func (*ArticleApp) GetArticle(ctx context.Context, id int64) (*entity.ArticleDo, error) {
	return nil, nil
}

// AddArticle 添加文章
func (*ArticleApp) AddArticle(ctx context.Context, do entity.ArticleDo) (*entity.ArticleDo, error) {
	return nil, nil
}

// UpdateArticle 修改文章
func (*ArticleApp) UpdateArticle(ctx context.Context, id int64, do entity.ArticleDo) (*entity.ArticleDo, error) {
	return nil, nil
}

// DelArticle 删除文章
func (*ArticleApp) DelArticle(ctx context.Context, id int64) error {
	return nil
}
