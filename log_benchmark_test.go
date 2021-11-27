package dologger

import (
	"os"
	"testing"
)

func BenchmarkLogger_Debug(b *testing.B) {
	log := New(os.Stdout)
	log.WithPlain()

	for n := 0; n < b.N; n++ {
		log.Debug("for debug").
			Str("name", "ddddd").
			Int("id", n).
			Out()
	}
}
