package contracts

type Subscription struct {
	ID            uint     `json:"id" gorm:"primaryKey"`
	Slug          string   `json:"slug" gorm:"uniqueIndex"`
	MaxAICalls    int      `json:"max_ai_calls" gorm:"default:0"`
	Price         float64  `json:"price" gorm:"default:0"`
	WhatsIncluded []string `json:"whats_included" gorm:"type:jsonb"`
	Limitations   []string `json:"limitations" gorm:"type:jsonb"`
	FrontName     string   `json:"front_name" gorm:"not null"`
}

func (s *Subscription) ToToon() string {
	return ToToon(s)
}
