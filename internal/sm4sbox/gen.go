package main

import (
	"encoding/binary"
	"fmt"
	"math/big"
	"math/bits"
)

func xorBit(a uint8, ai int, b uint8, bi int) uint8 {
	ab := big.NewInt(int64(a))
	bb := big.NewInt(int64(b))
	ab.SetBit(ab, ai, ab.Bit(ai)^bb.Bit(bi))
	return uint8(ab.Uint64())
}

func cipherT(a uint8) uint32 {
	b := sBox[a]
	x := []byte{byte(bits.RotateLeft8(b, 6)), byte(b), byte(b), byte(b)}
	x[1] = xorBit(x[1], 7, b, 1)
	x[1] = xorBit(x[1], 6, b, 0)
	x[3] = xorBit(x[3], 5, b, 7)
	x[3] = xorBit(x[3], 4, b, 6)
	x[3] = xorBit(x[3], 3, b, 5)
	x[3] = xorBit(x[3], 2, b, 4)
	x[3] = xorBit(x[3], 1, b, 3)
	x[3] = xorBit(x[3], 0, b, 2)
	r := binary.BigEndian.Uint32(x)
	return bits.RotateLeft32(r, 26)
}

func keyT(a uint8) uint32 {
	b := uint32(sBox[a])
	return b<<24 | b<<15 | b<<5
}

func genSBox(transform func(uint8) uint32, start int) {
	for i := 0; i < 4; i++ {
		if i != 0 {
			fmt.Println()
		}
		fmt.Printf("var sBox%d = [256]uint32{\n", start+i)
		for j := 0; j < 16; j++ {
			fmt.Print("\t")
			for k := 0; k < 16; k++ {
				if k != 0 {
					fmt.Print(" ")
				}
				a := uint8(j<<4 | k)
				b := bits.RotateLeft32(transform(a), -i*8)
				fmt.Printf("0x%08x,", b)
			}
			fmt.Println()
		}
		fmt.Println("}")
	}
}

func main() {
	genSBox(cipherT, 0)
	genSBox(keyT, 4)
}
