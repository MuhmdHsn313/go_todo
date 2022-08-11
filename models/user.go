package models

import "time"

type User struct {
	BaiscModel BaiscModel `gorm:"embedded"`
	FullName   string     `json:"full_name" gorm:"not null"`
	Email      string     `json:"email" gorm:"not null;unique"`
	Password   string     `json:"password" gorm:"not null"`
}

type Session struct {
	ID          uint      `json:"id" gorm:"primarykey"`
	UserID      uint      `json:"user_id" gorm:"not null"`
	AccessToken string    `json:"access_token" gorm:"not null;unique"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime:true"`
}
