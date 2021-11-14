package main

import (
	"fmt"
	"sort"
)

func IsPalindrome(s sort.Interface) bool {
	start := 0
	end := s.Len() - 1
	for start < end {
		if s.Less(start, end) || s.Less(end, start) {
			return false
		}
		start++
		end--
	}
	return true
}

type sentence []byte

func (x sentence) Len() int           { return len(x) }
func (x sentence) Less(i, j int) bool { return x[i] < x[j] }
func (x sentence) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func main() {
	message1 := sentence([]byte("Hello, World!"))
	fmt.Printf("%s is palindrome?: %v\n", string(message1), IsPalindrome(message1))

	message2 := sentence([]byte("level"))
	fmt.Printf("%s is palindrome?: %v\n", string(message2), IsPalindrome(message2))
}
