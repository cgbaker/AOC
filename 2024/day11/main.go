package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"github.com/cgbaker/AOC/2024/utils"
)

type Problem = []Stone
type Stone = int

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	input := readInput(file)
	prob1(input)
	prob2(input)
}

func prob1(problem Problem) {
	fmt.Printf("prob1: %d\n", solve(25,problem))
}

func prob2(problem Problem) {
	fmt.Printf("prob2: %d\n", solve(75,problem))
}

func solve(numBlinks int, problem Problem) int {
	stones := map[Stone]int{}
	for _, p := range problem {
		stones[p] = stones[p]+1
	}
	for range numBlinks {
		next := map[Stone]int{}
		for k, v := range stones {
			if k == 0 {
				next[1] = next[1] + v
			} else if splits := split(k); splits != nil {
				next[splits[0]] = next[splits[0]] + v
				next[splits[1]] = next[splits[1]] + v
			} else {
				next[k*2024] = next[k*2024]+v
			}
		}
		stones = next
	}
	sum := 0
	for _, v := range stones {
		sum += v
	}
	return sum
}

func split(s Stone) []Stone {
	if sstr := fmt.Sprintf("%d",s); len(sstr) % 2 == 0 {
		s1 := utils.Atoi(sstr[:len(sstr)/2])
		s2 := utils.Atoi(sstr[len(sstr)/2:])
		return []Stone{s1,s2}
	}	
	return nil
}

func readInput(file *os.File) Problem {
	stones := []Stone{}
	wordScanner := bufio.NewScanner(file)
	wordScanner.Split(bufio.ScanWords)
	for wordScanner.Scan() {
		s := utils.Atoi(wordScanner.Text())
		stones = append(stones, s)
	}
	return stones
}


