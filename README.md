# wmsg
A general purpose command line tool to send messages to slack.

## Install
```
go get -u github.com/offftherecord/wmsg
```
## Help Menu
```
Usage: wmsg -w <webhook>
  -c    Format message to use code block
  -w string
        Webhook to post to
```
## Basic usage
```
echo Hello World | wmsg -w <webhook url>
```
If you want a nicer looking format you can use the `-c` flag which will wrap your input around code blocks

```
echo '{"test": "value"}' | jq | wmsg -w <webhook url> -c
```

# Known Issues
- Build Discord service