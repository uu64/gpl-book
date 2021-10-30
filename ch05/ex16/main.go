package main

import (
	"strings"
)

func join(sep string, elems ...string) string {
	var sb strings.Builder
	if len(elems) == 0 {
		return ""
	}
	for _, s := range elems[:len(elems)-1] {
		sb.WriteString(s)
		sb.WriteString(sep)
	}
	sb.WriteString(elems[len(elems)-1])
	return sb.String()
}
