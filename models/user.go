package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	TelegramID int64  `gorm:"uniqueIndex;column:telegram_id"`
	Username   string `gorm:"column:username"`
	Language   string `gorm:"column:language"`
}
