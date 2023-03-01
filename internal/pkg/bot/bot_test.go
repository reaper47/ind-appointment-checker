package bot_test

import (
	"os"
	"testing"

	"github.com/reaper47/ind-appointment-checker/internal/pkg/bot"
)

func TestInit(t *testing.T) {
	defer func() {
		bot.Clear()
	}()

	testcases := []struct {
		name string
		in   struct {
			chatID string
			botID  string
		}
		wantErr bool
	}{
		{
			name: "bot ok",
			in: struct {
				chatID string
				botID  string
			}{"-846671787", "6945786513:NHFwLDeuUZs69uH7qvaUuYuFhy90tVPNu9Z"},
			wantErr: false,
		},
		{
			name: "chatID not only numbers",
			in: struct {
				chatID string
				botID  string
			}{"-8JK671787", "6945786513:NHFwLDeuUZs69uH7qvaUuYuFhy90tVPNu9Z"},
			wantErr: true,
		},
		{
			name: "botID separated into two parts",
			in: struct {
				chatID string
				botID  string
			}{"-846671787", "6945786513NHFwLDeuUZs69uH7qvaUuYuFhy90tVPNu9Z"},
			wantErr: true,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			_ = os.Setenv("TELEGRAM_CHATID", tc.in.chatID)
			_ = os.Setenv("TELEGRAM_BOTID", tc.in.botID)
			defer teardown()

			err := bot.Init()
			if err == nil && tc.wantErr {
				t.Errorf("got %q but want %v with %v", err, tc.wantErr, tc.in)
			}
		})
	}
}

func TestSendMessage(t *testing.T) {
	t.Run("bot is not initialized", func(t *testing.T) {
		err := bot.SendMessage("test")

		if err == nil {
			t.Error(err)
		}
	})

	t.Run("wrong bot config", func(t *testing.T) {
		_ = os.Setenv("TELEGRAM_CHATID", "-8466717879232985793")
		_ = os.Setenv("TELEGRAM_BOTID", "6945786513NHFwLDeuUZs")
		defer teardown()

		err := bot.SendMessage("test")

		if err == nil {
			t.Error(err)
		}
	})

	t.Run("wrong bot config", func(t *testing.T) {
		_ = os.Setenv("TELEGRAM_CHATID", "-8466717879232985793")
		_ = os.Setenv("TELEGRAM_BOTID", "6945786513NHFwLDeuUZs")
		defer teardown()

		err := bot.SendMessage("test")

		if err == nil {
			t.Error(err)
		}
	})

	t.Run("bot is valid", func(t *testing.T) {
		_ = os.Setenv("TELEGRAM_CHATID", "-8466717879232985793")
		_ = os.Setenv("TELEGRAM_BOTID", "694578651:3NHFwLDeuUZs")
		defer teardown()
		bot.Init()

		err := bot.SendMessage("test")

		if err != nil {
			t.Error(err)
		}
	})
}

func teardown() {
	_ = os.Unsetenv("TELEGRAM_CHATID")
	_ = os.Unsetenv("TELEGRAM_BOTID")
	bot.Clear()
}
