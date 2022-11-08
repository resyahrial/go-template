package models

import (
	"time"
)

type CommonField struct {
	Id        string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	IsDeleted bool
}
