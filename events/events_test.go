package events

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	object := Event{Name: "test_name", Info: "test info"}
	str, err := object.Encode()

	expectedStr := `{"name":"test_name","info":"test info"}`
	assert.Equal(t, str, expectedStr)
	assert.Nil(t, err)
}

func TestDecode(t *testing.T) {
	testString := `{"name":"test_name","info":"test info"}`

	newEvent, err := DecodeEvent(testString)

	assert.Nil(t, err)
	assert.Equal(t, newEvent.Name, "test_name")
	assert.Equal(t, newEvent.Info, "test info")
}

func TestDecodeError(t *testing.T) {
	testString := "Not a valid object"

	_, err := DecodeEvent(testString)

	assert.NotNil(t, err)
}
