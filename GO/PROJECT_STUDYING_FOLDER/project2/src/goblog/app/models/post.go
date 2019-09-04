package models

import (
	"time"
)

type Post struct {
	Id int `gorm:"column:id"` //tag : mysql에서 컬럼명을 정의.
	Title string `gorm:"column:title"`
	Body string `gorm:"column:body"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}