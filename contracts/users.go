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
	ID           int       `json:"id" gorm:"primaryKey;autoIncrement"`
	UserEmail    string    `json:"user_email" gorm:"not null"`
	InputTokens  int       `json:"input_tokens" gorm:"not null"`
	OutputTokens int       `json:"output_tokens" gorm:"not null"`
	CallType     string    `json:"call_type" gorm:"not null"`
	TotalCalls   int       `json:"total_calls" gorm:"not null"`
	UsedModel    string    `json:"used_model" gorm:"not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
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
