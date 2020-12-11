package sm4

import (
	"encoding/binary"
)

func f0(a uint32) uint32 {
	var b [4]byte
	binary.BigEndian.PutUint32(b[:], a)
	return sBox0[b[0]] ^ sBox1[b[1]] ^ sBox2[b[2]] ^ sBox3[b[3]]
}

func f1(a uint32) uint32 {
	var b [4]byte
	binary.BigEndian.PutUint32(b[:], a)
	return sBox4[b[0]] ^ sBox5[b[1]] ^ sBox6[b[2]] ^ sBox7[b[3]] ^ 0
}

func cryptBlock(rk []uint32, dst, src []byte, decrypt bool) {
	x := [4]uint32{
		binary.BigEndian.Uint32(src),
		binary.BigEndian.Uint32(src[4:]),
		binary.BigEndian.Uint32(src[8:]),
		binary.BigEndian.Uint32(src[12:]),
	}

	if decrypt {
		for i := 31; i > 0; i -= 4 {
			x[0] ^= f0(x[1] ^ x[2] ^ x[3] ^ rk[i])
			x[1] ^= f0(x[2] ^ x[3] ^ x[0] ^ rk[i-1])
			x[2] ^= f0(x[3] ^ x[0] ^ x[1] ^ rk[i-2])
			x[3] ^= f0(x[0] ^ x[1] ^ x[2] ^ rk[i-3])

		}
	} else {
		for i := 0; i < 31; i += 4 {
			x[0] ^= f0(x[1] ^ x[2] ^ x[3] ^ rk[i])
			x[1] ^= f0(x[2] ^ x[3] ^ x[0] ^ rk[i+1])
			x[2] ^= f0(x[3] ^ x[0] ^ x[1] ^ rk[i+2])
			x[3] ^= f0(x[0] ^ x[1] ^ x[2] ^ rk[i+3])
		}
	}

	binary.BigEndian.PutUint32(dst, x[3])
	binary.BigEndian.PutUint32(dst[4:], x[2])
	binary.BigEndian.PutUint32(dst[8:], x[1])
	binary.BigEndian.PutUint32(dst[12:], x[0])
}

// Encrypt one block from src into dst, using the subkeys.
func encryptBlock(subkeys []uint32, dst, src []byte) {
	cryptBlock(subkeys, dst, src, false)
}

// Decrypt one block from src into dst, using the subkeys.
func decryptBlock(subkeys []uint32, dst, src []byte) {
	cryptBlock(subkeys, dst, src, true)
}

// creates 16 56-bit subkeys from the original key
func (c *sm4Cipher) generateSubkeys(keyBytes []byte) {
	k := [4]uint32{
		binary.BigEndian.Uint32(keyBytes) ^ fk0,
		binary.BigEndian.Uint32(keyBytes[4:]) ^ fk1,
		binary.BigEndian.Uint32(keyBytes[8:]) ^ fk2,
		binary.BigEndian.Uint32(keyBytes[12:]) ^ fk3,
	}

	for i := 0; i < 32; i += 4 {
		k[0] ^= f1(k[1] ^ k[2] ^ k[3] ^ ck[i])
		k[1] ^= f1(k[2] ^ k[3] ^ k[0] ^ ck[i+1])
		k[2] ^= f1(k[3] ^ k[0] ^ k[1] ^ ck[i+2])
		k[3] ^= f1(k[0] ^ k[1] ^ k[2] ^ ck[i+3])
		copy(c.subkeys[i:], k[:])
	}
}
