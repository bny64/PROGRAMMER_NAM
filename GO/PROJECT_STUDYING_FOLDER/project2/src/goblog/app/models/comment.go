package models

import (
	"time"
)

type Comment struct {
	Id int64 `gorm:"column:id;type:INT;PRIMARY_KEY;AUTO_INCREMENT"`
	Body string `gorm:"column:body;type:TEXT;NOT NULL"`
	Commenter string `gorm:"column:commenter;type:TEXT;NOT NULL"`
	PostId string `gorm:"column:post_id;type:INT;NOT NULL"`
	CreatedAt time.Time `gorm:"column:created_at;type:DATETIME;NOT NULL"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:DATETIME;NOT NULL"`
}