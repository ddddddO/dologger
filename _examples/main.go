package main

import (
	"os"

	dl "github.com/ddddddO/dologger"
)

func main() {
	logger := dl.New(os.Stdout)

	logger.Debug("for debug").
		Str("name", "ddddd").
		Int("id", 1111).
		Output()
}
