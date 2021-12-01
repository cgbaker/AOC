package main

import (
	"bufio"
)

func part2(scanner *bufio.Scanner) int {
	increases := 0
	a := next(scanner)
	b := next(scanner)
	c := next(scanner)
	d := next(scanner)
	for d != -1 {
		if d > a {
			increases++
		}
		a, b, c, d = b, c, d, next(scanner)
	}
	return increases
}
