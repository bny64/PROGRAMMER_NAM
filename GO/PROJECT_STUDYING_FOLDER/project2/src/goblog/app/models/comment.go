package models

import (
	"time"
)

type Comment struct {
	Id int64 `gorm:"column:id;type:INT AUTO_INCREMENT;PRIMARY_KEY"`
	Body string `gorm:"column:body;type:TEXT;NOT NULL"`
	Commenter string `gorm:"column:commenter;type:TEXT;NOT NULL"`
	PostId int `gorm:"column:post_id;type:INT;NOT NULL"`
	CreatedAt time.Time `gorm:"column:created_at;type:DATETIME;NOT NULL"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:DATETIME;NOT NULL"`
}