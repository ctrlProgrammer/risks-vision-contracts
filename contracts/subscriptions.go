package contracts

import "time"

type Subscription struct {
	ID            uint     `json:"id" gorm:"primaryKey"`
	Slug          string   `json:"slug" gorm:"uniqueIndex"`
	MaxAICalls    int      `json:"max_ai_calls" gorm:"default:0"`
	Price         float64  `json:"price" gorm:"default:0"`
	WhatsIncluded []string `json:"whats_included" gorm:"type:jsonb"`
	Limitations   []string `json:"limitations" gorm:"type:jsonb"`
	FrontName     *string  `json:"front_name" gorm:""`
}

func (s *Subscription) ToToon() string {
	return ToToon(s)
}

type TokenBundles struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	AmountOfTokens int       `json:"amount_of_tokens" gorm:"default:0"`
	Price          float64   `json:"price" gorm:"default:0"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
