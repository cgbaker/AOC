package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	crabs := readCrabs(scanner)
	N := len(crabs)
	sort.Ints(crabs)
	median := crabs[N/2]
	mean := mean(crabs)
	fmt.Println("Min:", crabs[0])
	fmt.Println("Max:", crabs[N-1])
	fmt.Println("Median:", median)
	fmt.Println("Mean:", mean)
	fmt.Println("Part1:", part1(crabs, median))
	s1 := part2(crabs,mean-1)
	s2 := part2(crabs,mean)
	s3 := part2(crabs,mean+1)
	s := min3(s1,s2,s3)
	fmt.Println("Part2:", s)
}

func part1(crabs []int, x int) int {
	cost := 0
	for _, c := range crabs {
		if c > x {
			cost += c- x
		} else {
			cost += x -c
		}
	}
	return cost
}

func part2(crabs []int, x int) int {
	cost := 0
	for _, c := range crabs {
		if x < c {
			cost += gauss(c - x)
		} else if c < x {
			cost += gauss(x - c)
		}
	}
	return cost
}

func gauss(n int) int {
	if n == 0 {
		return 0
	}
	return (n*n + n)/2
}

func readCrabs(scanner *bufio.Scanner) []int {
	crabs := []int{}
	if ok := scanner.Scan(); !ok  {
		panic("scanning first word")
	}
	tokens := strings.Split(scanner.Text(), ",")
	for _, t := range tokens {
		var v int
		if n, err := fmt.Sscan(t, &v); n == 1 && err == nil {
			crabs = append(crabs, v)
		}
	}
	return crabs
}

func mean(crabs []int) int {
	sum := 0
	for _, p := range crabs {
		sum += p
	}
	return int(math.Round(float64(sum)/float64(len(crabs))))
}

func min3(x,y,z int) int {
	return min2(x, min2(y,z))
}

func min2(x,y int) int {
	if x < y {
		return x
	}
	return y
}