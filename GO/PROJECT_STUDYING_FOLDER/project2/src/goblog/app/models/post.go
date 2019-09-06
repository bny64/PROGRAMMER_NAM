package models

import (
	"time"
)

type Post struct {
	Id int `gorm:"column:id;type:INT;PRIMARY_KEY;AUTO_INCREMENT"` //tag : mysql에서 컬럼명을 정의.
	Title string `gorm:"column:title;type:TEXT;NOT NULL"`
	Body string `gorm:"column:body;type:TEXT;NOT NULL"`
	CreatedAt time.Time `gorm:"column:created_at;type:DATETIME;NOT NULL"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:DATETIME;NOT NULL"`		
}