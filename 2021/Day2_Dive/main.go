package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	//fmt.Println(part1(scanner))
	fmt.Println(part2(scanner))
}

func next(scanner *bufio.Scanner) (ok bool, dir string, val int) {
	if ok = scanner.Scan(); ok {
		input := scanner.Text()
		if n, err := fmt.Sscan(input, &dir, &val); err != nil {
			panic(err)
		} else if n != 2 {
			panic("didn't scan two tokens")
		}
	}
	return
}

func part1(scanner *bufio.Scanner) int {
	horiz := 0
	depth := 0
	for {
		ok, dir, val := next(scanner)
		if !ok {
			break
		}
		switch dir {
		case "forward":
			horiz += val
		case "down":
			depth += val
		case "up":
			depth -= val
		}
	}
	return horiz*depth
}

func part2(scanner *bufio.Scanner) int {
	aim := 0
	horiz := 0
	depth := 0
	for {
		ok, dir, val := next(scanner)
		if !ok {
			break
		}
		switch dir {
		case "forward":
			horiz += val
			depth += aim*val
		case "down":
			aim += val
		case "up":
			aim -= val
		}
	}
	return horiz*depth
}