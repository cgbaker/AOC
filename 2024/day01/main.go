package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"sort"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	left, right := readLists(file)
	sort.Ints(left)
	sort.Ints(right)
	fmt.Println("problem 1: ", prob1(left,right))
	fmt.Println("problem 2: ", prob2(left,right))
}

func prob1(left, right []int) int {
	sum := 0
	for i := range left {
		sum += AbsInt(left[i] - right[i])
	}
	return sum
}

func prob2(left, right []int) int {
	rcounts := map[int]int{}
	for _, v := range right {
		rcounts[v] = rcounts[v] + 1
	}
	sum := 0
	for _, l := range left {
		sum += rcounts[l]*l
	}
	return sum
}

func readLists(file *os.File) ([]int, []int) {
	left := []int{}
	right := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var l, r int
		fmt.Sscanf(scanner.Text(), "%d %d", &l, &r)
		left = append(left, l)
		right = append(right, r)
	}
	return left, right
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

