package counter

import "testing"

func TestWriteCounter(t *testing.T) {
	emptyStr1 := ""
	emptyStr2 := "\n 　\t"
	singleLine1 := "test"
	singleLine2 := "Hello, World!"
	multiLine1 := "Hello, World!\nThis is a sample text.\n"

	tests := []struct {
		input []byte
		want  int
	}{
		{[]byte(emptyStr1), 0},
		{[]byte(emptyStr2), 0},
		{[]byte(singleLine1), 1},
		{[]byte(singleLine2), 2},
		{[]byte(multiLine1), 7},
	}
	for _, test := range tests {
		var c WordCounter
		_, err := c.Write(test.input)
		if err != nil {
			t.Errorf("error: c.Write(%v) returns error %v", string(test.input), err)
		}
		if WordCounter(test.want) != c {
			t.Errorf("error: c.Write(%v) = %v\n", string(test.input), int(c))
		}

		_, err = c.Write(test.input)
		if err != nil {
			t.Errorf("error: c.Write(%v) returns error %v", string(test.input), err)
		}
		if WordCounter(test.want*2) != c {
			t.Errorf("error: c.Write(%v) = %v\n", string(test.input), int(c))
		}
	}
}
