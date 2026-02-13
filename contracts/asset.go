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
	Time   time.Time `json:"time" gorm:"primaryKey;not null"`
	Symbol string    `json:"symbol" gorm:"primaryKey;not null"`
	Open   float64   `json:"open" gorm:"not null"`
	High   float64   `json:"high" gorm:"not null"`
	Low    float64   `json:"low" gorm:"not null"`
	Close  float64   `json:"close" gorm:"not null"`
	Volume float64   `json:"volume" gorm:"not null"`
}

func (RVAssetOHLCV) TableName() string {
	return "assets_ohlcv"
}
