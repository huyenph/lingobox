package models

import (
	"gorm.io/gorm"
)

type Word struct {
	gorm.Model
	UserID   uint   `gorm:"index:idx_user_word,unique"`
	User     User   `gorm:"constraint:OnDelete:CASCADE;"`
	Word     string `gorm:"index:idx_user_word,unique;column:word"`
	Meaning  string `gorm:"column:meaning"`
	Language string `gorm:"column:language"`
}
