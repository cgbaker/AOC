package main

import (
	"bufio"
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
	algo, img := readInput(file, numSteps)
	for s := 0; s < numSteps; s++ {
		img = algo.enhance(img)
	}
	img.print()
	count := 0
	for y := 2; y < img.size-2; y++ {
		for _, p := range img.pixels[y*img.size + 2 : y*img.size + img.size - 2] {
			count += int(p)
		}
	}
	fmt.Println("num lit:",count)
}

type Algorithm []byte

func (a Algorithm) enhance(input *Image) *Image {
	output := &Image{
		border: input.border,
		size: input.size,
		pixels: make([]byte,input.size*input.size),
	}
	for y := 1; y < output.size-1; y++ {
		for x := 1; x < output.size-1; x++ {
			idx := input.applyStencil(y,x)
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

func (i *Image) applyStencil(y int, x int) int16 {
	acc := int16(0)
	idx := (y-1)*i.size + (x-1)
	acc += int16(i.pixels[idx+0]) << 8
	acc += int16(i.pixels[idx+1]) << 7
	acc += int16(i.pixels[idx+2]) << 6
	idx += i.size
	acc += int16(i.pixels[idx+0]) << 5
	acc += int16(i.pixels[idx+1]) << 4
	acc += int16(i.pixels[idx+2]) << 3
	idx += i.size
	acc += int16(i.pixels[idx+0]) << 2
	acc += int16(i.pixels[idx+1]) << 1
	acc += int16(i.pixels[idx+2]) << 0
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