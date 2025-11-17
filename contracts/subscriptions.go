package contracts

type Subscription struct {
	ID         uint    `json:"id" gorm:"primaryKey"`
	Slug       string  `json:"slug" gorm:"uniqueIndex"`
	MaxAICalls int     `json:"max_ai_calls" gorm:"default:0"`
	Price      float64 `json:"price" gorm:"default:0"`
}

func (s *Subscription) ToToon() string {
	return ToToon(s)
}
