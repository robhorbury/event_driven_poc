package listener

import (
	"bufio"

	"example.com/event_driven_poc/events"

	"os"
	"strings"
)

func Listener(eventsChannel chan<- string, exit chan<- string) {
	for {
		var reader = bufio.NewReader(os.Stdin)
		message, _ := reader.ReadString('\n')
		_, err := events.DecodeEvent(message)

		if strings.TrimSpace(message) == "exit" {
			exit <- message
		} else if err == nil {
			eventsChannel <- message
		}
	}
}
