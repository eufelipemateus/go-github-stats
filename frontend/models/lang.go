package models

type Lang struct {
	Model

	Language string `json:"language"`
	Weight int `json:"weight"`
	Percent float32 `json:"percent" gorm:"default:0;not null" sql:"type:double(10,2);"` 
	CollectID uint
	Collect   Collect `gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // use Company.CompanyID as re
}



type ApiLang struct {  
    Language   string `json:"language"`
    Weight    int  `json:"weight"`
	Percent  float64 `json:"percent"`
}
