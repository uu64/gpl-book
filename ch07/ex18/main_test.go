package main

import (
	"bytes"
	"testing"
)

const doc = `<body>
  <div>
    <div>
      <span>span1</span>
    </div>
    <div>
      <p class="span">paragraph1</p>
    </div>
  </div>
  <button id="submit">submit</button>
</body>
`

func TestParse(t *testing.T) {
	tests := []struct {
		path []string
		want string
	}{
		{[]string{"div", "div", "p"}, "body() div() div() p(class=span): paragraph1\n"},
		{[]string{"submit"}, "body() button(id=submit): submit\n"},
		{[]string{"div", "div", "span"}, "body() div() div() span(): span1\nbody() div() div() p(class=span): paragraph1\n"},
	}
	for _, test := range tests {
		out := new(bytes.Buffer)
		parse(bytes.NewBufferString(doc), out, test.path)

		if out.String() != test.want {
			t.Errorf("path: %v, out:\n%s\n", test.path, out.String())
		}
	}
}
