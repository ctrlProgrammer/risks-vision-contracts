package contracts

import (
	"encoding/json"
	"time"
)

type Strategy struct {
	ID               uint      `gorm:"primaryKey;autoIncrement" json:"id,omitempty"`
	UserEmail        string    `json:"user_email" gorm:"not null;index" description:"Owner user email; links the strategy to users.email"`
	AssetSymbol      string    `json:"asset_symbol"`
	Description      string    `json:"description"`
	BaseInstructions string    `json:"base_instructions"`
	Timeframe        string    `json:"timeframe" gorm:"default:'1h'"`
	CreatedAt        time.Time `json:"created_at,omitempty" gorm:"autoCreateTime"`
}

func (Strategy) TableName() string {
	return "strategies"
}

// StrategyBuild is a versioned snapshot of a strategy (processing code, prompt, and computed output).
type StrategyBuild struct {
	ID             uint            `gorm:"primaryKey;autoIncrement" json:"id,omitempty"`
	StrategyID     uint            `json:"strategy_id" gorm:"not null;index;uniqueIndex:idx_strategy_build_version" description:"Parent strategy this build belongs to"`
	VersionNumber  int             `json:"version_number" gorm:"not null;uniqueIndex:idx_strategy_build_version" description:"Monotonic version for this strategy (unique per strategy_id)"`
	ProcessingCode string          `json:"processing_code" gorm:"type:text;not null" description:"Executable or DSL text used to process the initial market/strategy data"`
	UserPrompt     string          `json:"user_prompt" gorm:"type:text" description:"The user prompt text saved when this strategy version was created or refined"`
	Results        json.RawMessage `json:"results,omitempty" gorm:"type:jsonb" description:"Structured output of the strategy run (metrics, signals, etc.)"`
	CreatedAt      time.Time       `json:"created_at,omitempty" gorm:"autoCreateTime"`
	UpdatedAt      time.Time       `json:"updated_at,omitempty" gorm:"autoUpdateTime"`
}

func (StrategyBuild) TableName() string {
	return "strategy_builds"
}
