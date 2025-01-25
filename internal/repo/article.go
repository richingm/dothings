package repo

import (
	"context"
	"gorm.io/gorm"
	"richingm/dothings/internal/entity"
	"richingm/dothings/middleware"
)

type ArticleRepo struct {
}

func NewArticleRepo() *ArticleRepo {
	return &ArticleRepo{}
}

func (r *ArticleRepo) Create(ctx context.Context, po *entity.ArticlePo) (*entity.ArticlePo, error) {
	db := ctx.Value(middleware.DbContextKey).(*gorm.DB)
	err := db.Model(&entity.ArticlePo{}).Create(po).Error
	if err != nil {
		return nil, err
	}
	return po, nil
}

func (r *ArticleRepo) Update(ctx context.Context, id int, fields map[string]interface{}) error {
	db := ctx.Value(middleware.DbContextKey).(*gorm.DB)
	err := db.Model(&entity.ArticlePo{}).Where("id = ?", id).Updates(fields).Error
	if err != nil {
		return err
	}
	return nil
}
