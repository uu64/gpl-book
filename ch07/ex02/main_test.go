package main

import (
	"bytes"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	w := new(bytes.Buffer)
	cw, len := CountingWriter(w)

	cw.Write([]byte("Hello, World!\n"))
	if w.String() != "Hello, World!\n" || *len != 14 {
		t.Errorf("error: input %v, got %v, len %v\n", "Hello, World\n", w.String(), *len)
	}

	cw.Write([]byte("test"))
	if w.String() != "Hello, World!\ntest" || *len != 18 {
		t.Errorf("error: input %v, got %v, len %v\n", "Hello, World\ntest", w.String(), *len)
	}
}
