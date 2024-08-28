package handler

import (
	"testing"
)

func TestHandlerExit(t *testing.T) {
	testChan := make(chan string)
	exitChan := make(chan string)
	go Handler(testChan, exitChan)

	for {
		testChan <- "exit"
		select {
		case <-exitChan:
			return
		}
	}

}
