package contracts

import "time"

type New struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id,omitempty"`
	Asset      string    `json:"asset"`
	Symbol     string    `json:"symbol"`
	Info       string    `json:"info"`
	Sentiment  string    `json:"sentiment"`
	Title      string    `json:"title"`
	Source     string    `json:"source"`
	Type       string    `json:"type"`
	SourceLink string    `json:"source_link"`
	CreatedAt  time.Time `json:"created_at,omitempty" gorm:"autoCreateTime"`
}

func (New) TableName() string {
	return "news"
}

type ReducedNews struct {
	Title     string `json:"title"`
	Info      string `json:"info"`
	Sentiment string `json:"sentiment"`
}

type ReducedNewsList []ReducedNews

func (n *ReducedNews) ToToon() string {
	return ToToon(n)
}

func (n *ReducedNewsList) ToToon() string {
	return ToToon(n)
}

func (n *New) ToToon() string {
	return ToToon(n)
}
