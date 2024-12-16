package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"slices"
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

type Coord interface {
	Row() int
	Col() int
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
	Chars []byte 
}

var (
	DIR_N Point = Point{-1,0}
	DIR_S Point = Point{ 1,0}
	DIR_W Point = Point{0,-1}
	DIR_E Point = Point{0, 1}
	DIRS []Point = []Point{DIR_N,DIR_E,DIR_S,DIR_W}
)

type Point struct {
	r, c int
}
func (c *Point) Row() int {
	return c.r
}
func (c *Point) Col() int {
	return c.c
}
func (a Point) Minus(b Point) Point {
	return Point{
		r: a.r - b.r,
		c: a.c - b.c,
	}
}
func (a Point) Plus(b Point) Point {
	return Point{
		r: a.r + b.r,
		c: a.c + b.c,
	}
}

func NewPoint(r, c int) Point {
	return Point{
		r: r,
		c: c,
	}
}

func NewPointFromCoord(c Coord) Point {
	return NewPoint(c.Row(), c.Col())
}

func NewCoord(r, c int) Coord {
	return &Point{
		r: r,
		c: c,
	}
}

func (g *CharGrid) Index(r, c int) int {
	if r >= g.NumRows || r < 0 || c >=g.NumCols || c < 0 {
		return -1
	}
	return r*g.NumCols + c
}

func (g *CharGrid) Point(index int) Point {
	r, c := g.RowCol(index)
	return Point{
		r: r,
		c: c,
	}
}

func (g *CharGrid) RowCol(index int) (int, int) {
	return index/g.NumCols, index%g.NumCols
}

func (g *CharGrid) GetCharFromIndex(index int) byte {
	if index >= len(g.Chars) || index < 0 {
		return 0
	}
	return g.Chars[index]
}

func (g *CharGrid) GetChar(coord Coord) byte {
	r, c := coord.Row(), coord.Col()
	if r >= g.NumRows || r < 0 || c >=g.NumCols || c < 0 {
		return 0;
	}
	return g.Chars[r*g.NumCols + c];
}

func (g *CharGrid) Print() {
	ln := g.Chars
	for range g.NumRows {
		fmt.Println(string(ln[:g.NumCols]))
		ln = ln[g.NumCols:]
	}
}

func (g *CharGrid) SetChar(coord Coord, val byte) {
	r, c := coord.Row(), coord.Col()
	if r >= g.NumRows || r < 0 || c >=g.NumCols || c < 0 {
		return
	}
	g.Chars[r*g.NumCols + c] = val;
}

func (g *CharGrid) SetCharFromIndex(index int, val byte) {
	if index < 0 || index > len(g.Chars) {
		return
	}
	g.Chars[index] = val;
}


func (g *CharGrid) NSEWPoints(p Point) []Point {
	neighbors := make([]Point,0,4)
	for _, d := range DIRS {
		neighbors = append(neighbors, p.Plus(d))
	}
	return neighbors
}

func (g *CharGrid) NSEW(index int) []int {
	neighbors := make([]int,0,4)
	p := g.Point(index)
	for _, n := range g.NSEWPoints(p) {
		neighbors = append(neighbors, g.Index(n.r,n.c))
	}
	return neighbors
}

func (g *CharGrid) Clone() *CharGrid {
	clone := &CharGrid{
		NumCols: g.NumCols,
		NumRows: g.NumRows,
		Chars: slices.Clone(g.Chars),
	}
	return clone
}

func NewCharGrid(r, c int) *CharGrid {
	return &CharGrid{
		NumCols: c, 
		NumRows: r,
		Chars: make([]byte, r*c),
	}
}

func ReadCharGrid(file *os.File) (*CharGrid, *bufio.Scanner) {
	grid := &CharGrid{
		NumCols: 0,
		NumRows: 0,
		Chars:  []byte{},
	}
	lineScanner := bufio.NewScanner(file)
	lineScanner.Split(bufio.ScanLines)
	if lineScanner.Scan() {
		line := lineScanner.Text()
		grid.NumCols = len(line)
		if grid.NumCols > 0 {
			grid.NumRows++
			grid.Chars = append(grid.Chars, []byte(line)...)
		}
	}
	for lineScanner.Scan() {
		line := lineScanner.Text()
		if len(line) == 0 {
			break
		}
		if len(line) != grid.NumCols {
			panic("read unexpected line number")
		}
		grid.NumRows++
		grid.Chars = append(grid.Chars, []byte(line)...)
	}
	return grid, lineScanner
}

func GetMiddle(x []int) int {
	if len(x) == 0 {
		panic("undefined")
	}
	return x[len(x)/2]
}

func Pow(a, b int) int {
	res := 1
	for range b {
		res *= a
	}
	return res
}
