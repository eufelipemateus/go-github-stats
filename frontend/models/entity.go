package models

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint `json:"id"  gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time`json:"updated_at"`

}

type Entity interface {
	Count(db *gorm.DB) int64
	Take(db *gorm.DB, limit int, offset int) interface{}
	All(db *gorm.DB) interface {}
}

func All(db *gorm.DB, entity Entity) interface{} {
	data := entity.All(db)
	return data
}