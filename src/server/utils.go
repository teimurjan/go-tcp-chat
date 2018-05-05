package server

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func formatAddress(port int) string {
	return ":" + strconv.Itoa(port)
}

func handleError(e error) {
	if e != nil {
		logError(e.Error())
		os.Exit(errorCodeAtExit)
	}
}

func log(text string, isError bool) {
	message := fmt.Sprintf("%v: %v\n", time.Now().UTC(), text)
	if isError {
		message = "ERROR: " + message
	} else {
		message = "DEBUG: " + message
	}
	fmt.Fprint(os.Stdout, message)
}

func logError(text string) {
	log(text, true)
}

func logMessage(text string) {
	log(text, false)
}
