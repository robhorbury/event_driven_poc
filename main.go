package main

import (
	"os"

	"example.com/event_driven_poc/handler"
	"example.com/event_driven_poc/listener"
)

func main() {
	events := make(chan string)
	exit := make(chan string)

	go listener.Listener(events)
	go handler.Handler(events, exit)

	for {
		select {
		case <-exit:
			os.Exit(0)
		}
	}

}
