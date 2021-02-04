# wmsg
A general purpose command line tool to send messages to slack.

**Note:** Take caution not to spam Slack with messages

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
## Example uses
Send notifications when long running tasks are done

```
nmap -sS target; echo "Nmap scan completed"| wmsg -w <webhook>
```
Send formatted output from tools
```
amass enum -d domain -o results; cat results | anew | wmsg -w <webhook> -c
```
