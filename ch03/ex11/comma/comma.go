package comma

import "bytes"

func comma(s string) string {
	var frac, sign, dot []byte

	sbuf := []byte(s)

	if sbuf[0] == '+' || sbuf[0] == '-' {
		sign = []byte{sbuf[0]}
		sbuf = sbuf[1:]
	}

	if idx := bytes.Index(sbuf, []byte(".")); idx != -1 {
		dot = []byte{'.'}
		frac = sbuf[idx+1:]
		sbuf = sbuf[:idx]
	}

	n := len(sbuf)
	if n <= 3 {
		return toString(sbuf, frac, sign, dot)
	}
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		if (n-i)%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(sbuf[i])
	}
	return toString(buf.Bytes(), frac, sign, dot)
}

func toString(intg, frac, sign, dot []byte) string {
	return string(bytes.Join([][]byte{sign, intg, dot, frac}, []byte{}))
}
