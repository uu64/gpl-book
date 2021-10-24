package main

import (
	"fmt"
	"strings"
)

var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	keys, cycle := topoSort(prereqs)
	for i, course := range keys {
		var suffix string
		if _, ok := cycle[course]; ok {
			suffix = fmt.Sprintf("(<-> %s)", strings.Join(cycle[course], ", "))
		}
		fmt.Printf("%d:\t%s %s\n", i+1, course, suffix)
	}
}

func topoSort(m map[string][]string) ([]string, map[string][]string) {
	var order []string
	seen := make(map[string]bool)
	cycle := make(map[string][]string)

	contains := func(path []string, item string) bool {
		for _, p := range path {
			if p == item {
				return true
			}
		}
		return false
	}

	var visitAll func(item string, path []string)

	visitAll = func(item string, path []string) {
		if contains(path, item) {
			for i, p := range path {
				cycle[p] = append(cycle[p], path[:i]...)
				cycle[p] = append(cycle[p], path[i+1:]...)
			}
			return
		}
		if !seen[item] {
			seen[item] = true
			for _, i := range m[item] {
				path = append(path, item)
				visitAll(i, path)
			}
			order = append(order, item)
		}
	}

	for key := range m {
		visitAll(key, []string{})
	}
	return order, cycle
}
