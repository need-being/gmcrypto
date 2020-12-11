package sm4

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
	panic("not implemented")
}
