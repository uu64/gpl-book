package bzip

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"sync"
)

const bzip2Path = "/usr/bin/bzip2"

type writer struct {
	w   io.Writer // underlying output stream
	buf bytes.Buffer
	mu  sync.Mutex
}

// NewWriter returns a writer for bzip2-compressed streams.
func NewWriter(out io.Writer) io.WriteCloser {
	w := &writer{w: out}
	return w
}

func (w *writer) Write(data []byte) (int, error) {
	var total int // uncompressed bytes written

	w.mu.Lock()

	total, err := w.buf.Write(data)
	if err != nil {
		return total, err
	}

	w.mu.Unlock()

	return total, nil
}

// Close flushes the compressed data and closes the stream.
// It does not close the underlying io.Writer.
func (w *writer) Close() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	dir := os.TempDir()
	f, err := os.CreateTemp(dir, "gobzip2")
	if err != nil {
		return err
	}
	defer f.Close()
	f.Write(w.buf.Bytes())

	b, err := exec.Command(bzip2Path, "-c", f.Name()).Output()
	if err != nil {
		return err
	}

	_, err = w.w.Write(b)
	if err != nil {
		return err
	}
	return nil
}

//!-close
