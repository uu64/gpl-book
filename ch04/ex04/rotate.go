package main

import "fmt"

func rotate(x []int, count int) []int {
	if len(x) == 0 {
		return x
	}
	for i := 0; i < count; i++ {
		x = append(x, x[i])
	}
	return x[count:]
}

func main() {
	slice := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(slice)
	slice = rotate(slice, 2)
	fmt.Println(slice)
}
