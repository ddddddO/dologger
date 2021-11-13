package dologger

import (
	"bytes"
	"testing"
)

func TestLogger_Debug(t *testing.T) {
	buf := &bytes.Buffer{}
	log := New(buf)
	log.Debug().Str("name", "ddddd").Int("id", 1111).Output()
	got := buf.String()
	want := "\x1b[33mDEBUG\x1b[0m name:ddddd id:1111\n"

	if got != want {
		t.Errorf("got:%s\nwant:%s", got, want)
	}
}

func TestLogger_Info(t *testing.T) {
	buf := &bytes.Buffer{}
	log := New(buf)
	log.Info().Str("name", "aaaaa").Int("id", 22222).Output()
	got := buf.String()
	want := "\x1b[34mINFO\x1b[0m name:aaaaa id:22222\n"

	if got != want {
		t.Errorf("got:%s\nwant:%s", got, want)
	}
}
