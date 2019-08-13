package models

import (
	"time"
)

type Atr struct {
	ID   uint   `gorm:"primary_key" json:"ID"`
	Name string `json:"Name"`

	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"-"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"-"`
}
