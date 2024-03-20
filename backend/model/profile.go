package model

import "gorm.io/gorm"

type Profile struct {
	UserID     uint64 `json:"id" gorm:"type:bigint;not null;uniqueIndex"`
	Username   string `json:"username" gorm:"not null"`
	Email      string `json:"email" gorm:"not null"`
	gorm.Model `json:"-"`
}
