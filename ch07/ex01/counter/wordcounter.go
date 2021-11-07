package counter

import (
	"bufio"
	"bytes"
)

// WordCounter counts the number of the words.
type WordCounter int

// Write writes the bytes to WordCounter.
func (c *WordCounter) Write(p []byte) (int, error) {
	sc := bufio.NewScanner(bytes.NewReader(p))
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		*c++
	}
	return len(p), nil
}
