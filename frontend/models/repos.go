

package models

type Repository struct {
	Model

	Name string `json:"name"`
	CountStars int `json:"stars"`
	CountForks int `json:"count_forks"`
	CollectID uint
	Collect   Collect `gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // use Company.CompanyID as re
}