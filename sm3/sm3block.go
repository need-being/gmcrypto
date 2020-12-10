package sm3

import (
	"encoding/binary"
	"math/bits"
)

const (
	t0 = 0x79cc4519
	t1 = 0x7a879d8a
)

func ff0(x, y, z uint32) uint32 { return x ^ y ^ z }

func ff1(x, y, z uint32) uint32 { return (x & y) | (x & z) | (y & z) }

func gg0(x, y, z uint32) uint32 { return x ^ y ^ z }

func gg1(x, y, z uint32) uint32 { return (x & y) | (^x & z) }

func p0(x uint32) uint32 { return x ^ bits.RotateLeft32(x, 9) ^ bits.RotateLeft32(x, 17) }

func p1(x uint32) uint32 { return x ^ bits.RotateLeft32(x, 15) ^ bits.RotateLeft32(x, 23) }

// blockGeneric is a portable, pure Go version of the SM3 block step.
// It's used by sm3block_generic.go and tests.
func blockGeneric(dig *digest, p []byte) {
	var w [68]uint32

	h0, h1, h2, h3, h4, h5, h6, h7 := dig.h[0], dig.h[1], dig.h[2], dig.h[3], dig.h[4], dig.h[5], dig.h[6], dig.h[7]
	for len(p) >= chunk {
		for i := 0; i < 16; i++ {
			w[i] = binary.BigEndian.Uint32(p[i*4:])
		}
		for i := 16; i < 68; i++ {
			w[i] = p1(w[i-16]^w[i-9]^bits.RotateLeft32(w[i-3], 15)) ^ bits.RotateLeft32(w[i-13], 7) ^ w[i-6]
		}

		a, b, c, d, e, f, g, h := h0, h1, h2, h3, h4, h5, h6, h7
		for i := 0; i < 16; i++ {
			ss1 := bits.RotateLeft32(bits.RotateLeft32(a, 12)+e+bits.RotateLeft32(t0, i), 7)
			ss2 := ss1 ^ bits.RotateLeft32(a, 12)
			tt1 := ff0(a, b, c) + d + ss2 + (w[i] ^ w[i+4])
			tt2 := gg0(e, f, g) + h + ss1 + w[i]
			d = c
			c = bits.RotateLeft32(b, 9)
			b = a
			a = tt1
			h = g
			g = bits.RotateLeft32(f, 19)
			f = e
			e = p0(tt2)
		}
		for i := 16; i < 64; i++ {
			ss1 := bits.RotateLeft32(bits.RotateLeft32(a, 12)+e+bits.RotateLeft32(t1, i), 7)
			ss2 := ss1 ^ bits.RotateLeft32(a, 12)
			tt1 := ff1(a, b, c) + d + ss2 + (w[i] ^ w[i+4])
			tt2 := gg1(e, f, g) + h + ss1 + w[i]
			d = c
			c = bits.RotateLeft32(b, 9)
			b = a
			a = tt1
			h = g
			g = bits.RotateLeft32(f, 19)
			f = e
			e = p0(tt2)
		}

		h0 ^= a
		h1 ^= b
		h2 ^= c
		h3 ^= d
		h4 ^= e
		h5 ^= f
		h6 ^= g
		h7 ^= h

		p = p[chunk:]
	}

	dig.h[0], dig.h[1], dig.h[2], dig.h[3], dig.h[4], dig.h[5], dig.h[6], dig.h[7] = h0, h1, h2, h3, h4, h5, h6, h7
}
