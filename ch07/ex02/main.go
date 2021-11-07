package main

import (
	"bytes"
	"fmt"
	"io"
)

// ByteCounter counts the number of the bytes.
type ByteCounter struct {
	len    int64
	writer io.Writer
}

// Write writes the bytes to ByteCounter.
func (c *ByteCounter) Write(p []byte) (int, error) {
	n, err := c.writer.Write(p)
	if err != nil {
		return 0, err
	}
	c.len += int64(n)
	return n, nil
}

// CountingWriter returns ByteCounter and the length of the content.
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := ByteCounter{0, w}
	return &c, &c.len
}

func main() {
	w := new(bytes.Buffer)
	cw, len := CountingWriter(w)

	cw.Write([]byte("test"))
	fmt.Println(w.String())
	fmt.Println(*len)
}
