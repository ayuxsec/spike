package reporter

import (
	"fmt"
	"os"

	"github.com/ayuxsec/spike/internal/pkg/scanner/core"
	"github.com/ayuxsec/spike/pkg/config"
	"github.com/ayuxsec/spike/pkg/logger"
)

type TelegramNotifier struct {
	cfg *config.ReporterConfig
}

func NewTelegramNotifier(cfg *config.ReporterConfig) *TelegramNotifier {
	return &TelegramNotifier{cfg: cfg}
}

func (t *TelegramNotifier) OnDomainScanned(out *core.ScannerOutput, errs []error) {
	if err := t.sendReport(out, errs); err != nil {
		fmt.Fprintln(os.Stderr)
		logger.Errorf("Failed to send Telegram report for domain %s: %v", out.DomainScanned, err)
	}
}

func (t *TelegramNotifier) sendReport(out *core.ScannerOutput, errs []error) error {
	if !t.cfg.Telegram.Enabled {
		logger.Warn("Telegram notifications are disabled in the configuration.")
		return nil
	}

	report, err := GenerateScanReport(out)
	if err != nil {
		return fmt.Errorf("failed to generate scan report: %v", err)
	}

	err = sendTelegramMessage(
		t.cfg.Telegram.ChatID,
		t.cfg.Telegram.BotToken,
		report,
		t.cfg.Telegram.Timeout,
	)
	if err != nil {
		return fmt.Errorf("failed to send Telegram message: %v", err)
	}

	if errReport := generateErrsReport(out.DomainScanned, errs); errReport != "" {
		err = sendTelegramMessage(
			t.cfg.Telegram.ChatID,
			t.cfg.Telegram.BotToken,
			generateErrsReport(out.DomainScanned, errs),
			t.cfg.Telegram.Timeout,
		)
		if err != nil {
			return fmt.Errorf("failed to send Telegram error report: %v", err)
		}
	}
	return nil
}
