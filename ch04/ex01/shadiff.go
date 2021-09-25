package main

// 2**8 = 256 (max of 1byte)
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func shadiff(v1, v2 *[32]byte) int {
	var c byte
	for i := 0; i < 32; i++ {
		c += pc[v1[i]^v2[i]]
	}
	return int(c)
}
