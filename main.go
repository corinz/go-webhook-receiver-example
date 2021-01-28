package main

import wh "github.com/corinz/go-webhook-receiver/tree/dev"

func main() {
	wh.ExecuteThisWhen("slack-notify.sh", "")
	wh.Startup("/")
}
