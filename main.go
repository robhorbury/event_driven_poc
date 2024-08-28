package main

import (
	"fmt"
	"os"

	"example.com/event_driven_poc/events"
	"example.com/event_driven_poc/handler"
	"example.com/event_driven_poc/listener"
)

func main() {

	test := events.Event{Name: "rob", Info: "test"}

	fmt.Println(test.Encode())
	events := make(chan string)
	exit := make(chan string)

	go listener.Listener(events, exit)
	go handler.Handler(events, exit)

	for {
		select {
		case <-exit:
			os.Exit(0)
		}
	}

}
