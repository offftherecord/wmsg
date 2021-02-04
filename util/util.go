package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

const (
	defaultTimeout = 10
	blockStart     = "```"
	blockEnd       = "```"
	newline        = "\n"
)

type slackRequestBody struct {
	Text string `json:"text"`
}

func blockFormat(msg string) string {
	return blockStart + msg + blockEnd
}

// https://stackoverflow.com/questions/25686109/split-string-by-length-in-golang
func chunks(s string, chunkSize int) []string {
	if chunkSize >= len(s) {
		return []string{s}
	}
	var chunks []string
	chunk := make([]rune, chunkSize)
	len := 0
	for _, r := range s {
		chunk[len] = r
		len++
		if len == chunkSize {
			chunks = append(chunks, string(chunk))
			len = 0
		}
	}
	if len > 0 {
		chunks = append(chunks, string(chunk[:len]))
	}
	return chunks
}

// Send will post to an 'Incoming Webook' url
func Send(webookURL string, m string, code bool, timeout int) error {

	// Slack has a maximum message limit so chunk large messages
	msgs := chunks(m, 3000)

	for _, msg := range msgs {
		if code {
			msg = blockFormat(msg)
		}

		slackBody, _ := json.Marshal(slackRequestBody{Text: msg})
		req, err := http.NewRequest(http.MethodPost, webookURL, bytes.NewBuffer(slackBody))
		if err != nil {
			return err
		}

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{Timeout: time.Duration(timeout) * time.Second}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}

		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		if buf.String() != "ok" {
			return errors.New("Error sending message to webhook")
		}
	}
	return nil
}
