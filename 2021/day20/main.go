package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	numSteps := 2
	file, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	algo, img := readInput(file, 5)
	img.print()
	for s := 0; s < numSteps; s++ {
		img = algo.enhance(img)
		img.print()
	}
	fmt.Println("num lit:",bytes.Count(img.pixels, []byte{1}))
}

type Algorithm []byte

func (a Algorithm) enhance(input *Image) *Image {
	if input.border <= 0 {
		panic("not enough room")
	}
	output := &Image{
		border: input.border-1,
		size: input.size,
		pixels: make([]byte,input.size*input.size),
	}
	for y := output.border; y < output.size-output.border; y++ {
		for x := output.border; x < output.size-output.border; x++ {
			idx := input.applyStencil(y,x)
			//fmt.Printf("y: %d, x: %d, idx: %d\n",y,x,idx)
			output.pixels[y*output.size + x] = a[idx]
		}
	}
	return output
}

type Image struct {
	border int
	size int
	pixels []byte
}

func (i *Image) append(row int, str string) {
	if i.size == 0 {
		i.size = len(str)+2*i.border
		i.pixels = make([]byte, i.size*i.size)
	}
	start := row*i.size + i.border
	copy(i.pixels[start:], toBytes(str))
}

func (i *Image) applyStencil(y int, x int) int {
	acc := 0
	idx := (y-1)*i.size + (x-1)
	acc += int(i.pixels[idx+0]) << 8
	acc += int(i.pixels[idx+1]) << 7
	acc += int(i.pixels[idx+2]) << 6
	idx += i.size
	acc += int(i.pixels[idx+0]) << 5
	acc += int(i.pixels[idx+1]) << 4
	acc += int(i.pixels[idx+2]) << 3
	idx += i.size
	acc += int(i.pixels[idx+0]) << 2
	acc += int(i.pixels[idx+1]) << 1
	acc += int(i.pixels[idx+2]) << 0
	return acc
}

func (img *Image) print() {
	for i, p := range img.pixels {
		switch p {
		case 0:
			fmt.Print(".")
		case 1:
			fmt.Print("#")
		}
		if ((i+1) % img.size) == 0 {
			fmt.Println("")
		}
	}
	fmt.Println("")
}

func readInput(file *os.File, border int) (Algorithm, *Image) {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	algo := toBytes(scanner.Text())
	image := &Image{
		border: border,
	}
	atRow := border
	for scanner.Scan() {
		image.append(atRow, scanner.Text())
		atRow++
	}
	return algo, image
}

func toBytes(str string) []byte {
	bts := make([]byte,len(str))
	for i, s := range str {
		switch s {
		case '.':
			bts[i] = 0
		case '#':
			bts[i] = 1
		}
	}
	return bts
}