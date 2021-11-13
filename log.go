package dologger

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
)

type Logger struct {
	buf *bytes.Buffer
	w   io.Writer
}

func New(w io.Writer) *Logger {
	return &Logger{
		w: w,
	}
}

func (l *Logger) Debug(msg string) *Logger {
	l.buf = &bytes.Buffer{}
	_, err := l.buf.WriteString(addDebugColor("DEBUG") + " " + "message" + ":" + msg)
	if err != nil {
		panic(err)
	}
	return l
}

func (l *Logger) Info(msg string) *Logger {
	l.buf = &bytes.Buffer{}
	_, err := l.buf.WriteString(addInfoColor("INFO") + " " + "message" + ":" + msg)
	if err != nil {
		panic(err)
	}
	return l
}

func (l *Logger) Str(key, msg string) *Logger {
	tmp := " " + key + ":" + msg
	_, err := l.buf.WriteString(tmp)
	if err != nil {
		panic(err)
	}
	return l
}

func (l *Logger) Int(key string, n int) *Logger {
	tmp := " " + key + ":" + strconv.Itoa(n)
	_, err := l.buf.WriteString(tmp)
	if err != nil {
		panic(err)
	}
	return l
}

// TODO: Loggerにio.Writerを持たせて、File or 標準出力できるようにする
func (l *Logger) Output() {
	_, err := l.buf.WriteString("\n")
	if err != nil {
		panic(err)
	}
	l.w.Write(l.buf.Bytes())
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
