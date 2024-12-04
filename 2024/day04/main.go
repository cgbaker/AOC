package main

import (
	"fmt"
	"os"
	"log"
	"github.com/cgbaker/AOC/2024/utils"
)

type Input = utils.CharGrid

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	input := utils.ReadCharGrid(file)
	prob1(input)
	prob2(input)
}

func prob1(input *Input) {
	numSols := 0
	for r := range input.NumRows {
		for c := range input.NumCols {
			numSols += checkSol(r,c,input)
		}
	}
	fmt.Printf("prob1: %d\n",numSols)
}

func prob2(input *Input) {
	numSols := 0
	for r := 1; r < input.NumRows-1; r++ {
		for c := 1; c < input.NumCols-1; c++ {
			if checkMasSol(r,c,input) {
				numSols++
			}
		}
	}
	fmt.Printf("prob2: %d\n",numSols)
}

func checkSol(r, c int, input *Input) int {
	numSols := 0
	dirs := [][2]int {
		{-1, -1},
		{-1,  0},
		{-1,  1},
		{ 0, -1},
		{ 0,  1},
		{ 1, -1},
		{ 1,  0},
		{ 1,  1},
	}
	for _, dir := range dirs {
		if checkSolDir(r,c,input,dir) {
			numSols++
		}
	}
	return numSols
}

func checkSolDir(r, c int, input *Input, dir [2]int) bool {
	sol := []byte{'X','M','A','S'}
	for i := range 4 {
		if input.GetChar(r, c) != sol[i] {
			return false
		}
		r += dir[0]
		c += dir[1]
	}
	return true
}

func checkMasSol(r, c int, input *Input) bool {
	if input.GetChar(r,c) != 'A' {
		return false
	}
	NW := input.GetChar(r-1, c-1)
	NE := input.GetChar(r+1, c-1)
	SW := input.GetChar(r-1, c+1)
	SE := input.GetChar(r+1, c+1)
	check := string([]byte{NW,NE,SW,SE})
	switch check {
	case "MMSS", "MSMS", "SSMM", "SMSM":
		return true
	default:
		return false
	}
}
