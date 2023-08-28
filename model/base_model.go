package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
