package sm4

import (
	"testing"
)

type cryptTest struct {
	key []byte
	in  []byte
	out []byte
}

var encryptTests = []cryptTest{
	{
		// Appendix A.1 of GB/T 32907-2016
		[]byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10},
		[]byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10},
		[]byte{0x68, 0x1e, 0xdf, 0x34, 0xd2, 0x06, 0x96, 0x5e, 0x86, 0xb3, 0xe9, 0x4f, 0x53, 0x6e, 0x42, 0x46},
	},
}

var repeatedTests = []cryptTest{
	{
		// Appendix A.2 of GB/T 32907-2016
		[]byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10},
		[]byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10},
		[]byte{0x59, 0x52, 0x98, 0xc7, 0xc6, 0xfd, 0x27, 0x1f, 0x04, 0x02, 0xf8, 0x04, 0xc3, 0x3d, 0x3f, 0x66},
	},
}

func TestCipherEncrypt(t *testing.T) {
	for i, tt := range encryptTests {
		c, err := NewCipher(tt.key)
		if err != nil {
			t.Errorf("NewCipher(%d bytes) = %s", len(tt.key), err)
			continue
		}
		out := make([]byte, len(tt.in))
		c.Encrypt(out, tt.in)
		for j, v := range out {
			if v != tt.out[j] {
				t.Errorf("Cipher.Encrypt %d: out[%d] = %#x, want %#x", i, j, v, tt.out[j])
				break
			}
		}
	}
}

func TestCipherDecrypt(t *testing.T) {
	for i, tt := range encryptTests {
		c, err := NewCipher(tt.key)
		if err != nil {
			t.Errorf("NewCipher(%d bytes) = %s", len(tt.key), err)
			continue
		}
		plain := make([]byte, len(tt.in))
		c.Decrypt(plain, tt.out)
		for j, v := range plain {
			if v != tt.in[j] {
				t.Errorf("decryptBlock %d: plain[%d] = %#x, want %#x", i, j, v, tt.in[j])
				break
			}
		}
	}
}

func TestCipherEncryptRepeated(t *testing.T) {
	for i, tt := range repeatedTests {
		c, err := NewCipher(tt.key)
		if err != nil {
			t.Errorf("NewCipher(%d bytes) = %s", len(tt.key), err)
			continue
		}
		out := make([]byte, len(tt.in))
		copy(out, tt.in)
		for j := 0; j < 1000000; j++ {
			c.Encrypt(out, out)
		}
		for j, v := range out {
			if v != tt.out[j] {
				t.Errorf("Cipher.Encrypt %d: out[%d] = %#x, want %#x", i, j, v, tt.out[j])
				break
			}
		}
	}
}

func TestCipherDecryptRepeated(t *testing.T) {
	for i, tt := range repeatedTests {
		c, err := NewCipher(tt.key)
		if err != nil {
			t.Errorf("NewCipher(%d bytes) = %s", len(tt.key), err)
			continue
		}
		plain := make([]byte, len(tt.in))
		copy(plain, tt.out)
		for j := 0; j < 1000000; j++ {
			c.Decrypt(plain, plain)
		}
		for j, v := range plain {
			if v != tt.in[j] {
				t.Errorf("decryptBlock %d: plain[%d] = %#x, want %#x", i, j, v, tt.in[j])
				break
			}
		}
	}
}

func BenchmarkNewCipher(b *testing.B) {
	tt := encryptTests[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := NewCipher(tt.key)
		if err != nil {
			b.Fatal("NewCipher:", err)
		}
	}
}

func BenchmarkEncrypt(b *testing.B) {
	tt := encryptTests[0]
	c, err := NewCipher(tt.key)
	if err != nil {
		b.Fatal("NewCipher:", err)
	}
	out := make([]byte, len(tt.in))
	b.SetBytes(int64(len(out)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Encrypt(out, tt.in)
	}
}

func BenchmarkDecrypt(b *testing.B) {
	tt := encryptTests[0]
	c, err := NewCipher(tt.key)
	if err != nil {
		b.Fatal("NewCipher:", err)
	}
	out := make([]byte, len(tt.out))
	b.SetBytes(int64(len(out)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Decrypt(out, tt.out)
	}
}
