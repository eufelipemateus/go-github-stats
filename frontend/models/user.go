package models

type User struct {
	Model
	Username  string `json:"username" gorm:"unique;not null"`
	Relevance int  `json:"relevance" gorm:"default:100;not null"`
	Show bool `json:"show"  gorm:"default:false"`
	Following []*Collect `gorm:"many2many:followers;"`
}

type APIUser struct {
	ID   uint `json:"id"`
	Username string `json:"username"`
}
