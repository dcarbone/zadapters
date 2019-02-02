package zstdlog_test

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
