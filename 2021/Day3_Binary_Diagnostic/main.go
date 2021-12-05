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

	//fmt.Println(part1(scanner))
	fmt.Println(part2(scanner))
}

func next(scanner *bufio.Scanner) (ok bool, bits string) {
	if ok = scanner.Scan(); ok {
		input := scanner.Text()
		if n, err := fmt.Sscan(input, &bits); err != nil {
			panic(err)
		} else if n != 1 {
			panic("didn't scan two tokens")
		}
	}
	return
}

func next2(scanner *bufio.Scanner) (ok bool, n int64, width int) {
	var err error
	if ok = scanner.Scan(); ok {
		input := scanner.Text()
		width = len(input)
		n, err = strconv.ParseInt(input, 2, 32)
		if err != nil {
			panic(err)
		}
	}
	return
}


func part1(scanner *bufio.Scanner) int {
	total := 0
	ok, bits := next(scanner)
	ones := make([]int, len(bits))
	for ok {
		total++
		for i, b := range bits {
			if b == '1' {
				ones[i]++
			}
		}
		ok, bits = next(scanner)
	}
	fmt.Println(ones)
	gamma := 0
	epsilon := 0
	for _, c := range ones {
		gamma *= 2
		epsilon *= 2
		if c >= (total - c) {
			gamma++
		} else {
			epsilon++
		}
	}
	fmt.Printf("total: %v\n", total)
	fmt.Printf("gamma: %v\n", gamma)
	fmt.Printf("epsilon: %v\n", epsilon)
	return gamma*epsilon
}

func part2(scanner *bufio.Scanner) int64 {
	// get all values
	values := []int64{}
	width := 0
	for true {
		ok, v, w := next2(scanner)
		if !ok {
			break
		}
		values = append(values, v)
		width = w
	}
	fmt.Println("Width: ", width)
	fmt.Println("Count: ", len(values))

	O2vals := values
	CO2vals := values
	for mask := 1 << (width - 1); mask > 0; mask >>= 1 {
		fmt.Println("Processing stage ", mask)
		if len(O2vals) > 1 {
			O2vals = applyBitCriteria(O2vals, int64(mask), false)
		}
		if len(CO2vals) > 1 {
			CO2vals = applyBitCriteria(CO2vals, int64(mask), true)
		}

	}
	if len(O2vals) != 1 {
		panic(fmt.Sprintln("finished with incorrect number of O2 vals", len(O2vals)))
	}
	if len(CO2vals) != 1 {
		panic(fmt.Sprintln("finished with incorrect number of O2 vals", len(CO2vals)))
	}
	fmt.Println(" O2: ", O2vals[0])
	fmt.Println("CO2: ", CO2vals[0])
	return O2vals[0] * CO2vals[0]
}

func applyBitCriteria(input []int64, mask int64, flip bool) []int64 {
	filtered := make([]int64, 0, len(input))
	mostCommon := countMostCommon(input, mask)
	for _, v := range input {
		match := v & mask == mostCommon
		if match != flip {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func countMostCommon(input []int64, mask int64) int64 {
	count := 0
	for _, v := range input {
		if v & mask != 0 {
			count++
		}
	}
	if count >= (len(input) - count) {
		return mask
	}
	return 0
}