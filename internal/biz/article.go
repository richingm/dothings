package biz

type ArticleRepoIf interface {
}

type ArticleBiz struct {
	articleRepo ArticleRepoIf
}

func NewArticleBiz(articleRepo ArticleRepoIf) *ArticleBiz {
	return &ArticleBiz{}
}

func (b *ArticleBiz) Create() {

}
