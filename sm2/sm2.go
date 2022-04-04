// Package sm2 is implemented based on GB/T 32918.1-2016, GB/T 32918.2-2016, and
// GB/T 32918.5-2017.
package sm2

import (
	"crypto"
	"crypto/elliptic"
	"io"
	"math/big"
)

// PublicKey represents an SM2 public key.
type PublicKey struct {
	elliptic.Curve
	X, Y *big.Int
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

// PrivateKey represents an ECDSA private key.
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

// GenerateKey generates a public and private key pair.
func GenerateKey(c elliptic.Curve, rand io.Reader) (*PrivateKey, error) {
	panic("not implemented")
}
