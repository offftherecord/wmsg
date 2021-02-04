package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/offftherecord/wmsg/util"
)

const defaultTimeout = 10

func main() {
	var webhook string
	flag.StringVar(&webhook, "w", "", "Webhook to post to.")

	var code bool
	flag.BoolVar(&code, "c", false, "Format message to use code block. (default false)")

	var timeout int
	flag.IntVar(&timeout, "t", defaultTimeout, "Timeout in seconds.")

	var help bool
	flag.BoolVar(&help, "h", false, "Print this help information.")

	flag.Parse()

	if help {
		fmt.Println("Usage: wmsg -w <webhook>")
		flag.PrintDefaults()
		os.Exit(0)
	}

	if webhook == "" {
		fmt.Println("Usage: wmsg -w <webhook>")
		fmt.Println("Missing webhook.")
		flag.PrintDefaults()
		os.Exit(1)
	}

	var msg string
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		msg += sc.Text() + "\n"
	}

	err := util.Send(webhook, msg, code, timeout)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
