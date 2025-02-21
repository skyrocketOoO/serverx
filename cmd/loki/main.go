package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type LokiPayload struct {
	Streams []Stream `json:"streams"`
}

type Stream struct {
	Stream map[string]string `json:"stream"`
	Values [][]string        `json:"values"`
}

func sendLogToLoki(level, message string) {
	// Construct the payload
	timestamp := fmt.Sprintf("%d", time.Now().UnixNano())
	payload := LokiPayload{
		Streams: []Stream{
			{
				Stream: map[string]string{
					"job":   "my-go-server",
					"env":   "production",
					"level": level,
				},
				Values: [][]string{
					{timestamp, message},
				},
			},
		},
	}

	// Convert to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal payload: %v", err)
		return
	}

	// Send to Loki
	resp, err := http.Post("http://localhost:3100/loki/api/v1/push", "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Printf("Failed to send log to Loki: %v", err)
		return
	}
	defer resp.Body.Close()

	// Check response
	if resp.StatusCode != http.StatusNoContent {
		log.Printf("Unexpected Loki response: %d %s", resp.StatusCode, resp.Status)
	}
}

func main() {
	sendLogToLoki("info", "This is an info log sent to Loki")
	sendLogToLoki("error", "This is an error log sent to Loki")
	sendLogToLoki("warning", "This is an warning log sent to Loki")
	sendLogToLoki("debug", "This is an debug log sent to Loki")
	sendLogToLoki("panic", "This is an panic log sent to Loki")
	sendLogToLoki("fatal", "This is an fatal log sent to Loki")
	sendLogToLoki("trace", "This is an trace log sent to Loki")
}
