package reporter

import (
	"fmt"
	"spike/internal/scanner/core"
	"spike/pkg/config"
	"spike/pkg/logger"
)

type TelegramNotifier struct {
	cfg *config.ReporterConfig
}

func NewTelegramNotifier(cfg *config.ReporterConfig) *TelegramNotifier {
	return &TelegramNotifier{cfg: cfg}
}

func (t *TelegramNotifier) OnDomainScanned(out *core.ScannerOutput, errs []error) {
	if err := t.sendReport(out, errs); err != nil {
		logger.Errorf("Failed to send Telegram report: %v", err)
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

	err = sendTelegramMessage(
		t.cfg.Telegram.ChatID,
		t.cfg.Telegram.BotToken,
		generateErrsReport(out.DomainScanned, errs),
		t.cfg.Telegram.Timeout,
	)
	if err != nil {
		return fmt.Errorf("failed to send Telegram error report: %v", err)
	}

	return nil
}
