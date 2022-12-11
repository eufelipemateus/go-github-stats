package models

import "time"


type Collect struct {
	ID        uint `json:"id"  gorm:"primary_key"`

	CountLanguages int `json:"count_languages"`
	CountFollowers int `json:"count_followers"`
	CountRepos int `json:"count_repos"`
	CountStars int  `json:"count_stars"`
	UserID uint
	User   User `gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // use Company.CompanyID as re

	CollectStartedAt time.Time `json:"collected_at" gorm:"autoCreateTime:true" `
	FinishedAt time.Time `json:"finished_at" gorm:"autoCreateTime:false" `

	Followers []*User `gorm:"many2many:followers;"`
}
