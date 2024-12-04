package utils

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"strconv"
)

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func SignInt(x int) int {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	} else {
		return 0
	}
}

func Atoi(s string) int {
	x, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return x
}

// from: https://www.reddit.com/r/golang/comments/enzpes/comment/fe8q1lj
func SplitRegex(re *regexp.Regexp) bufio.SplitFunc {
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if loc := re.FindIndex(data); loc != nil {
			return loc[1], data[loc[0]:loc[1]], nil
		}
		if atEOF {
			return 0, nil, io.EOF
		}
		return 0, nil, nil
	}
}

type CharGrid struct {
	NumRows, NumCols int
	// row-wise array
	chars []byte 
}

func (g *CharGrid) GetChar(r, c int) byte {
	if r >= g.NumRows || r < 0 || c >=g.NumCols || c < 0 {
		return 0;
	}
	return g.chars[c*g.NumRows + r];
}

func ReadCharGrid(file *os.File) *CharGrid {
	grid := &CharGrid{
		NumCols: 0,
		NumRows: 0,
		chars:  []byte{},
	}
	lineScanner := bufio.NewScanner(file)
	lineScanner.Split(bufio.ScanLines)
	if lineScanner.Scan() {
		line := lineScanner.Text()
		grid.NumCols = len(line)
		if grid.NumCols > 0 {
			grid.NumRows++
			grid.chars = append(grid.chars, []byte(line)...)
		}
	}
	for lineScanner.Scan() {
		line := lineScanner.Text()
		if len(line) != grid.NumCols {
			panic("read unexpected line number")
		}
		grid.NumRows++
		grid.chars = append(grid.chars, []byte(line)...)
	}
	return grid
}
