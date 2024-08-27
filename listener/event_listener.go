package listener

import (
	"bufio"
	"os"
)

func Listener(events chan<- string) {
	for {
		var reader = bufio.NewReader(os.Stdin)
		message, _ := reader.ReadString('\n')

		events <- message
	}
}
