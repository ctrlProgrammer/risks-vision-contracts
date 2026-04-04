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

// AI Report Analysis

type NewsMarketSector struct {
	Sector    string `json:"sector" description:"The market sector name (e.g. Crypto, Equities, Commodities, Forex)." required:"true"`
	Sentiment string `json:"sentiment" description:"Overall sentiment for this sector: bullish, bearish, or neutral." required:"true"`
	Summary   string `json:"summary" description:"Two to three sentence summary of relevant news and events affecting this sector today." required:"true"`
}

type NewsKeyEvent struct {
	Title  string `json:"title" description:"Short title of the event or news item." required:"true"`
	Impact string `json:"impact" description:"Expected market impact: high, medium, or low." required:"true"`
	Detail string `json:"detail" description:"One sentence describing what happened and why it matters." required:"true"`
}

type NewsReportAnalysis struct {
	Date             string             `json:"date" description:"The date this report covers in YYYY-MM-DD format." required:"true"`
	OverallSentiment string             `json:"overall_sentiment" description:"Single word overall market sentiment for the day: bullish, bearish, neutral, or mixed." required:"true"`
	MarketSummary    string             `json:"market_summary" description:"Concise 60-80 word structured summary of the day covering the most relevant news and economic events and their combined effect on financial markets." required:"true"`
	Sectors          []NewsMarketSector `json:"sectors" description:"Per-sector breakdown of sentiment and highlights." required:"true"`
	KeyEvents        []NewsKeyEvent     `json:"key_events" description:"List of the most impactful individual news items or economic events of the day." required:"true"`
	Risks            []string           `json:"risks" description:"List of key risks or uncertainties the market should watch for the coming sessions." required:"true"`
	Opportunities    []string           `json:"opportunities" description:"List of potential trading or investment opportunities arising from today's events and news." required:"true"`
	ReadableReport   string             `json:"readable_report" description:"Full natural-language report of the day written in clear, flowing prose. This field must be entirely free of markdown syntax, bullet points, or special characters — plain sentences and paragraphs only — so it can be read aloud by a text-to-speech engine without any formatting artifacts." required:"true"`
	Error            string             `json:"error" description:"Populate only if the analysis could not be completed; otherwise leave empty." required:"true"`
}

type NewsDailyMarketReport struct {
	ID        int                `json:"id" gorm:"primaryKey;autoIncrement"`
	Date      time.Time          `json:"date"`
	Report    NewsReportAnalysis `json:"report"`
	UserEmail string             `json:"user_email" gorm:"not null"`
	Market    string             `json:"market" gorm:"not null"`
	CreatedAt time.Time          `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time          `json:"updated_at" gorm:"autoUpdateTime"`
}
