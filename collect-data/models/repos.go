

package models

type Repository struct {
	Model

	Name string `json:"name" gorm:"not null"`
	CountStars int `json:"stars" gorm:"not null"`
	CountForks int `json:"count_forks" gorm:"not null"`
	CollectID uint `gorm:"not null"`
	Collect   Collect `gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // use Company.CompanyID as re
}