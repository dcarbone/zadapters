package zstdlog_test

import (
	"bytes"
	"testing"

	"github.com/dcarbone/zadapters/zstdlog"
	"github.com/rs/zerolog"
)

var (
	expectedStdLoggerNoLevelJSON = []byte("{\"message\":\"test logger\"}\n")
	expectedStdLoggerLevelJSON   = []byte("{\"level\":\"info\",\"message\":\"test logger with level\"}\n")

	expectedStdLoggerNoLevelConsole = []byte("<nil> %!S test logger\n")
	expectedStdLoggerLevelConsole   = []byte("<nil> INF test logger with level\n")
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
	return len(p), nil
}

func TestStdLogger(t *testing.T) {
	t.Parallel()

	t.Run("JSON", func(t *testing.T) {
		w := newCapture(1)
		defer close(w.w)
		l := zerolog.New(w)

		lg := zstdlog.NewStdLogger(l)
		lg.Println("test logger")

		if len(w.w) != 1 {
			t.Fatalf("Expected writer chan to have 1 message, but it has %d messages", len(w.w))
		}

		if msg := <-w.w; !bytes.Equal(msg, expectedStdLoggerNoLevelJSON) {
			t.Logf("%q does not match expected %q", string(msg), string(expectedStdLoggerNoLevelJSON))
			t.FailNow()
		}

		if len(w.w) > 0 {
			t.Fatalf("Expected writer chan to be empty, but it has %d elements", len(w.w))
		}

		lgl := zstdlog.NewStdLoggerWithLevel(l, zerolog.InfoLevel)
		lgl.Println("test logger with level")

		if len(w.w) != 1 {
			t.Fatalf("Expected writer chan to have 1 message, but it has %d messages", len(w.w))
		}

		if msg := <-w.w; !bytes.Equal(msg, expectedStdLoggerLevelJSON) {
			t.Logf("%q does not match expected %q", string(msg), string(expectedStdLoggerLevelJSON))
			t.FailNow()
		}

		if len(w.w) > 0 {
			t.Fatalf("Expected writer chan to be empty, but it has %d elements", len(w.w))
		}
	})

	t.Run("Console", func(t *testing.T) {
		w := newCapture(1)
		defer close(w.w)

		l := zerolog.New(zerolog.NewConsoleWriter(func(writer *zerolog.ConsoleWriter) {
			writer.NoColor = true
			writer.Out = w
		}))

		lg := zstdlog.NewStdLogger(l)
		lg.Println("test logger")

		if len(w.w) != 1 {
			t.Fatalf("Expected writer chan to have 1 message, but it has %d messages", len(w.w))
		}

		if msg := <-w.w; !bytes.Equal(msg, expectedStdLoggerNoLevelConsole) {
			t.Logf("%q does not match expected %q", string(msg), string(expectedStdLoggerNoLevelConsole))
			t.FailNow()
		}

		if len(w.w) > 0 {
			t.Fatalf("Expected writer chan to be empty, but it has %d elements", len(w.w))
		}

		lgl := zstdlog.NewStdLoggerWithLevel(l, zerolog.InfoLevel)
		lgl.Println("test logger with level")

		if len(w.w) != 1 {
			t.Fatalf("Expected writer chan to have 1 message, but it has %d messages", len(w.w))
		}

		if msg := <-w.w; !bytes.Equal(msg, expectedStdLoggerLevelConsole) {
			t.Logf("%q does not match expected %q", string(msg), string(expectedStdLoggerLevelConsole))
			t.FailNow()
		}

		if len(w.w) > 0 {
			t.Fatalf("Expected writer chan to be empty, but it has %d elements", len(w.w))
		}
	})
}
