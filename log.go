package dologger

import (
	"bytes"
	"fmt"
	"strconv"
)

type Logger struct {
	buf *bytes.Buffer
}

func New() *Logger {
	return &Logger{}
}

func (l *Logger) Debug() *Logger {
	l.buf = &bytes.Buffer{}
	_, err := l.buf.WriteString(l.addDebugColor("DEBUG"))
	if err != nil {
		panic(err)
	}
	return l
}

func (l *Logger) Info() *Logger {
	l.buf = &bytes.Buffer{}
	_, err := l.buf.WriteString(l.addInfoColor("INFO"))
	if err != nil {
		panic(err)
	}
	return l
}

func (l *Logger) Str(key, msg string) *Logger {
	tmp := key + ":" + msg
	if !l.isBufEmpty() {
		tmp = " " + tmp
	}

	_, err := l.buf.WriteString(tmp)
	if err != nil {
		panic(err)
	}
	return l
}

func (l *Logger) Int(key string, n int) *Logger {
	i := strconv.Itoa(n)
	tmp := key + ":" + i
	if !l.isBufEmpty() {
		tmp = " " + tmp
	}

	_, err := l.buf.WriteString(tmp)
	if err != nil {
		panic(err)
	}
	return l
}

func (l *Logger) isBufEmpty() bool {
	return l.buf.Len() == 0
}

// TODO: Loggerにio.Writerを持たせて、File or 標準出力できるようにする
func (l *Logger) Output() {
	fmt.Println(l.buf.String())
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

func (*Logger) addDebugColor(level string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", uint8(Yellow), level)
}

func (*Logger) addInfoColor(level string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", uint8(Blue), level)
}
