package intset

import (
	"bytes"
	"fmt"
	"sort"
)

// An MapIntSet is a set of small non-negative integers.
type MapIntSet map[uint64]bool

// Has reports whether the set contains the non-negative value x.
func (s MapIntSet) Has(x int) bool {
	return s[uint64(x)]
}

// Add adds the non-negative value x to the set.
func (s MapIntSet) Add(x int) {
	s[uint64(x)] = true
}

// UnionWith sets s to the union of s and t.
func (s MapIntSet) UnionWith(t MapIntSet) {
	for i, tword := range t {
		if tword {
			s[i] = true
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s MapIntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')

	keys := make([]uint64, 0, len(s))
	for i := range s {
		keys = append(keys, i)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	for _, k := range keys {
		if buf.Len() > 1 {
			buf.WriteByte(' ')
		}
		fmt.Fprintf(&buf, "%d", k)
	}
	buf.WriteByte('}')
	return buf.String()
}
