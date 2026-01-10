package reporter

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/ayuxsec/spike/pkg/logger"
)

func sendTelegramMessage(chatID int, botToken string, message string, requestTimeout int) error {
	endpoint := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)
	form := url.Values{}
	form.Set("chat_id", fmt.Sprintf("%d", chatID))
	form.Set("text", message)
	data := form.Encode()
	logger.Debugf("Sending POST request to Telegram API: %s", endpoint)
	req, err := http.NewRequest("POST", endpoint, strings.NewReader(data))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{Timeout: time.Duration(requestTimeout) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}
	defer resp.Body.Close()

	if !logger.DisableDebug {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to read response body from telegram API: %v", err)
		}
		logger.Debugf("Telegram API response recieved: %s", string(respBody))
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send message recieved invalid status code. Response: %v", resp)
	}
	return nil
}
