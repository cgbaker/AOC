package main

import (
	"container/list"
	"fmt"
	"os"
	"log"
	"bufio"
	"github.com/cgbaker/AOC/2024/utils"
	"slices"
)

type Space = int

type Pair struct {File;Space}

type File struct {
	id int
	size int
	moved bool
}

type Problem = []Pair

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	problem := readInput(file)
	prob1(problem)
	prob2(problem)
}

func prob1(problem Problem) {
	clone := slices.Clone(problem)	
	fmt.Printf("prob1: %d\n",computeChecksum1(clone))
}

func prob2(problem Problem) {
	fmt.Printf("prob2: %d\n",computeChecksum2(problem))
}

func computeChecksum1(problem Problem) int {
	checksum := 0
	pos := 0
	outer:
	for i, p := range problem {
		// accumulate the file contribution
		for range p.File.size {
			checksum += p.File.id * pos
			pos++
		}
		// now file the empty space with a file from the back
		for range p.Space {
			last := lastNonEmptyFile(problem[i+1:])
			if last == nil {
				break outer
			}
			checksum += last.File.id * pos
			last.File.size--
			pos++
		}
	}
	return checksum
}

func lastNonEmptyFile(pairs []Pair) *Pair {
	for j := len(pairs)-1; j>=0; j-- {
		if pairs[j].File.size > 0 {
			return &pairs[j]
		}
	}
	return nil
}

func computeChecksum2(problem Problem) int {
	// I don't see a streaming solution for this one, gonna actually have to implement the compaction
	// because I have to compact in order, attempting each file only once, I don't see an obvious streaming solution
	list := list.New()
	for _, p := range problem {
		list.PushBack(p)
	}
	// compactify: back to front
	for e := list.Back(); e != list.Front(); e = e.Prev() {
		file := e.Value.(Pair).File
		if file.moved {
			continue
		}
		// look for a blank space
		for f := list.Front(); f != e; f = f.Next() {
			if file.size <= f.Value.(Pair).Space {
				file.moved = true
				list.InsertAfter( Pair{file, f.Value.(Pair).Space - file.size}, f )
				f.Value = Pair{f.Value.(Pair).File, 0}
				e.Value = Pair{File{}, e.Value.(Pair).Space + file.size}
				break
			}
		}
	}
	// compute checksum
	checksum := 0
	pos := 0
	for e := list.Front(); e != nil; e = e.Next() {
		file := e.Value.(Pair).File
		for range file.size {
			checksum += file.id * pos
			pos++
		}
		pos += e.Value.(Pair).Space
	}
	return checksum
}

func Print(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		for range e.Value.(Pair).File.size {
			fmt.Printf("%d",e.Value.(Pair).File.id)
		}
		for range e.Value.(Pair).Space {
			fmt.Printf(".")
		}
	}
	fmt.Println("")
}


func readInput(file *os.File) Problem {
	pairs := []Pair{}
	byteScanner := bufio.NewScanner(file)
	byteScanner.Split(bufio.ScanBytes)
	parsingFile := true
	curFileId := 0
	for byteScanner.Scan() {
		c := byteScanner.Text()
		if c[0] == '\n' {
			break
		}
		sz := utils.Atoi(c)
		if parsingFile {
			newFile := File{id: curFileId, size: sz}
			pairs = append(pairs, Pair{newFile,0})
			curFileId++
		}  else {
			pairs[len(pairs)-1].Space = sz
		}
		parsingFile = !parsingFile
	}
	return pairs
}


