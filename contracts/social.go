package contracts

import (
	"encoding/json"
	"time"
)

// SocialScheduledPost is one row in the social posting scheduler queue.
type SocialScheduledPost struct {
	ID                 uint            `json:"id" gorm:"primaryKey;autoIncrement"`
	PostText           string          `json:"post_text" gorm:"type:text;not null"`
	Summary            string          `json:"summary" gorm:"type:text"`
	SocialNetwork      string          `json:"social_network" gorm:"size:32;not null;index"`
	PostingTime        time.Time       `json:"posting_time" gorm:"not null;index"`
	PostType           string          `json:"post_type" gorm:"size:64;not null"`
	Version1           string          `json:"version_1" gorm:"column:version_1;type:text"`
	Version2           string          `json:"version_2" gorm:"column:version_2;type:text"`
	Version3           string          `json:"version_3" gorm:"column:version_3;type:text"`
	PerformanceMetrics json.RawMessage `json:"performance_metrics,omitempty" gorm:"type:jsonb"`
	State              string          `json:"state" gorm:"size:32;not null;index"`
	PublishedAt        *time.Time      `json:"published_at,omitempty"`
}

func (SocialScheduledPost) TableName() string {
	return "social_posting_scheduler"
}

type SocialTemplate struct {
	ID             uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	TemplateName   string    `json:"template_name" gorm:"size:128;not null"`
	TemplateText   string    `json:"template_text" gorm:"type:text;not null"`
	SocialNetwork  string    `json:"social_network" gorm:"size:32;not null;default:x"`
	TargetLanguage string    `json:"target_language" gorm:"size:8;not null;default:en"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Version        int       `json:"version" gorm:"not null;default:1"`
}

func (SocialTemplate) TableName() string {
	return "social_templates"
}
