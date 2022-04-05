// Package sm2 is implemented based on GB/T 32918.1-2016, GB/T 32918.2-2016, and
// GB/T 32918.5-2017.
package sm2

import (
	"crypto"
	"crypto/elliptic"
	"errors"
	"io"
	"math/big"

	"github.com/need-being/gmcrypto/sm2/internal/convert"
	"github.com/need-being/gmcrypto/sm3"
)

// precomputed big integers.
var (
	one   = new(big.Int).SetInt64(1)
	two   = new(big.Int).SetInt64(2)
	three = new(big.Int).SetInt64(3)
)

// PublicKey represents an SM2 public key.
type PublicKey struct {
	elliptic.Curve
	X, Y *big.Int

	// ID is the identifier of the signer.
	// The max size of an ID is 65,535 bytes.
	//
	// ID is not a part of the public key defined in GB/T 32918.2-2016.
	// Since it is commonly a good practice to use one key pair per identity,
	// it makes more sense to bind the ID with the public key.
	// In the scenario that a key pair is associated with multiple identities,
	// multiple instances of PublicKey or PrivateKey should be created for each
	// identity with other fields remaining the same.
	// This also enables SM2 to support crypto.Signer.
	ID []byte
}

// Equal reports whether pub and x have the same value.
func (pub *PublicKey) Equal(x crypto.PublicKey) bool {
	xx, ok := x.(*PublicKey)
	if !ok {
		return false
	}
	// check curve pointers only since SM2 curve is a singleton.
	return pub.Curve == xx.Curve && pub.X.Cmp(xx.X) == 0 && pub.Y.Cmp(xx.Y) == 0
}

// Digest returns the value Z defined in GB/T 32918.2-2016 by hashing.
// SM3 is used for hash algorithm.
func (pub *PublicKey) Digest() ([]byte, error) {
	idLength := len(pub.ID) << 3
	if idLength > 0xffff {
		return nil, errors.New("sm2: ID too large")
	}

	// Z = H( len(ID) || ID || a || b || Gx || Gy || X || Y )
	h := sm3.New() // write on sm3 never returns error

	// write ID
	h.Write([]byte{byte(idLength >> 8), byte(idLength)})
	h.Write(pub.ID)

	// write curve
	params := pub.Curve.Params()
	buf := make([]byte, (params.BitSize+7)/8)

	a := new(big.Int).Sub(params.P, three)
	if err := convert.FieldToBytes(a, params.P, buf); err != nil {
		return nil, err
	}
	h.Write(buf)
	if err := convert.FieldToBytes(params.B, params.P, buf); err != nil {
		return nil, err
	}
	h.Write(buf)
	if err := convert.FieldToBytes(params.Gx, params.P, buf); err != nil {
		return nil, err
	}
	h.Write(buf)
	if err := convert.FieldToBytes(params.Gy, params.P, buf); err != nil {
		return nil, err
	}
	h.Write(buf)

	// write public key
	if err := convert.FieldToBytes(pub.X, params.P, buf); err != nil {
		return nil, err
	}
	h.Write(buf)
	if err := convert.FieldToBytes(pub.Y, params.P, buf); err != nil {
		return nil, err
	}
	h.Write(buf)

	return h.Sum(nil), nil
}

// PrivateKey represents an SM2 private key.
type PrivateKey struct {
	PublicKey
	D *big.Int
}

// Public returns the public key corresponding to priv.
func (priv *PrivateKey) Public() crypto.PublicKey {
	return &priv.PublicKey
}

// Equal reports whether priv and x have the same value.
func (priv *PrivateKey) Equal(x crypto.PrivateKey) bool {
	xx, ok := x.(*PrivateKey)
	if !ok {
		return false
	}
	return priv.PublicKey.Equal(&xx.PublicKey) && priv.D.Cmp(xx.D) == 0
}

// Sign signs the given message with priv.
// SM2 relies on hash over message and the identity of the signer, and therefore
// cannot handle pre-hashed messages. Thus opts.HashFunc() must return zero to
// indicate the message hasn't been hashed if opts presents. This can be
// achieved by passing crypto.Hash(0) as the value for opts.
func (priv *PrivateKey) Sign(rand io.Reader, message []byte, opts crypto.SignerOpts) (signature []byte, err error) {
	if opts != nil && opts.HashFunc() != crypto.Hash(0) {
		return nil, errors.New("sm2: cannot sign hashed message")
	}

	return Sign(rand, priv, message)
}

// GenerateKey generates a public and private key pair.
func GenerateKey(c elliptic.Curve, rand io.Reader) (*PrivateKey, error) {
	// generate d in [1, n-2].
	params := c.Params()
	b := make([]byte, params.BitSize/8+8) // 64 more bits to reduce bias from mod.
	if _, err := io.ReadFull(rand, b); err != nil {
		return nil, err
	}
	d := new(big.Int).SetBytes(b)
	n := new(big.Int).Sub(params.N, two)
	d.Mod(d, n)
	d.Add(d, one)

	// compute public key
	x, y := c.ScalarBaseMult(d.Bytes())

	// pack private key
	return &PrivateKey{
		PublicKey: PublicKey{
			Curve: c,
			X:     x,
			Y:     y,
		},
		D: d,
	}, nil
}

// Sign signs the message with a private key and returns a signature.
// The signature is in the form of (r, s) where r and s have the same length.
// SM3 is used for hash algorithm.
func Sign(rand io.Reader, priv *PrivateKey, message []byte) ([]byte, error) {
	// A1, A2: compute hash value e
	z, err := priv.Digest()
	if err != nil {
		return nil, err
	}
	h := sm3.New() // write on sm3 never returns error
	h.Write(z)
	h.Write(message)
	e := convert.BytesToInteger(h.Sum(nil))

	// A3: generate random k
	params := priv.Curve.Params()
	r := new(big.Int)
	s := new(big.Int)
	for {
		b := make([]byte, params.BitSize/8+8) // 64 more bits to reduce bias from mod.
		if _, err = io.ReadFull(rand, b); err != nil {
			return nil, err
		}
		k := new(big.Int).SetBytes(b)
		n := new(big.Int).Sub(params.N, one)
		k.Mod(k, n)
		k.Add(k, one)

		// A4: compute (x, y) = kG where y is dropped
		x, _ := priv.Curve.ScalarBaseMult(k.Bytes())

		// A5: compute r = (e + x) mod n
		r.Add(e, x)
		r.Mod(r, params.N)
		if r.Sign() == 0 {
			continue // goto A3
		}
		if t := new(big.Int).Add(r, k); t.Cmp(params.N) == 0 {
			continue // goto A3
		}

		// A6: compute s = ((1 + d)^-1 * (k - rd)) mod n
		s.Add(one, priv.D)
		s.ModInverse(s, params.N)
		s2 := new(big.Int).Mul(r, priv.D)
		s2.Sub(k, s2)
		s.Mul(s, s2)
		s.Mod(s, params.N)
		if s.Sign() != 0 {
			break // goto A3
		}
	}

	// A7: convert r, s to byte strings
	n := (params.BitSize + 7) / 8
	sig := make([]byte, n*2)
	if err = convert.IntegerToBytes(r, sig[:n]); err != nil {
		return nil, err
	}
	if err = convert.IntegerToBytes(s, sig[n:]); err != nil {
		return nil, err
	}
	return sig, nil
}

// Verify reports whether sig is a valid signature of message by the given
// public key.
func Verify(pub *PublicKey, message, sig []byte) bool {
	// parse (r, s)
	params := pub.Curve.Params()
	n := (params.BitSize + 7) / 8
	if len(sig) != n*2 {
		return false
	}
	r := convert.BytesToInteger(sig[:n])
	s := convert.BytesToInteger(sig[n:])

	// B1: check r in [1, n-1]
	if r.Sign() == 0 || r.Cmp(params.N) >= 0 {
		return false
	}

	// B2: check s in [1, n-1]
	if s.Sign() == 0 || s.Cmp(params.N) >= 0 {
		return false
	}

	// B3, B4: compute hash value e.
	z, err := pub.Digest()
	if err != nil {
		return false
	}
	h := sm3.New() // write on sm3 never returns error
	h.Write(z)
	h.Write(message)
	e := convert.BytesToInteger(h.Sum(nil))

	// B5: compute t = (r + s) mod n
	t := new(big.Int).Add(r, s)
	t.Mod(t, params.N)
	if t.Sign() == 0 {
		return false
	}

	// B6: compute (x, y) = sG + tP where y is dropped
	x, y := pub.Curve.ScalarBaseMult(s.Bytes())
	x2, y2 := pub.Curve.ScalarMult(pub.X, pub.Y, t.Bytes())
	x, _ = pub.Curve.Add(x, y, x2, y2)

	// B7: compute R = (e + x) mod n
	rr := new(big.Int).Add(e, x)
	rr.Mod(rr, params.N)

	return rr.Cmp(r) == 0
}
