package main

import (
	"fmt"
)

func minWithError(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("min: requires at least one argument")
	}

	min := vals[0]
	for _, v := range vals[1:] {
		if min > v {
			min = v
		}
	}
	return min, nil
}

func min(val int, vals ...int) int {
	min := val
	for _, v := range vals {
		if min > v {
			min = v
		}
	}
	return min

}

func maxWithError(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("max: requires at least one argument")
	}

	max := vals[0]
	for _, v := range vals[1:] {
		if max < v {
			max = v
		}
	}
	return max, nil

}

func max(val int, vals ...int) int {
	max := val
	for _, v := range vals {
		if max < v {
			max = v
		}
	}
	return max
}
