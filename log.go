package dologger

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
)

type Logger struct {
	buf  *bytes.Buffer
	w    io.Writer
	mode outputMode
}

func New(w io.Writer) *Logger {
	return &Logger{
		w: w,
	}
}

type outputMode int

// TODO: 各処理内でmodeで分岐してそれぞれの処理をする、ではなく、interface越しに呼び出すようにする
const (
	modeJSON outputMode = iota // default
	modePlain
	modeTree
)

func (l *Logger) WithJSON() {
	l.mode = modeJSON
}

// for console
func (l *Logger) WithPlain() {
	l.mode = modePlain
}

func (l *Logger) WithTree() {
	l.mode = modeTree
}

func (l *Logger) Debug(msg string) *Logger {
	l.buf = &bytes.Buffer{}
	tmp := ""

	switch l.mode {
	case modePlain:
		tmp = addDebugColor("DEBUG") + " " + "message" + ":" + msg
	case modeJSON:
		tmp = quote("level") + ":" + quote("DEBUG") + "," + quote("message") + ":" + quote(msg)
	}

	_, err := l.buf.WriteString(tmp)
	if err != nil {
		panic(err)
	}

	return l
}

func quote(s string) string {
	return `"` + s + `"`
}

func (l *Logger) Info(msg string) *Logger {
	l.buf = &bytes.Buffer{}
	tmp := ""

	switch l.mode {
	case modePlain:
		tmp = addInfoColor("INFO") + " " + "message" + ":" + msg
	case modeJSON:
		tmp = quote("level") + ":" + quote("INFO") + "," + quote("message") + ":" + quote(msg)
	}

	_, err := l.buf.WriteString(tmp)
	if err != nil {
		panic(err)
	}

	return l
}

func (l *Logger) Str(key, msg string) *Logger {
	tmp := ""

	switch l.mode {
	case modePlain:
		tmp = " " + key + ":" + msg
	case modeJSON:
		tmp = "," + quote(key) + ":" + quote(msg)
	}

	_, err := l.buf.WriteString(tmp)
	if err != nil {
		panic(err)
	}

	return l
}

func (l *Logger) Int(key string, n int) *Logger {
	tmp := ""

	switch l.mode {
	case modePlain:
		tmp = " " + key + ":" + strconv.Itoa(n)
	case modeJSON:
		tmp = "," + quote(key) + ":" + strconv.Itoa(n)
	}

	_, err := l.buf.WriteString(tmp)
	if err != nil {
		panic(err)
	}

	return l
}

// TODO: Loggerにio.Writerを持たせて、File or 標準出力できるようにする
func (l *Logger) Output() {
	var tmp []byte

	switch l.mode {
	case modePlain:
		_, err := l.buf.WriteString("\n")
		if err != nil {
			panic(err)
		}
		tmp = l.buf.Bytes()
	case modeJSON:
		tmp = append([]byte("{"), l.buf.Bytes()...)
		tmp = append(tmp, []byte("}\n")...)
	}

	l.w.Write(tmp)
}

// https://github.com/uber-go/zap/blob/master/internal/color/color.go
const (
	Black Color = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

type Color uint8

func addDebugColor(level string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", uint8(Yellow), level)
}

func addInfoColor(level string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", uint8(Blue), level)
}
