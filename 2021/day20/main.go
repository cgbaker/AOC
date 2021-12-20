package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	numSteps := 50
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	algo, img := readInput(file)
	void := byte(0)
	for s := 0; s < numSteps; s++ {
		img = algo.enhance(img,void)
		if algo[0] != 0 {
			void ^= 1
		}
	}
	img.print()
	fmt.Println("num lit:",bytes.Count(img.pixels, []byte{1}))
}

type Algorithm []byte

func (a Algorithm) enhance(input *Image, void byte) *Image {
	sz := input.size+2
	output := &Image{
		size: sz,
		pixels: make([]byte,sz*sz),
	}
	for y := 0; y < output.size; y++ {
		for x := 0; x < output.size; x++ {
			idx := input.applyStencil(y-1,x-1,void)
			output.pixels[y*output.size + x] = a[idx]
		}
	}
	return output
}

type Image struct {
	size int
	pixels []byte
}

func (i *Image) append(str string) {
	if i.size == 0 {
		i.size = len(str)
		i.pixels = make([]byte, 0, i.size*i.size)
	}
	i.pixels = append(i.pixels, toBytes(str)...)
}

func (i *Image) applyStencil(y int, x int, void byte) int16 {
	acc := int16(0)
	acc += int16(i.get(y-1,x-1, void)) << 8
	acc += int16(i.get(y-1,x  , void)) << 7
	acc += int16(i.get(y-1,x+1, void)) << 6
	acc += int16(i.get(y  ,x-1, void)) << 5
	acc += int16(i.get(y  ,x  , void)) << 4
	acc += int16(i.get(y  ,x+1, void)) << 3
	acc += int16(i.get(y+1,x-1, void)) << 2
	acc += int16(i.get(y+1,x  , void)) << 1
	acc += int16(i.get(y+1,x+1, void)) << 0
	return acc
}

func (i *Image) get(y, x int, void byte) byte {
	if 0 <= y && y < i.size && 0 <= x && x < i.size {
		return i.pixels[i.size*y + x]
	}
	return void
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

func readInput(file *os.File) (Algorithm, *Image) {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	algo := toBytes(scanner.Text())
	image := &Image{}
	for scanner.Scan() {
		image.append(scanner.Text())
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