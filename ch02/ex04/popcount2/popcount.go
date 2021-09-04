package popcount2

// PopCount2 returns the population count (number of set bits) of x.
func PopCount2(x uint64) int {
	var c uint64
	for i := 0; i < 64; i++ {
		c += (x >> i) & 1
	}
	return int(c)
}
