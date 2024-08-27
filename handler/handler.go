package handler

import (
	"fmt"
	"strings"
)

func Handler(events chan string, exit chan<- string) {
	for {
		event := strings.TrimSpace(<-events)
		fmt.Println(event)
		switch event {
		case "exit":
			fmt.Println("stopping program flow")
			exit <- event
		}
	}

}
