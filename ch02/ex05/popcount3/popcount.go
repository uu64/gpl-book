package popcount3

// PopCount3 returns the population count (number of set bits) of x.
func PopCount3(x uint64) int {
	var c int
	for {
		prev := x
		x = x & (x - 1)
		if prev == x {
			break
		} else {
			c++
		}
	}
	return c
}
