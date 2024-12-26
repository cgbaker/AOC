package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	// "github.com/cgbaker/AOC/2024/utils"
)

type Problem struct {
	locks [][]int
	keys [][]int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	problem := readInput(file)
	fmt.Printf("problem: %d locks, %d keys\n",len(problem.locks),len(problem.keys))
	part1(problem)
	part2(problem)
}

func part1(problem *Problem) {
	sum := 0
	for _, lock := range problem.locks {
		for _, key := range problem.keys {
			if fit(key,lock) {
				sum++
			}
		}
	}
	fmt.Printf("part1: %d\n",sum)
}

func part2(_ *Problem) {
	fmt.Printf("part2: %d\n",0)
}

func fit(key, lock []int) bool {
	for c := range 5 {
		if key[c] + lock[c] > 5 {
			return false
		}
	}
	return true
}

func readInput(file *os.File) *Problem {
	locks := [][]int{}
	keys := [][]int{}
	lineScanner := bufio.NewScanner(file)
	lineScanner.Split(bufio.ScanLines)
	for lineScanner.Scan() {
		if len(lineScanner.Text()) == 0 {
			continue
		}
		if lineScanner.Text() == "....." {
			// key
			key := make([]int,5)
			lineScanner.Scan()
			for range 5 {
				s := lineScanner.Text()
				for c, v := range s {
					if v == '#' {
						key[c]++
					}
				}
				lineScanner.Scan()
			}
			keys = append(keys,key)
			lineScanner.Scan()
		} else if lineScanner.Text() == "#####" {
			// lock
			lock := make([]int,5)
			lineScanner.Scan()
			for range 5 {
				s := lineScanner.Text()
				for c, v := range s {
					if v == '#' {
						lock[c]++
					}
				}
				lineScanner.Scan()
			}
			locks = append(locks,lock)
			lineScanner.Scan()
		} else {
			panic(fmt.Sprintf("bad parsing assumption: %s",lineScanner.Text()))
		}
	}
	return &Problem{
		keys: keys,
		locks: locks,
	}
}


