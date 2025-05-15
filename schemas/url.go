package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Url struct {
	gorm.Model
	OriginalUrl string
	ShortId     string
}

type UrlResponse struct {
	ID          uint       `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" gorm:"index"`
	OriginalUrl string     `json:"original_url"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
