package model

import (
	"time"
)

type CommonField struct {
	Id        string    `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	IsDeleted bool
}
