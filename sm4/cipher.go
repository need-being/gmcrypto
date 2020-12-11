// Package sm4 is implemented based on GB/T 32907-2016.
package sm4

import (
	"crypto/cipher"
	"strconv"
)

// BlockSize is the SM4 block size in bytes.
const BlockSize = 8

// KeySizeError indicates invalid key size
type KeySizeError int

func (k KeySizeError) Error() string {
	return "crypto/sm4: invalid key size " + strconv.Itoa(int(k))
}

// sm4Cipher is an instance of SM4 encryption.
type sm4Cipher struct {
	subkeys [32]uint32
}

// NewCipher creates and returns a new cipher.Block.
func NewCipher(key []byte) (cipher.Block, error) {
	if len(key) != 16 {
		return nil, KeySizeError(len(key))
	}

	c := new(sm4Cipher)
	c.generateSubkeys(key)
	return c, nil
}

func (c *sm4Cipher) BlockSize() int { return BlockSize }

func (c *sm4Cipher) Encrypt(dst, src []byte) {
	if len(src) < BlockSize {
		panic("crypto/sm4: input not full block")
	}
	if len(dst) < BlockSize {
		panic("crypto/sm4: output not full block")
	}
	encryptBlock(c.subkeys[:], dst, src)
}

func (c *sm4Cipher) Decrypt(dst, src []byte) {
	if len(src) < BlockSize {
		panic("crypto/sm4: input not full block")
	}
	if len(dst) < BlockSize {
		panic("crypto/sm4: output not full block")
	}
	decryptBlock(c.subkeys[:], dst, src)
}
