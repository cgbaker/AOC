package main

import (
	"bufio"
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

	packet := readPacket(file)
	fmt.Println(packet)

	fmt.Println("Part1:", packet.sumOfVersions())
	// fmt.Println("Part2:", packet.???)
}

func readPacket(file *os.File) *Packet {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	bytes := []uint8{}
	scanner.Scan()
	for _, b := range scanner.Text() {
		bytes = append(bytes, fromHex(b))
	}
	return &Packet{
		bytes: bytes,
	}
}

type Packet struct {
	// i'm storing 4 bits into 8 bits... i may regret this later...
	bytes []uint8
}

func (p *Packet) getSubPackets() []Packet {
	return nil
}

func (p *Packet) sumOfVersions() int {
	sum := p.getVersion()
	for _, sub := range p.getSubPackets() {
		sum += sub.sumOfVersions()
	}
	return sum
}

func (p *Packet) getVersion() int {
	return 0
}

func fromHex(b int32) uint8 {
	if b >= '0' && b <= '9' {
		return uint8(b - '0')
	} else if b >= 'A' && b <= 'F' {
		return uint8(b - 'A' + 10)
	}
	panic(fmt.Sprintln("i don't know what i'm reading:",string(b)))
}
