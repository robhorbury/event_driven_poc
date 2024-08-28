package events

import "encoding/json"

type Event struct {
	Name string `json:"name"`
	Info string `json:"info"`
}

func (input *Event) Encode() (string, error) {
	res, err := json.Marshal(input)
	return string(res), err
}

func DecodeEvent(input string) (Event, error) {
	newEvent := Event{}

	err := json.Unmarshal([]byte(input), &newEvent)

	return newEvent, err
}
