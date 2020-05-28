package util

import (
	"io"
	"strings"
)

func NewBufferPipe(actual io.Writer) *BufferPipe {
	return &BufferPipe{Passthrough: actual}
}

type BufferPipe struct {
	Buffer      strings.Builder
	Passthrough io.Writer
}

func (b *BufferPipe) Write(p []byte) (n int, err error) {
	b.Buffer.Write(p)
	return b.Passthrough.Write(p)
}
