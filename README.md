# To Hook
A command line tool to post messages to a webhook from stdin.

## Install
```
go get github.com/offftherecord/tohook
```
## Help Menu
```
Usage: tohook -w <webhook>
  -c    Wrap message in code block
  -w string
        Webhook to post to
```
## Basic usage
**Note:** Slack is the default service used
```
echo Hello World | tohook -w <webhook url>
```
If you want a nicer looking format you can use the `-c` flag which will wrap your input around code blocks

```
echo '{"test": "value"}' | jq | tohook -w <webhook url> -c
```

# TODO:
- Build Discord service