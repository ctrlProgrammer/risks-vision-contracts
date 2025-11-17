package contracts

import "time"

type User struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	Email          string `json:"email" default:"NULL"`
	TotalAICalls   int    `json:"total_ai_calls" gorm:"default:0"`
	SubscriptionID int    `json:"subscription_id" gorm:"default:7"`
	Wallet         string `json:"wallet" default:"NULL"`
}

type UserBetaTester struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" default:"NULL"`
	Wallet    string    `json:"wallet" default:"NULL"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

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

type UserCallLogsStats struct {
	TotalTokens   uint        `json:"total_tokens"`
	TotalCalls    uint        `json:"total_calls"`
	GroupedLogs   []UserCalls `json:"grouped_logs"`
	UnGroupedLogs []UserCalls `json:"un_grouped_logs"`
}

type UserPlatformStats struct {
	TotalTokens   uint              `json:"total_tokens"`
	TotalCalls    uint              `json:"total_calls"`
	UserCallLogs  UserCallLogsStats `json:"user_call_logs"`
	User          User              `json:"user"`
	Subscriptions []Subscription    `json:"subscriptions"`
}
