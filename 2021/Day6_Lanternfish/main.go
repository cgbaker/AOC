package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	population := readPopulation(scanner)

	fmt.Println("Initial population:", sum(population))

	fmt.Println(part1(population, 256))
}

func part1(population []int, numDays int) int {
	for i := 0; i < numDays; i++ {
		population = cycle(population)
	}
	return sum(population)
}

func cycle(population []int) []int {
	newPop := make([]int,9)
	newPop[0] = population[1]
	newPop[1] = population[2]
	newPop[2] = population[3]
	newPop[3] = population[4]
	newPop[4] = population[5]
	newPop[5] = population[6]
	newPop[6] = population[7] + population[0]
	newPop[7] = population[8]
	newPop[8] = population[0]
	return newPop
}

func readPopulation(scanner *bufio.Scanner) []int {
	population := make([]int,9)
	if ok := scanner.Scan(); !ok  {
		panic("scanning first word")
	}
	tokens := strings.Split(scanner.Text(), ",")
	for _, t := range tokens {
		var v int
		if n, err := fmt.Sscan(t, &v); n == 1 && err == nil {
			population[v]++
		}
	}
	return population
}

func sum(population []int) int {
	sum := 0
	for _, p := range population {
		sum += p
	}
	return sum
}