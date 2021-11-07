package counter

import (
	"bufio"
	"bytes"
)

// LineCounter counts the number of the lines.
type LineCounter int

// Write writes the bytes to LineCounter.
func (c *LineCounter) Write(p []byte) (int, error) {
	sc := bufio.NewScanner(bytes.NewReader(p))
	sc.Split(bufio.ScanLines)
	for sc.Scan() {
		*c++
	}
	return len(p), nil
}
