package service

import (
	"errors"

	"github.com/huyenph/lingobox/config"
	"github.com/huyenph/lingobox/model"
	"gorm.io/gorm"
)

func InsertUser(telegramID int64, username string, language string) (*model.User, error) {
	var user model.User

	err := config.DB.Where("telegram_id = ?", telegramID).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		user = model.User{
			TelegramID: telegramID,
			Username:   username,
			Language:   language,
		}
		if err := config.DB.Create(&user).Error; err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserByTelegramID(telegramID int64) (*model.User, error) {
	var user model.User
	err := config.DB.Where("telegram_id = ?", telegramID).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
