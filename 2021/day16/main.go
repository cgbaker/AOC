package main

import (
	"bufio"
	"encoding/hex"
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
	bytes := []byte{}
	scanner.Scan()
	str := scanner.Text()
	if len(str) % 2 != 0 {
		str += "0"
	}
	bytes, err := hex.DecodeString(str)
	if err != nil {
		panic("parsing: " + err.Error())
	}
	return &Packet{
		bytes: bytes,
	}
}

type Packet struct {
	version int
	typeId  int
	subPackets []Packet
}

type Type byte
const (
	LITERAL Type = 4
)

func (p *Packet) getSubPackets() []Packet {
	if p.getType() == LITERAL {
		return nil
	}
	//var subPacketBytes []byte
	lenType := p.getLengthType()
	length := 0
	shift := 0
	switch lenType {
	case 0:
		// need 15 == 1 + 8 + 6
		length = int(0b00000001 & p.bytes[0]) << 14
		length += int(p.bytes[1]) << 6
		length += int(p.bytes[2] & 0b11111100) >> 2
		shift = 22
	case 1:
		// need 11 == 1 + 8 + 2
		length = int(0b00000001 & p.bytes[0]) << 10
		length += int(p.bytes[1]) << 2
		length += int(p.bytes[2] & 0b11000000) >> 6
		shift = 18
	}
	subPacketBytes := shiftBytes(p.bytes, shift)
	return dividePackets(subPacketBytes, lenType, length)
}

func shiftBytes(bytes []byte, shift int) []byte {
	shiftBytes := shift / 8
	shiftBits := shift % 8
	trunc := bytes[shiftBytes:]
	if shiftBits == 0 {
		return trunc
	}
	shifted := make([]byte,0,len(trunc))
	high := 0
	low := 0
	for i, t := range trunc {
		high = ((int(t) << shiftBits) & 0xff00) >> 8
		if i > 0 {
			if high | low > 255 {
				panic("too big")
			}
			shifted = append(shifted, byte(high | low))
		}
		low = int(t) << shiftBits & 255
	}
	shifted = append(shifted, byte(low))
	return shifted
}

func (p *Packet) sumOfVersions() int {
	sum := int(p.getVersion())
	for _, sub := range p.getSubPackets() {
		sum += sub.sumOfVersions()
	}
	return sum
}

func (p *Packet) getVersion() byte {
	return (p.bytes[0] & 0b11100000) >> 5
}

func (p *Packet) getType() Type {
	return Type(p.bytes[0] & 0b00011100 >> 2)
}

func (p *Packet) getLengthType() byte {
	if p.getType() == LITERAL {
		panic("no length type")
	}
	return (p.bytes[0] & 0b00000010) >> 1
}