package dologger

import (
	"bytes"
	"fmt"
	"strconv"
)

// console出力時のカラーは以下がつかえそう
// https://github.com/uber-go/zap/blob/master/internal/color/color.go
type Logger struct {
	buf *bytes.Buffer
}

func New() *Logger {
	return &Logger{}
}

func (l *Logger) Debug() *Logger {
	l.buf = &bytes.Buffer{}
	_, err := l.buf.WriteString("DEBUG")
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
