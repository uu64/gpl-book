package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	files := os.Args[1:]
	if len(files) == 0 {
		counts := countLines(os.Stdin)
		show(counts)
	} else {
		countsPerFile := make(map[string]map[string]int)
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countsPerFile[arg] = countLines(f)
			f.Close()
		}
		showAll(countsPerFile)
	}

}

func show(counts map[string]int) {
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func showAll(countsPerFile map[string]map[string]int) {
	countsAll := make(map[string]int)
	files := make(map[string]string)
	sep := " "
	for file, counts := range countsPerFile {
		for line, n := range counts {
			// fmt.Printf("%d\t%s\t%s\n", n, line, file)
			countsAll[line] += n
			if len(files[line]) > 0 && !strings.Contains(files[line], file) {
				files[line] += sep + file
			} else {
				files[line] = file
			}
		}
	}
	for line, n := range countsAll {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, files[line])
		}
	}
}

func countLines(f *os.File) map[string]int {
	counts := make(map[string]int)
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	return counts
}
