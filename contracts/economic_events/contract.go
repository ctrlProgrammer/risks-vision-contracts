package economic_events

import "time"

type EconomicCalendarEvent struct {
	ID                    int       `gorm:"primaryKey;autoIncrement"`
	EventID               string    `json:"eventId" gorm:"not null;unique"`
	DateUTC               string    `json:"dateUtc"`
	PeriodDateUTC         string    `json:"periodDateUtc"`
	Actual                float64   `json:"actual"`
	Revised               float64   `json:"revised"`
	Consensus             float64   `json:"consensus"`
	RatioDeviation        float64   `json:"ratioDeviation"`
	Previous              float64   `json:"previous"`
	IsBetterThanExpected  bool      `json:"isBetterThanExpected"`
	Name                  string    `json:"name"`
	CountryCode           string    `json:"countryCode"`
	CurrencyCode          string    `json:"currencyCode"`
	Unit                  string    `json:"unit"`
	Potency               string    `json:"potency"`
	Volatility            string    `json:"volatility"`
	IsSpeech              bool      `json:"isSpeech"`
	LastUpdated           float64   `json:"lastUpdated"`
	PreviousIsPreliminary bool      `json:"previousIsPreliminary"`
	CreatedAt             time.Time `json:"created_at,omitempty" gorm:"autoCreateTime"`
	UpdatedAt             time.Time `json:"updated_at,omitempty" gorm:"autoUpdateTime"`
}
