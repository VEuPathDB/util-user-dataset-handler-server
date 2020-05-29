package util

import (
	"io"
	"strings"
)

// NewBufferPipe returns a new BufferPipe instance wrapping the given io.Writer.
func NewBufferPipe(actual io.Writer) *BufferPipe {
	return &BufferPipe{Passthrough: actual}
}

// BufferPipe wraps an io.Writer instance and splits the written bytes into both
// an internal string buffer and the wrapped io.Writer.
type BufferPipe struct {
	Buffer      strings.Builder
	Passthrough io.Writer
}

func (b *BufferPipe) Write(p []byte) (n int, err error) {
	b.Buffer.Write(p)
	return b.Passthrough.Write(p)
}
