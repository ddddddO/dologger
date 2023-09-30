package main

import (
	"os"

	// "github.com/ddddddO/dologger"
)

func main() {
	// plainMode()
	// jsonMode()
	// treeMode()
}

/*
func plainMode() {
	log := dologger.New(os.Stdout)
	log.WithPlain()

	log.Debug("for debug").
		Str("name", "ddddd").
		Int("id", 1111).
		Out()

	// Output:
	// DEBUG message:for debug name:ddddd id:1111
}

func jsonMode() {
	log := dologger.New(os.Stdout)
	// default mode
	// log.WithJSON()

	log.Debug("for debug").
		Str("name", "ddddd").
		Int("id", 1111).
		Out()

	// Output:
	// {"level":"DEBUG","message":"for debug","name":"ddddd","id":1111}
}

func treeMode() {
	log := dologger.New(os.Stdout)
	log.WithTree()

	log.Debug("for debug").
		Str("name", "ddddd").
		Int("id", 1111).
		Out()

	// Output:
	// DEBUG
	// ├── message
	// │   └── for debug
	// ├── name
	// │   └── ddddd
	// └── id
	//     └── 1111
}
*/