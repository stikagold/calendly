package helpers

import (
	"bytes"
	"fmt"
	"log"
)

var (
	buf    bytes.Buffer
	logger = log.New(&buf, "logger: ", log.Ltime)
)

func HandleError(err error) {
	if err != nil {
		logger.Print(fmt.Errorf(fmt.Sprintf("[err] Handled error: %s", err.Error())))
		fmt.Print(&buf)
	}
}
