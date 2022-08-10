package models

import (
	"time"
)

type BaiscModel struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime:true"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime:true"`
}
