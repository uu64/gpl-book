package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// LimitedReader returns a Reader that reads from r but reports an EOF after n bytes.
type LimitedReader struct {
	r     io.Reader
	limit int64
}

// Read implements the io.Reader interface.
func (lr *LimitedReader) Read(p []byte) (n int, err error) {
	if lr.limit == 0 {
		return 0, io.EOF
	}
	tmp := make([]byte, lr.limit)
	n, err = lr.r.Read(tmp)
	lr.limit -= int64(n)
	copy(p, tmp)
	return
}

// LimitReader returns a LimitedReader.
func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitedReader{r, n}
}

func main() {
	r := strings.NewReader("Hello, World!")
	lr := LimitReader(r, 4)

	b := make([]byte, 8)
	for {
		n, err := lr.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

	if _, err := io.Copy(os.Stdout, lr); err != nil {
		log.Fatal(err)
	}
}
