package models

import (
	"time"
)

// Model is base model
type Model struct {
	CreatedAt time.Time  `json:"creation_timestamp" gorm:"not null" sql:"DEFAULT:current_timestamp"`
	UpdatedAt time.Time  `json:"update_timestamp" gorm:"not null" sql:"DEFAULT:current_timestamp on update current_timestamp"`
	DeletedAt *time.Time `sql:"index" json:"expire_timestamp"`
}
