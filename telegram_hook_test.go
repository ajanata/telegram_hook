package telegram_hook

import (
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestNewHook(t *testing.T) {
	_, err := NewTelegramHook("", "", "")
	if err == nil {
		t.Errorf("No error on invalid Telegram API token.")
	}

	_, err = NewTelegramHook("", os.Getenv("TELEGRAM_TOKEN"), "")
	if err != nil {
		t.Fatalf("Error on valid Telegram API token: %s", err)
	}

	h, _ := NewTelegramHook("testing", os.Getenv("TELEGRAM_TOKEN"), os.Getenv("TELEGRAM_TARGET"))
	if err != nil {
		t.Fatalf("Error on valid Telegram API token and target: %s", err)
	}
	log.AddHook(h)

	log.WithFields(log.Fields{
		"animal": "walrus",
		"number": 1,
		"size":   10,
	}).Errorf("A walrus appears")
}
