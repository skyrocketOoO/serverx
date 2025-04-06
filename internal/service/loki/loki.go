package loki

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type LokiWriter struct {
	LokiURL string
}

func NewLokiWriter() (*LokiWriter, error) {
	lokiURL := os.Getenv("LOKI_URL")
	fullPushURL := lokiURL + "/loki/api/v1/push"
	healthURL := lokiURL + "/ready"

	// Simple health check
	resp, err := http.Get(healthURL)
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("loki not available at %s: %v", healthURL, err)
	}

	return &LokiWriter{
		LokiURL: fullPushURL,
	}, nil
}

func (lw *LokiWriter) Write(p []byte) (n int, err error) {
	logEntry := map[string]interface{}{
		"streams": []map[string]interface{}{
			{
				"stream": map[string]string{
					"job": "alarmsystem",
					"env": "dev",
				},
				"values": [][]string{
					{fmt.Sprintf("%d", time.Now().UnixNano()), string(p)},
				},
			},
		},
	}

	payload, err := json.Marshal(logEntry)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal log entry: %w", err)
	}

	resp, err := http.Post(lw.LokiURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return 0, fmt.Errorf("failed to send log to Loki: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("unexpected Loki response: %d", resp.StatusCode)
	}

	return len(p), nil
}
