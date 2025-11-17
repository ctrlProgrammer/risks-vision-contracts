package contracts

import "time"

type ChatMessage struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	SessionId string    `json:"session_id" gorm:"not null;index"`
	Role      string    `json:"role" gorm:"type:text"`
	Message   string    `json:"message" gorm:"type:text"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

type ChatMessages []ChatMessage

type ChatSession struct {
	Id        uint         `json:"id" gorm:"primaryKey;autoIncrement"`
	UserEmail string       `json:"user_email" gorm:"not null"`
	SessionId string       `json:"session_id" gorm:"unique;not null"`
	Title     string       `json:"title" gorm:"type:text"`
	Messages  ChatMessages `json:"messages" gorm:"foreignKey:SessionId;references:SessionId"`
	CreatedAt time.Time    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time    `json:"updated_at" gorm:"autoUpdateTime"`
}

func (s *ChatSession) ToToon() string {
	return ToToon(s)
}

func (m *ChatMessage) ToToon() string {
	return ToToon(m)
}

func (ms *ChatMessages) ToToon() string {
	return ToToon(ms)
}
