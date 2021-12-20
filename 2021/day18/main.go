package main

import (
	"bufio"
	"crypto/x509/pkix"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("sample1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	algo, img := readInput(file)
	fmt.Println(sum.magnitude())
}

func readInput(file *os.File) (*Algorithm, *Image) {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
}

type Number struct {
	tree []int8
}

func readNumber(scanner *bufio.Scanner) Number {
	str := scanner.Text()
	t, str := str[0], str[1:]
	switch t {

	}
	return nil
}
