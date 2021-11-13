package main

import (
	dl "github.com/ddddddO/dologger"
)

func main() {
	logger := dl.New()

	logger.Debug().
		Str("name", "ddddd").
		Int("id", 1111).
		Output()
}
