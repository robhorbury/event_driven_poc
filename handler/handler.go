package handler

import (
	"fmt"
	"strings"

	"example.com/event_driven_poc/events"
)

func Handler(eventsChannel chan string, exit chan<- string) {
	for {
		eventString := strings.TrimSpace(<-eventsChannel)

		if strings.TrimSpace(eventString) == "exit" {
			exit <- eventString
		}

		// We have already tested that the events passed to this channel
		// will serialise into events Json, so parse and handle here.
		event, err := events.DecodeEvent(eventString)

		if err != nil {
			panic(err)
		}

		switch event.Name {
		case "print":
			fmt.Println(event.Info)
		}

	}
}
