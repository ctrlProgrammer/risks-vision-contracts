package users

import "time"

type UserCalls struct {
	ID           int       `gorm:"primaryKey;autoIncrement"`
	UserEmail    string    `gorm:"not null"`
	InputTokens  int       `gorm:"not null"`
	OutputTokens int       `gorm:"not null"`
	CallType     string    `gorm:"not null"`
	TotalCalls   int       `gorm:"not null"`
	UsedModel    string    `gorm:"not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}
