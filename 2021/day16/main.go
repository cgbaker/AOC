package main

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"fmt"
	bitstream "github.com/bearmini/bitstream-go"
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
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		str := scanner.Text()
		if len(str) < 80 {
			fmt.Println("Input:",str)
		}
		if len(str) % 2 != 0 {
			str += "0"
		}
		byts, err := hex.DecodeString(str)
		if err != nil {
			panic("parsing: " + err.Error())
		}
		r := bitstream.NewReader(bytes.NewReader(byts), nil)
		packet, _ := readPacket(r)
		fmt.Println("Part1:", sumOfVersions(packet))
		fmt.Println("Part2:", packet.value())
	}
}

type Packet struct {
	version int
	typeId  PacketType
	literal int64
	subPackets []*Packet
}

type PacketType int
const (
	SUM      PacketType = 0
	PRODUCT  PacketType = 1
	MINIMUM  PacketType = 2
	MAXIMUM  PacketType = 3
	LITERAL  PacketType = 4
	GRTRTHAN PacketType = 5
	LESSTHAN PacketType = 6
	EQUALTO  PacketType = 7
)

func readPacket(r *bitstream.Reader) (*Packet, int) {
	p := &Packet{
		version: readNBits(r, 3),
		typeId:  PacketType(readNBits(r, 3)),
	}
	bitsRead := 6
	if p.typeId == LITERAL {
		numNibbles := 0
		for {
			numNibbles++
			next := readNBits(r, 5)
			bitsRead += 5
			p.literal = (p.literal << 4) + int64(next & 0b1111)
			if next & 0b10000 == 0 {
				break
			} else if numNibbles >= 16 {
				panic("nibbled too much; need to use something bigger than int64")
			}
		}
	} else {
		lengthType, _ := r.ReadBit()
		bitsRead++
		switch lengthType {
		case 0:
			length := readNBits(r, 15)
			bitsRead += 15
			for {
				if length == 0 {
					break
				} else if length < 0 {
					panic("i read too much")
				}
				subP, read := readPacket(r)
				length -= read
				bitsRead += read
				p.subPackets = append(p.subPackets, subP)
			}
		case 1:
			length := readNBits(r, 11)
			bitsRead += 11
			for i := 0; i < length; i++ {
				subP, read := readPacket(r)
				bitsRead += read
				p.subPackets = append(p.subPackets, subP)
			}
		}
	}
	return p, bitsRead
}

func (p *Packet) value() int64 {
	value := int64(0)
	switch p.typeId {
	case SUM:
		for _, p := range p.subPackets {
			value += p.value()
		}
	case PRODUCT:
		value = 1
		for _, p := range p.subPackets {
			value *= p.value()
		}
	case MINIMUM:
		value = math.MaxInt
		for _, p := range p.subPackets {
			value = min(value, p.value())
		}
	case MAXIMUM:
		for _, p := range p.subPackets {
			value = max(value, p.value())
		}
	case LITERAL:
		value = p.literal
	case GRTRTHAN:
		if p.subPackets[0].value() > p.subPackets[1].value() {
			value = 1
		}
	case LESSTHAN:
		if p.subPackets[0].value() < p.subPackets[1].value() {
			value = 1
		}
	case EQUALTO:
		if p.subPackets[0].value() == p.subPackets[1].value() {
			value = 1
		}
	default:
		panic("i'm sorry, sir, this is a wendy's")
	}
	return value
}

func sumOfVersions(p *Packet) int {
	sum := p.version
	for _, sub := range p.subPackets {
		sum += sumOfVersions(sub)
	}
	return sum
}

func readNBits(r *bitstream.Reader, n uint8) int {
	i, err := r.ReadNBitsAsUint32BE(n)
	if err != nil {
		panic(err.Error())
	}
	return int(i)
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}