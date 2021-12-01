package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	//answer := part1(scanner)
	answer := part2(scanner)
	fmt.Println(answer)
}

func next(scanner *bufio.Scanner) int {
	var v int
	if scanner.Scan() {
		v, _ = strconv.Atoi(scanner.Text())
	} else {
		v = -1
	}
	return v
}
