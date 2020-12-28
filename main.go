package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"./internals"
)

func main() {

	projectID := os.Getenv("GCP_PROJECT_ID")

	if projectID == "" {
		_ = fmt.Errorf("failed to read gcp project id from environment vars")
		os.Exit(1)
		return
	}

	publishCommand := flag.NewFlagSet("publish", flag.ExitOnError)
	publishTopic := publishCommand.String("topic", "", "Topic to publish message to.")
	publishMsg := publishCommand.String("msg", "", "Message to publish.")

	subCommand := flag.NewFlagSet("subscribe", flag.ExitOnError)
	subID := subCommand.String("sub", "", "Subscription to subscribe to.")

	if len(os.Args) < 2 {
		_ = fmt.Errorf("expected 'publish' or 'subscribe' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "publish":
		publishCommand.Parse(os.Args[2:])
		internals.Publish(projectID, *publishTopic, *publishMsg)
	case "subscribe":
		subCommand.Parse(os.Args[2:])
		internals.Subscribe(projectID, *subID)
		time.Sleep(time.Duration(3) * time.Minute)
	default:
		_ = fmt.Errorf("expected 'publish' or 'subscribe' subcommands")
		os.Exit(1)
	}
}
