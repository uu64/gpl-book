package popcount3

// PopCount3 returns the population count (number of set bits) of x.
func PopCount3(x uint64) int {
	var c int
	for x != 0 {
		x = x & (x - 1)
		c++
	}
	return c
}
