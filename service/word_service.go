package service

import (
	"github.com/huyenph/lingobox/config"
	"github.com/huyenph/lingobox/model"
)

func InserWord(
	user *model.User,
	wordStr string,
	meaning string,
	language string,
	examples []string,
) (*model.Word, error) {

	word := model.Word{
		UserID:   user.ID,
		Word:     wordStr,
		Meaning:  meaning,
		Language: language,
	}

	for _, ex := range examples {
		word.Examples = append(word.Examples, model.Example{
			Sentence: ex,
		})
	}

	if err := config.DB.Create(&word).Error; err != nil {
		return nil, err
	}

	return &word, nil
}

func GetUserWords(userID uint) ([]model.Word, error) {
	var words []model.Word

	err := config.DB.
		Preload("Examples").
		Where("user_id = ?", userID).
		Find(&words).Error

	if err != nil {
		return nil, err
	}
	return words, nil
}
