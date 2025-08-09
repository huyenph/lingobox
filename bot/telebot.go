package bot

import (
	"fmt"
	"log"

	"github.com/huyenph/lingobox/config"
	"github.com/huyenph/lingobox/service"
	"github.com/huyenph/lingobox/utils"
	"gopkg.in/tucnak/telebot.v2"
)

var userStates = make(map[int64]*UserState)

type UserState struct {
	WaitingForWord             bool
	WaitingForMeaning          bool
	WaitingForExampleConfirm   bool
	WaitingForExampleSentences bool

	TempWord     string
	TempMeaning  string
	TempExamples []string
}

func SetupHandlers(b *telebot.Bot) {
	cfg := config.LoadConfig()

	b.Handle("/start", func(m *telebot.Message) {
		b.Send(m.Sender, utils.StartMessage)
	})
	b.Handle("/help", func(m *telebot.Message) {
		b.Send(m.Sender, utils.GetHelpMessage(cfg.AuthorUsername, cfg.AuthorEmail))
	})

	b.Handle("/newword", func(m *telebot.Message) {
		userID := m.Sender.ID

		userStates[userID] = &UserState{
			WaitingForWord: true,
		}

		b.Send(m.Sender, "Please type the new word:")
	})

	b.Handle("/list", func(m *telebot.Message) {
		user, err := service.GetUserByTelegramID(m.Sender.ID)

		if err != nil {
			b.Send(m.Sender, "❌ Failed to get user.")
			return
		}

		words, err := service.GetUserWords(user.ID)
		if err != nil {
			b.Send(m.Sender, "❌ Failed to load your words.")
			return
		}

		if len(words) == 0 {
			b.Send(m.Sender, "You don't have any saved words yet.")
			return
		}

		response := "📚 Your saved words:\n"
		response += "\n"
		for i, w := range words {
			response += fmt.Sprintf("%d. %s: %s\n", i+1, w.Word, w.Meaning)
			if len(w.Examples) > 0 {
				response += "Examples:\n"
				for _, ex := range w.Examples {
					response += fmt.Sprintf("  - %s\n", ex.Sentence)
				}
			}
			response += "\n"
		}

		b.Send(m.Sender, response)
	})

	b.Handle(telebot.OnText, func(m *telebot.Message) {
		userID := m.Sender.ID
		state, exists := userStates[userID]
		if !exists {
			return
		}

		if state.WaitingForWord {
			state.TempWord = m.Text
			state.WaitingForWord = false
			state.WaitingForMeaning = true

			b.Send(m.Sender, "Got it! Now please type the meaning:")
			return
		}

		if state.WaitingForMeaning {
			state.TempMeaning = m.Text
			state.WaitingForMeaning = false
			state.WaitingForExampleConfirm = true

			b.Send(m.Sender, "Do you want to add example sentences? (yes/no)")
			return
		}

		if state.WaitingForExampleConfirm {
			text := m.Text
			if text == "yes" || text == "Yes" {
				state.WaitingForExampleConfirm = false
				state.WaitingForExampleSentences = true
				state.TempExamples = []string{}

				b.Send(m.Sender, "Please type your example sentences one by one. Type 'done' when finished.")
			} else {
				user, err := service.InsertUser(m.Sender.ID, m.Sender.Username, m.Sender.LanguageCode)
				if err != nil {
					log.Println("CreateUser error:", err)
					return
				}

				word, err := service.InserWord(user, state.TempWord, state.TempMeaning, m.Sender.LanguageCode, nil)
				if err != nil {
					log.Println("CreateWord error:", err)
					return
				}

				log.Println("Word saved:", word.Word)

				delete(userStates, userID)
				b.Send(m.Sender, "Got it! Your word has been saved without examples.")
			}
			return
		}

		if state.WaitingForExampleSentences {
			if m.Text == "done" {
				user, err := service.InsertUser(
					m.Sender.ID,
					m.Sender.Username,
					m.Sender.LanguageCode,
				)
				if err != nil {
					log.Println("CreateUser error:", err)
					return
				}

				word, err := service.InserWord(
					user,
					state.TempWord,
					state.TempMeaning,
					m.Sender.LanguageCode,
					state.TempExamples,
				)
				if err != nil {
					log.Println("CreateWord error:", err)
					return
				}

				log.Println("Word saved:", word.Word)

				delete(userStates, userID)
				b.Send(m.Sender, "Thanks! Your word and examples are saved.")
			} else {
				state.TempExamples = append(state.TempExamples, m.Text)
				b.Send(m.Sender, "Example added. Add another or type 'done' to finish.")
			}
			return
		}
	})
}
