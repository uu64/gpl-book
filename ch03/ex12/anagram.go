package main

import (
	"bytes"
	"fmt"
	"os"
	"sort"
)

type runes []rune

func (s runes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s runes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s runes) Len() int {
	return len(s)
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(runes(r))
	return string(r)
}

func isAnagram(s1, s2 string) bool {
	buf1 := bytes.ToLower([]byte(s1))
	buf2 := bytes.ToLower([]byte(s2))

	buf1 = bytes.ReplaceAll(buf1, []byte{' '}, []byte{})
	buf2 = bytes.ReplaceAll(buf2, []byte{' '}, []byte{})

	return sortString(string(buf1)) == sortString(string(buf2))
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage: anagram WORD1 WORD2")
		os.Exit(1)
	}
	fmt.Println(isAnagram(os.Args[1], os.Args[2]))
}
