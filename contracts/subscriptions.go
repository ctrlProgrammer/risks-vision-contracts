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

type Bundle struct {
	ID                uint      `json:"id" gorm:"primaryKey"`
	Slug              string    `json:"slug" gorm:"uniqueIndex;size:64;not null"`
	NameEn            string    `json:"name_en" gorm:"size:128;not null"`
	AudienceLabelEs   string    `json:"audience_label_es" gorm:"size:128"`
	Description       string    `json:"description" gorm:"type:text"`
	AmountOfTokens    int       `json:"amount_of_tokens" gorm:"not null"`
	Price             float64   `json:"price" gorm:"not null"`
	Equivalence       string    `json:"equivalence" gorm:"type:text"`
	IsMostPopular     bool      `json:"is_most_popular" gorm:"default:false"`
	SortOrder         int       `json:"sort_order" gorm:"index;default:0"`
	InternalCostUSD   float64   `json:"internal_cost_usd" gorm:"column:internal_cost_usd"`
	InternalProfitUSD float64   `json:"internal_profit_usd" gorm:"column:internal_profit_usd"`
	MarginPercent     float64   `json:"margin_percent" gorm:"column:margin_percent"`
	CreatedAt         time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt         time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Bundle) TableName() string {
	return "token_bundles"
}
