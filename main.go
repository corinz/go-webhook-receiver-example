package main

import wh "github.com/corinz/go-webhook-receiver"

func main() {
	wh.ExecuteThisWhen("slack-notify.sh", "")
	wh.Startup("/")
}
