package reporter

import (
	"testing"
)

func TestSendTelegramMessage(t *testing.T) {
	chatId := 123456789          // replace with a valid chat ID for real testing
	botToken := "TEST_BOT_TOKEN" // replace with a valid bot token for real testing
	if err := sendTelegramMessage(chatId, botToken, "Test message", 5); err == nil {
		t.Errorf("Expected error for invalid bot token, got nil")
	}
}
