package sm3

import (
	"bytes"
	"crypto/rand"
	"encoding"
	"fmt"
	"io"
	"testing"
)

type sm3Test struct {
	out string
	in  string
}

var golden = []sm3Test{
	{"66c7f0f462eeedd9d1f2d46bdc10e4e24167c4875cf2f7a2297da02b8f4ba8e0", "\x61\x62\x63"},
	{"debe9ff92275b8a138604889c18e5a4d6fdb70e5387e5765293dcba39c0c5732", "\x61\x62\x63\x64\x61\x62\x63\x64\x61\x62\x63\x64\x61\x62\x63\x64\x61\x62\x63\x64\x61\x62\x63\x64\x61\x62\x63\x64\x61\x62\x63\x64\x61\x62\x63\x64\x61\x62\x63\x64\x61\x62\x63\x64\x61\x62\x63\x64\x61\x62\x63\x64\x61\x62\x63\x64\x61\x62\x63\x64\x61\x62\x63\x64"},
}

func TestGolden(t *testing.T) {
	for i := 0; i < len(golden); i++ {
		g := golden[i]
		s := fmt.Sprintf("%x", Sum([]byte(g.in)))
		if s != g.out {
			t.Fatalf("Sum function: sm3(%s) = %s want %s", g.in, s, g.out)
		}
		c := New()
		for j := 0; j < 3; j++ {
			if j < 2 {
				io.WriteString(c, g.in)
			} else {
				io.WriteString(c, g.in[0:len(g.in)/2])
				c.Sum(nil)
				io.WriteString(c, g.in[len(g.in)/2:])
			}
			s := fmt.Sprintf("%x", c.Sum(nil))
			if s != g.out {
				t.Fatalf("sm3[%d](%s) = %s want %s", j, g.in, s, g.out)
			}
			c.Reset()
		}
	}
}

func TestGoldenMarshal(t *testing.T) {
	h := New()
	h2 := New()
	for _, g := range golden {
		h.Reset()
		h2.Reset()

		io.WriteString(h, g.in[:len(g.in)/2])

		state, err := h.(encoding.BinaryMarshaler).MarshalBinary()
		if err != nil {
			t.Errorf("could not marshal: %v", err)
			continue
		}

		if err := h2.(encoding.BinaryUnmarshaler).UnmarshalBinary(state); err != nil {
			t.Errorf("could not unmarshal: %v", err)
			continue
		}

		io.WriteString(h, g.in[len(g.in)/2:])
		io.WriteString(h2, g.in[len(g.in)/2:])

		if actual, actual2 := h.Sum(nil), h2.Sum(nil); !bytes.Equal(actual, actual2) {
			t.Errorf("sm3(%q) = 0x%x != marshaled 0x%x", g.in, actual, actual2)
		}
	}
}

func TestSize(t *testing.T) {
	c := New()
	if got := c.Size(); got != Size {
		t.Errorf("Size = %d; want %d", got, Size)
	}
}

func TestBlockSize(t *testing.T) {
	c := New()
	if got := c.BlockSize(); got != BlockSize {
		t.Errorf("BlockSize = %d; want %d", got, BlockSize)
	}
}

// Tests that blockGeneric (pure Go) and block (in assembly for some architectures) match.
func TestBlockGeneric(t *testing.T) {
	for i := 1; i < 30; i++ { // arbitrary factor
		gen, asm := New().(*digest), New().(*digest)
		buf := make([]byte, BlockSize*i)
		rand.Read(buf)
		blockGeneric(gen, buf)
		block(asm, buf)
		if *gen != *asm {
			t.Errorf("For %#v block and blockGeneric resulted in different states", buf)
		}
	}
}

var bench = New()
var buf = make([]byte, 8192)

func benchmarkSize(b *testing.B, size int) {
	b.SetBytes(int64(size))
	sum := make([]byte, bench.Size())
	for i := 0; i < b.N; i++ {
		bench.Reset()
		bench.Write(buf[:size])
		bench.Sum(sum[:0])
	}
}

func BenchmarkHash8Bytes(b *testing.B) {
	benchmarkSize(b, 8)
}

func BenchmarkHash320Bytes(b *testing.B) {
	benchmarkSize(b, 320)
}

func BenchmarkHash1K(b *testing.B) {
	benchmarkSize(b, 1024)
}

func BenchmarkHash8K(b *testing.B) {
	benchmarkSize(b, 8192)
}
