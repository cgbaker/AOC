package main

import (
	"bufio"
)

func part1(scanner *bufio.Scanner) int {
    increases := 0
    prevVal := next(scanner)
    newVal := next(scanner)
	for newVal != -1 {
		if newVal > prevVal {
			increases++
		}
		prevVal = newVal
		newVal = next(scanner)
	}
	return increases
}
