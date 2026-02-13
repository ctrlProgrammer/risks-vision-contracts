package contracts

import "time"

type RVAsset struct {
	ID             uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name           string    `json:"name" gorm:"not null"`
	Symbol         string    `json:"symbol" gorm:"not null"`
	Region         string    `json:"region" gorm:"not null"`
	Type           string    `json:"type" gorm:"not null"`
	Currency       string    `json:"currency" gorm:"not null"`
	Price          float64   `json:"price" gorm:"not null"`
	PriceProvider  string    `json:"price_provider" gorm:"not null"`
	Logo           string    `json:"logo" gorm:"not null"`
	ProviderTarget string    `json:"provider_target" gorm:"not null"`
	Blacklist      bool      `json:"blacklist" gorm:"default:false"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (RVAsset) TableName() string {
	return "assets"
}

type RVAssetOHLCV struct {
	// Use time.Time for native Postgres TIMESTAMPTZ support
	Time   time.Time `json:"time" gorm:"primaryKey;column:time;not null"`
	Symbol string    `json:"symbol" gorm:"primaryKey;column:symbol;not null"`
	Open   float64   `json:"open" gorm:"column:open;not null"`
	High   float64   `json:"high" gorm:"column:high;not null"`
	Low    float64   `json:"low" gorm:"column:low;not null"`
	Close  float64   `json:"close" gorm:"column:close;not null"`
	Volume float64   `json:"volume" gorm:"column:volume;not null"`
}

func (RVAssetOHLCV) TableName() string {
	return "assets_ohlcv"
}
