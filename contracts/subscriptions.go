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

type TokenBundleOrder struct {
	ID                 uint      `json:"id" gorm:"primaryKey"`
	Reference          string    `json:"reference" gorm:"uniqueIndex;size:128;not null"`
	UserEmail          string    `json:"user_email" gorm:"index;not null;size:256"`
	BundleSlug         string    `json:"bundle_slug" gorm:"index;size:64;not null"`
	Tokens             int       `json:"tokens" gorm:"not null"`
	AmountInCents      int64     `json:"amount_in_cents" gorm:"not null"`
	AmountUSDCents     int64     `json:"amount_usd_cents" gorm:"column:amount_usd_cents;not null;default:0"`
	Currency           string    `json:"currency" gorm:"size:8;not null;default:COP"`
	Status             string    `json:"status" gorm:"size:24;not null;default:pending"` // pending, paid, failed
	WompiTransactionID string    `json:"wompi_transaction_id" gorm:"size:64"`
	CreatedAt          time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt          time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (TokenBundleOrder) TableName() string {
	return "token_bundle_orders"
}

// UserTokenWallet stores purchasable AI tokens (separate from usage call logs).
type UserTokenWallet struct {
	UserEmail    string    `json:"user_email" gorm:"primaryKey;size:256"`
	TokenBalance int       `json:"token_balance" gorm:"not null;default:0"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (UserTokenWallet) TableName() string {
	return "user_token_wallets"
}
