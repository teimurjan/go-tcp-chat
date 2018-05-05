package server

import (
	"fmt"
	"os"
	"strconv"
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

func logMessage(text string) {
	fmt.Fprint(os.Stdout, messageLog{body: text})
}

func logError(text string) {
	fmt.Fprint(os.Stderr, messageLog{body: text, isError: true})
}
