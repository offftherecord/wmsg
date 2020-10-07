# toslack
A command line tool to post messages to slack from stdin.

# Install
```
go get github.com/offftherecord/toslack
```
# Usage
```Usage: toslack -w <webhook>
Missing webhook.
  -c    Wrap message in code block
  -w string
        Webhook to post to
```
Basic usage
```
echo Hello World | toslack -w <webhook url>
```
If you want a nicer looking format you can use the `-c` flag which will wrap your input around code blocks

```
echo '{"test": "value"}' | jq | toslack -w <webhook url> -c
```