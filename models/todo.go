package models

type Todo struct {
	BaiscModel BaiscModel `gorm:"embedded"`
	Title      string     `json:"title" gorm:"not null"`
	Body       string     `json:"body" gorm:"not null"`
	IsDone     bool       `json:"is_done" gorm:"not null;default:false"`
}
