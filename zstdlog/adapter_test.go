package zstdlog_test

import (
	"bytes"
	"testing"

	"github.com/dcarbone/zadapters/zstdlog"
	"github.com/rs/zerolog"
)

var (
	expectedStdLoggerNoLevel = []byte("{\"message\":\"test logger\"}\n")
	expectedStdLoggerLevel   = []byte("{\"level\":\"info\",\"message\":\"test logger with level\"}\n")
)

type capture struct {
	w chan []byte
}

func newCapture(chlen int) *capture {
	return &capture{
		w: make(chan []byte, chlen),
	}
}

func (c capture) Write(p []byte) (n int, err error) {
	c.w <- p
	return
}

func TestStdLogger(t *testing.T) {
	w := newCapture(100)

	l := zerolog.New(w)

	lg := zstdlog.NewStdLogger(l)
	lg.Println("test logger")

	if msg := <-w.w; !bytes.Equal(msg, expectedStdLoggerNoLevel) {
		t.Logf("%q does not match expected %q", string(msg), string(expectedStdLoggerNoLevel))
		t.FailNow()
	}

	lgl := zstdlog.NewStdLoggerWithLevel(l, zerolog.InfoLevel)
	lgl.Println("test logger with level")
	if msg := <-w.w; !bytes.Equal(msg, expectedStdLoggerLevel) {
		t.Logf("%q does not match expected %q", string(msg), string(expectedStdLoggerLevel))
		t.FailNow()
	}
}
