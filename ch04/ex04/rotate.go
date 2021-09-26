package main

import "fmt"

func rotate(x []int, count int) {
	if len(x) != 0 {
		count = count % len(x)
	}
	if len(x) == 0 || count == 0 {
		return
	}

	tmp := x
	for len(tmp) > count {
		for i := count - 1; i > -1; i-- {
			tmp[i], tmp[len(tmp)-count+i] = tmp[len(tmp)-count+i], tmp[i]
		}
		tmp = tmp[:len(tmp)-count]
	}
}

func main() {
	slice1 := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(slice1)
	rotate(slice1, 2)
	fmt.Println(slice1)

	slice2 := []int{0, 1}
	fmt.Println(slice2)
	rotate(slice2, 3)
	fmt.Println(slice2)

	slice3 := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice3)
	rotate(slice3, 10)
	fmt.Println(slice3)
}
