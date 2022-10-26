package sequence

import (
	"time"
)

type Letters struct {
	Letters []string
}

type Sequence struct {
	Id        int64     `gorm:"column:id;primary_key;autoIncrement" json:"id"`
	Letters   string    `gorm:"column:letters;not null" json:"letters"`
	IsValid   bool      `gorm:"column:is_valid;not null" json:"is_valid"`
	CreatedAt time.Time `gorm:"column:created_at;not null" json:"created_at"`
}

type StatsInformationResponse struct {
	CountValid   int64   `json:"count_valid" example:"40"`
	CountInvalid int64   `json:"count_invalid" example:"60"`
	Ratio        float64 `json:"ratio" example:"0.4"`
	Message      string  `json:"message" example:"Bad Request"`
}

type SequenceResponseValidate struct {
	IsValid bool   `json:"is_valid" example:"true"`
	Message string `json:"message" example:"Bad Request"`
}
