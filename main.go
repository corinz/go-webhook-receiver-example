package main

import wh "github.com/corinz/go-webhook-receiver"

func main() {
	wh.ExecuteThisWhen("./slack-notify.sh", "after eq 1481a2de7b2a7d02428ad93446ab166be7793fbb")
	wh.Startup("/")
}
