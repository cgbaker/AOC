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
	scanner.Split(bufio.ScanWords)
	//num := part1(scanner)
	num := part2(scanner)
	fmt.Println("Part2:", num)
}

func part2(scanner *bufio.Scanner) int {
	total := 0
	for {
		input, output := readLine(scanner)
		if output == nil {
			break
		}
		key := deduce(input)

		num := 0
		for _, o := range output {
			num = 10*num + key[o]
		}
		total += num
	}
	return total
}

func deduce(input []string) map[string]int {
	allDigits := map[string]struct{}{}
	for _, i := range input {
		allDigits[i] = struct{}{}
	}
	key := map[string]int{}
	var zero, one, two, three, four, five, six, seven, eight, nine string
	for _, i := range input {
		switch len(i) {
		case 2:
			one = i
			key[one] = 1
		case 4:
			four = i
			key[four] = 4
		case 3:
			seven = i
			key[seven] = 7
		case 7:
			eight = i
			key[eight] = 8
		}
	}
	// know 1, 4, 7, and 8
	// 1 | 2 is non-existent, find 2
	// 1 | 6 is 8, find 6
	// 1 | 5 is 9, the only non-identity except 1 | 6 = 8
	for _, i := range input {
		if _, found := allDigits[add(one,i)]; !found {
			two = i
			key[two] = 2
		} else if i != eight && add(one,i) == eight {
			six = i
			key[six] = 6
		} else if add(one,i) != i && i != eight {
			five = i
			key[five] = 5
			nine = add(one,five)
			key[nine] = 9
		}
	}
	// we know them all except 0 and 3, which have a different number of segments
	for _, i := range []string{one,two,four,five,six,seven,eight,nine} {
		delete(allDigits,i)
	}
	for k, _ := range allDigits {
		if len(k) == 6 {
			zero = k
			key[zero] = 0
		} else if len(k) == 5 {
			three = k
			key[three] = 3
		}
	}
	fmt.Println(zero,one,two,three,four,five,six,seven,eight,nine)
	return key
}

// so much easier if i were using bitmasks...
func add(x, y string) string {
	sum := map[byte]struct{}{}
	for _, b := range []byte(x) {
		sum[b] = struct{}{}
	}
	for _, b := range []byte(y) {
		sum[b] = struct{}{}
	}
	bytes := []byte{}
	for b, _ := range sum {
		bytes = append(bytes, b)
	}
	return SortString(string(bytes))
}

func part1(scanner *bufio.Scanner) int {
	num := 0
	for {
		_, output := readLine(scanner)
		if output == nil {
			break
		}
		for _, o := range output {
			switch len(o) {
			case 2, 3, 4, 7:
				num++
			}
		}
	}
	return num
}

func nextWord(scanner *bufio.Scanner) (bool,string) {
	if scanner.Scan() {
		return true,SortString(scanner.Text())
	}
	return false, ""
}

func readLine(scanner *bufio.Scanner) ([]string,[]string) {
	var ok bool
	words := make([]string,14)
	for i := 0; i<10; i++ {
		ok, words[i] = nextWord(scanner)
		if !ok {
			return nil,nil
		}
	}
	if ok, delim := nextWord(scanner); !ok || delim != "|" {
		panic("missing delim")
	}
	for i := 10; i < 14; i++ {
		_, words[i] = nextWord(scanner)
	}
	return words[0:10], words[10:14]
}

