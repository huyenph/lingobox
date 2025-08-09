package model

import (
	"gorm.io/gorm"
)

type Example struct {
	gorm.Model
  WordID  uint
  Sentence string `gorm:"column:sentence"`
}
