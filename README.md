# LingoBox

LingoBox is a Telegram bot designed to help users learn and save new words step-by-step, with meanings and example sentences. It uses Go, GORM, and PostgreSQL for a robust backend and integrates with Telegram via the Telebot library.

---

## Features

- Add new words interactively via Telegram chat
- Save word meanings and multiple example sentences
- List all saved words with their details
- Supports user management by Telegram ID
- Clean, user-friendly conversational flow

---

## Technology Stack

- **Go** - backend language
- **GORM** - ORM for database interaction
- **PostgreSQL** - database
- **Telebot** - Telegram bot API client
- **dotenv** - environment variables management

---

## Setup and Installation

### Prerequisites

- Go 1.18+ installed
- PostgreSQL database ready
- Telegram bot token from [BotFather](https://telegram.me/BotFather)

### Steps

1. Clone the repository:

```bash
git clone https://github.com/huyenph/lingobox.git
cd lingobox
```

2. Create .env file with required environment variables:

```bash
PORT=your_port
TELEGRAM_BOT_TOKEN=your_telegram_bot_token
DATABASE_URL=postgres://<username>:<password>@<ip>:<port>/<dbname>?sslmode=disable
```

3. Run the application:

```bash
go run cmd/main.go
```

---

## Usage

Start the bot by sending `/start` in Telegram.

Add a new word with `/newword` and follow prompts for meaning and examples.

List saved words with `/list`.

Get help with `/help`.

---

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.
