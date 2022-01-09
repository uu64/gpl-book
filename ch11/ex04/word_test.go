package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
	"unicode"

	word "gopl.io/ch11/word2"
)

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25)
	tmp := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000))
		tmp[i] = r
		tmp[n-1-i] = r
	}

	runes := []rune{}
	for _, r := range tmp {
		runes = append(runes, r)
		switch rng.Intn(5) {
		case 0:
			runes = append(runes, rune(','))
		case 1:
			runes = append(runes, rune(' '))
		default:
			continue
		}
	}

	return string(runes)
}

func randomNotPalindrome(rng *rand.Rand) string {
	n := rng.Intn(23)
	// 長さ0と1は回文になるので避ける
	n += 2
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; {
		r1 := rune(rng.Intn(0x1000))
		r2 := rune(rng.Intn(0x1000))
		if unicode.ToLower(r1) == unicode.ToLower(r2) || !unicode.IsLetter(r1) || !unicode.IsLetter(r2) {
			continue
		}

		runes[i] = r1
		runes[n-1-i] = r2
		i++
	}
	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !word.IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}

	for i := 0; i < 1000; i++ {
		p := randomNotPalindrome(rng)
		if word.IsPalindrome(p) {
			for _, r := range p {
				fmt.Printf("%s, %v\n", string(r), unicode.IsLetter(r))
			}
			t.Errorf("IsPalindrome(%q) = true", p)
		}
	}
}
