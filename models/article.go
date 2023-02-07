package models

type Article struct {
	Id     int
	Title  string
	CateId int
	State  int
	// ArticleCate ArticleCate `gorm:"foreignKey:CateId"`
}

func (Article) TableName() string {
	return "article"
}
