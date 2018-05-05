package server

import (
	"fmt"
	"time"
)

type messageLog struct {
	body    string
	isError bool
}

func (message *messageLog) String() string {
	formattedMessage := fmt.Sprintf("%v: %v", time.Now().UTC(), message.body)
	if message.isError {
		formattedMessage = "ERROR: " + formattedMessage
	} else {
		formattedMessage = "DEBUG: " + formattedMessage
	}
	return formattedMessage
}
