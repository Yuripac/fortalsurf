package notifier

import (
	"fmt"
	"net/http"
	"os"
)

type Telegram struct {
	botToken string
	chatID string
}

func NewTelegram() Telegram {
	return Telegram{
		os.Getenv("TELEGRAM_BOT_TOKEN"),
		os.Getenv("TELEGRAM_CHAT_ID"),
	}
}

func (t Telegram) Send(text string) (*http.Response, error) {
	url := fmt.Sprintf(
		"https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s",
		t.botToken, t.chatID, text)

	return http.Get(url)
}
