package dologger

import (
	"bytes"
	"testing"
)

func TestLogger_Debug(t *testing.T) {
	buf := &bytes.Buffer{}
	log := New(buf)

	tests := []struct {
		name     string
		modeFunc func()
		in       map[string]interface{}
		want     string
	}{
		{
			name:     "output plain",
			modeFunc: log.WithPlain,
			in: map[string]interface{}{
				"message": "debug for test",
				"name":    "ddddd",
				"id":      1111,
			},
			want: "\x1b[33mDEBUG\x1b[0m message:debug for test name:ddddd id:1111\n",
		},
		{
			name:     "output json",
			modeFunc: log.WithJSON,
			in: map[string]interface{}{
				"message": "debug for test",
				"name":    "ddddd",
				"id":      1111,
			},
			want: "{\"level\":\"DEBUG\",\"message\":\"debug for test\",\"name\":\"ddddd\",\"id\":1111}\n",
		},
	}

	for _, tt := range tests {
		t.Log(tt.name)

		tt.modeFunc()

		log.Debug(tt.in["message"].(string)).Str("name", tt.in["name"].(string)).Int("id", tt.in["id"].(int)).Out()
		got := buf.String()

		if got != tt.want {
			t.Errorf("got:%s\nwant:%s", got, tt.want)
		}
		buf.Reset()
	}
}

func TestLogger_Info(t *testing.T) {
	buf := &bytes.Buffer{}
	log := New(buf)

	tests := []struct {
		name     string
		modeFunc func()
		in       map[string]interface{}
		want     string
	}{
		{
			name:     "output plain",
			modeFunc: log.WithPlain,
			in: map[string]interface{}{
				"message": "info for test",
				"name":    "aaaaa",
				"id":      22222,
			},
			want: "\x1b[34mINFO\x1b[0m message:info for test name:aaaaa id:22222\n",
		},
		{
			name:     "output json",
			modeFunc: log.WithJSON,
			in: map[string]interface{}{
				"message": "info for test",
				"name":    "aaaaa",
				"id":      22222,
			},
			want: "{\"level\":\"INFO\",\"message\":\"info for test\",\"name\":\"aaaaa\",\"id\":22222}\n",
		},
	}

	for _, tt := range tests {
		t.Log(tt.name)

		tt.modeFunc()

		log.Info(tt.in["message"].(string)).Str("name", tt.in["name"].(string)).Int("id", tt.in["id"].(int)).Out()
		got := buf.String()

		if got != tt.want {
			t.Errorf("got:%s\nwant:%s", got, tt.want)
		}
		buf.Reset()
	}
}
