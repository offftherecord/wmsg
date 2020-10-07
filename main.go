package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	timeout   = 10
	codeStart = "```"
	codeEnd   = "```"
	newLine   = "\n"
)

// slackRequestBody something
type slackRequestBody struct {
	Text string `json:"text"`
}

// Send will post to an 'Incoming Webook' url
func Send(webookURL string, msg string, code bool, timeout int) error {

	if code {
		msg = codeStart + msg + codeEnd
	}

	slackBody, _ := json.Marshal(slackRequestBody{Text: msg})
	req, err := http.NewRequest(http.MethodPost, webookURL, bytes.NewBuffer(slackBody))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if buf.String() != "ok" {
		return errors.New("Non-ok response returned from Slack")
	}
	return nil
}

func main() {
	var webhook string
	flag.StringVar(&webhook, "w", "", "Webhook to post to")

	var code bool
	flag.BoolVar(&code, "c", false, "Wrap message in code block")

	flag.Parse()

	// Webhook required
	if webhook == "" {
		fmt.Println("Usage: toslack -w <webhook>")
		fmt.Println("Missing webhook.")
		flag.PrintDefaults()
		os.Exit(1)
	}

	var msg string
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		msg += sc.Text() + newLine
	}

	err := Send(webhook, msg, code, 10)
	if err != nil {
		log.Fatalln(err)
	}
}
