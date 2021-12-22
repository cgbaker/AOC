package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"math"
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

	//// MAGNITUDE CHECK
	//for scanner.Scan() {
	//	n := readNumber(scanner)
	//	fmt.Println("magnitude:",n.magnitude(0))
	//}

	// PART 1
	//sum := &Number{
	//	tree: nil,
	//}
	//num := 0
	//for scanner.Scan() {
	//	next := readNumber(scanner)
	//	num++
	//	sum.add(next)
	//	sum = sum.reduce()
	//}
	//fmt.Print("sum:")
	//sum.print(0)
	//fmt.Println(sum.magnitude(0))

	// PART 2
	nums := []*Number{}
	for scanner.Scan() {
		nums = append(nums, readNumber(scanner))
	}
	m := 0
	for i, n1 := range nums {
		for _, n2 := range nums[i+1:] {
			c := n1.copy()
			c.add(n2)
			c = c.reduce()
			m = max(m, c.magnitude(0))

			c = n2.copy()
			c.add(n1)
			c = c.reduce()
			m = max(m, c.magnitude(0))
		}
	}
	fmt.Println(m)
}

const (
	PAIR = int8(-1)
	NOT_FOUND = int(-1)
)

type Number struct {
	tree []int8
}

func (n *Number) add(o *Number) {
	if n.tree == nil {
		n.tree = o.tree
		return
	}
	oldSize := max(treelen(n.tree), treelen(o.tree))
	newSize := 2*oldSize + 1
	if len(n.tree) < newSize {
		oldTree := n.tree
		n.tree = make([]int8, newSize)
		copy(n.tree, oldTree)
	}
	breadth := oldSize + 1
	start := oldSize
	// copy from the bottom up
	for breadth >= 1 {
		fromIdx := parent(start)
		if fromIdx < len(n.tree) {
			copy(n.tree[start:], n.tree[fromIdx:fromIdx+breadth/2])
		}
		if fromIdx < len(o.tree) {
			copy(n.tree[start+breadth/2:], o.tree[fromIdx:fromIdx+breadth/2])
		}
		breadth /= 2
		start -= breadth
	}
	n.tree[0] = PAIR
}

func treelen(tree []int8) int {
	l := 0
	for i, v := range tree {
		if v == PAIR {
			l = max(l,right(i))
		}
	}
	return 2 << int(math.Floor(math.Log2(float64(l+1)))) - 1
}

func (n *Number) magnitude(loc int) int {
	if n == nil {
		return 0
	}
	if loc >= len(n.tree) {
		panic(fmt.Sprint("magnitude arg out-of-bounds",loc))
	}
	switch v := n.tree[loc]; v {
	case PAIR:
		return 3*n.magnitude(left(loc)) + 2*n.magnitude(right(loc))
	default:
		return int(v)
	}
}

func (n *Number) insert(loc int, val int8) *Number {
	if loc >= len(n.tree) {
		fmt.Println("growing tree",loc)
		newTree := make([]int8, loc+1)
		copy(newTree, n.tree)
		n.tree = newTree
	}
	n.tree[loc] = val
	return n
}

func (n *Number) reduce() *Number {
	r := n
	for {
		if loc := r.findLeftmostDeepPair(); loc != NOT_FOUND {
			r.explode(loc)
		} else if loc := r.findLeftmostBigNumber(); loc != NOT_FOUND {
			r = r.split(loc)
		} else {
			break
		}
	}
	return r
}

func (n *Number) explode(loc int) {
	l, r := n.tree[left(loc)], n.tree[right(loc)]
	if l == PAIR || r == PAIR {
		panic("they said this wouldn't happen")
	}
	n.tree[loc] = 0

	foundMe := false
	leftNeighbor := NOT_FOUND
	rightNeighbor := NOT_FOUND
	n.dfsFind(0, func(idx int, val int8) bool {
		if idx == loc {
			foundMe = true
			return false
		}
		if !foundMe && val != PAIR {
			leftNeighbor = idx
		} else if foundMe && val != PAIR && rightNeighbor == NOT_FOUND {
			rightNeighbor = idx
		}
		return false
	})
	if leftNeighbor != NOT_FOUND {
		n.tree[leftNeighbor] += l
	}
	if rightNeighbor != NOT_FOUND {
		n.tree[rightNeighbor] += r
	}
}

func (n *Number) findLeftmostDeepPair() int {
	isDeepPair := func(l int, v int8) bool {
		return l >= 15 && v == PAIR
	}
	return n.dfsFind(0, isDeepPair)
}

func (n *Number) findLeftmostBigNumber() int {
	isBig := func(i int, v int8) bool {
		return v >= 10
	}
	return n.dfsFind(0, isBig)
}

func (n *Number) split(loc int) *Number {
	num := n.tree[loc]
	if num == PAIR {
		panic("wtf?")
	}
	l, r := (num >> 1), (num >> 1 + num % 2)
	n.insert(loc, PAIR)
	n.insert(left(loc), l)
	n.insert(right(loc), r)
	return n
}

func (n *Number) print(loc int) {
	switch v := n.tree[loc]; v {
	case PAIR:
		fmt.Print("[")
		n.print(left(loc))
		fmt.Print(",")
		n.print(right(loc))
		fmt.Print("]")
	default:
		fmt.Print(v)
	}
	if loc == 0 {
		fmt.Println("")
	}
}

func (n *Number) dfsFind(loc int, match func(i int, v int8) bool) int {
	if match(loc, n.tree[loc]) {
		return loc
	} else if n.tree[loc] == PAIR {
		if f := n.dfsFind(left(loc), match); f != NOT_FOUND {
			return f
		} else if f := n.dfsFind(right(loc), match); f != NOT_FOUND {
			return f
		}
	}
	return NOT_FOUND
}

func (n *Number) copy() *Number {
	c := &Number{
		tree: make([]int8, len(n.tree)),
	}
	copy(c.tree, n.tree)
	return c
}

func readNumber(scanner *bufio.Scanner) *Number {
	str := scanner.Text()
	maxDepth := 0
	curDepth := 0
	for _, b := range str {
		if b == '[' {
			curDepth++
			maxDepth = max(maxDepth, curDepth)
		} else if b == ']' {
			curDepth--
		}
	}
	n := &Number{
		tree: make([]int8, 1 << (maxDepth+1) - 1),
	}
	reader := bytes.NewReader([]byte(str))
	return recursiveRead(reader, n, 0)
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func recursiveRead(reader *bytes.Reader, num *Number, loc int) *Number {
	t, _ := reader.ReadByte()
	switch t {
	case '[':
		num = num.insert(loc, PAIR)
		num = recursiveRead(reader, num, left(loc))
		if t, err := reader.ReadByte(); t != ',' || err != nil {
			panic(fmt.Sprint(t, err))
		}
		num = recursiveRead(reader, num, right(loc))
		if t, err := reader.ReadByte(); t != ']' || err != nil {
			panic(fmt.Sprint(t, err))
		}
	default:
		num = num.insert(loc, int8(t - '0'))
	}
	return num
}

func left(loc int) int {
	return loc*2+1
}

func right(loc int) int {
	return loc*2+2
}

func parent(loc int) int {
	return (loc-1)/2
}