package main

import "fmt"

func rmdups(strings []string) []string {
	i := 1
	for _, s := range strings[1:] {
		if s != strings[i-1] {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func main() {
	strings1 := []string{
		"hello", "world",
	}
	fmt.Println(rmdups(strings1))

	strings2 := []string{
		"hello", "hello", "a", "b", "b", "b", "!",
	}
	fmt.Println(rmdups((strings2)))
}
