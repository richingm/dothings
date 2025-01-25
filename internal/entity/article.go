package entity

type ArticleDo struct {
	Id       int64
	Title    string
	BookId   int64
	Children []ArticleDo
	Sort     int64
}

type ArticleWithContentDo struct {
	Id       int64
	Title    string
	Content  string
	Sort     int64
	Children []ArticleWithContentDo
}

type ArticlePo struct {
	ID         int     `gorm:"primarykey" uri:"id"`
	CreatedAt  []uint8 `gorm:"column:created_at"`
	Pid        int     `gorm:"column:pid"`
	CategoryID int     `gorm:"column:category_id"`
	Title      string  `gorm:"column:title"`
	Sort       int     `gorm:"column:sort"`
}

func (ArticlePo) TableName() string {
	return "article"
}
