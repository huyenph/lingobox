package utils

import (
	"fmt"
)

func GetHelpMessage(authorName string, authorEmail string) string {
	return fmt.Sprintf(`🆘 LingoBox Help

Welcome to the LingoBox bot. Here are some tips to get you started:

1. Use /newword to add new vocabulary words step-by-step.
2. You can add multiple example sentences for each word.
3. Use /list to see all the words you've saved.
4. Use /start anytime to see the introduction message again.

If you have any questions, feedback, or need support, feel free to reach out:

📞 Contact Us

• Telegram: @%s
• Email: %s

We’re here to help!

Happy learning! 📚`, authorName, authorEmail)
}

const (
	StartMessage = `👋 Welcome to LingoBox!

LingoBox is your personal vocabulary assistant designed to help you learn new words effectively and effortlessly.

Available commands:
1. /newword – Add a new word with its meaning and examples, step-by-step.
2. /list – View all the words you've saved so far.
3. /help – Get detailed instructions on how to use the bot.

We’re excited to help you expand your vocabulary! To begin, simply type /newword and follow the prompts.

Happy learning! 📚`
)
