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
	defaultTimeout = 10
	blockStart     = "```"
	blockEnd       = "```"
	newline        = "\n"
)

// slackRequestBody something
type slackRequestBody struct {
	Text string `json:"text"`
}

func blockFormat(msg string) string {
	return blockStart + msg + blockEnd
}

// Send will post to an 'Incoming Webook' url
func Send(webookURL string, msg string, code bool, timeout int) error {

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
		return errors.New("Non-ok response returned from Slack")
	}
	return nil
}

func main() {
	var webhook string
	flag.StringVar(&webhook, "w", "", "Webhook to post to")

	var code bool
	flag.BoolVar(&code, "c", false, "Wrap message in code block (default: false)")

	var timeout int
	flag.IntVar(&timeout, "t", defaultTimeout, "Timeout in seconds (default: 10)")

	var help bool
	flag.BoolVar(&help, "h", false, "Print this help screen")

	flag.Parse()

	if help {
		fmt.Println("Usage: tohook -w <webhook>")
		flag.PrintDefaults()
		os.Exit(0)
	}

	if webhook == "" {
		fmt.Println("Usage: tohook -w <webhook>")
		fmt.Println("Missing webhook.")
		flag.PrintDefaults()
		os.Exit(1)
	}

	var msg string
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		msg += sc.Text() + newline
	}

	err := Send(webhook, msg, code, timeout)
	if err != nil {
		log.Fatalln(err)
	}
}
