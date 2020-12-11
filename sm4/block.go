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

func cryptBlock(subkeys []uint32, dst, src []byte, decrypt bool) {
	panic("not implemented")
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
	var k [36]uint32

	k[0] = binary.BigEndian.Uint32(keyBytes) ^ fk0
	k[1] = binary.BigEndian.Uint32(keyBytes[4:]) ^ fk1
	k[2] = binary.BigEndian.Uint32(keyBytes[8:]) ^ fk2
	k[3] = binary.BigEndian.Uint32(keyBytes[12:]) ^ fk3

	for i := 0; i < 32; i++ {
		k[i+4] = k[i] ^ f1(k[i+1]^k[i+2]^k[i+3]^ck[i])
	}
	copy(c.subkeys[:], k[4:])
}
