package bzip

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"
)

const bzip2Path = "/usr/bin/bzip2"

type writer struct {
	w      *bufio.Writer // underlying output stream
	mu     sync.Mutex
	tempfile string
}

// NewWriter returns a writer for bzip2-compressed streams.
func NewWriter(out io.Writer) io.WriteCloser {
	fmt.Println("init")
	w := &writer{w: bufio.NewWriter(out)}
	return w
}

func (w *writer) Write(data []byte) (int, error) {
	var total int // uncompressed bytes written

	w.mu.Lock()

	dir := os.TempDir()
	f, err := os.CreateTemp(dir, "gobzip2")
	if err != nil {
		return total, err
	}
	defer f.Close()

	bw := bufio.NewWriter(f)
	total, err = bw.Write(data)
	if err != nil {
		return total, err
	}

	w.tempfile = f.Name()

	w.mu.Unlock()

	return total, nil
}

// Close flushes the compressed data and closes the stream.
// It does not close the underlying io.Writer.
func (w *writer) Close() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	b, err := exec.Command(bzip2Path, "-c", w.tempfile).Output()
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
