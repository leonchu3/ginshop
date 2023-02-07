package models

type ArticleCate struct {
	Id      int //主建
	Title   string
	State   int
	Article []Article `gorm:"foreignKey:CateId"`
}

func (ArticleCate) TableName() string {
	return "article_cate"
}
