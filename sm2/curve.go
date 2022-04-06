package sm2

import (
	"crypto/elliptic"
	"math/big"
)

// curve represents an SM2 curve.
// The curve for SM2 is also a short-form Weierstrass curve with a=-3.
var curve = &elliptic.CurveParams{
	P: new(big.Int).SetBytes([]byte{
		0xff, 0xff, 0xff, 0xfe, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0x00, 0x00, 0x00, 0x00,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	}),
	N: new(big.Int).SetBytes([]byte{
		0xff, 0xff, 0xff, 0xfe, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0x72, 0x03, 0xdf, 0x6b, 0x21, 0xc6, 0x05, 0x2b,
		0x53, 0xbb, 0xf4, 0x09, 0x39, 0xd5, 0x41, 0x23,
	}),
	B: new(big.Int).SetBytes([]byte{
		0x28, 0xe9, 0xfa, 0x9e, 0x9d, 0x9f, 0x5e, 0x34,
		0x4d, 0x5a, 0x9e, 0x4b, 0xcf, 0x65, 0x09, 0xa7,
		0xf3, 0x97, 0x89, 0xf5, 0x15, 0xab, 0x8f, 0x92,
		0xdd, 0xbc, 0xbd, 0x41, 0x4d, 0x94, 0x0e, 0x93,
	}),
	Gx: new(big.Int).SetBytes([]byte{
		0x32, 0xc4, 0xae, 0x2c, 0x1f, 0x19, 0x81, 0x19,
		0x5f, 0x99, 0x04, 0x46, 0x6a, 0x39, 0xc9, 0x94,
		0x8f, 0xe3, 0x0b, 0xbf, 0xf2, 0x66, 0x0b, 0xe1,
		0x71, 0x5a, 0x45, 0x89, 0x33, 0x4c, 0x74, 0xc7,
	}),
	Gy: new(big.Int).SetBytes([]byte{
		0xbc, 0x37, 0x36, 0xa2, 0xf4, 0xf6, 0x77, 0x9c,
		0x59, 0xbd, 0xce, 0xe3, 0x6b, 0x69, 0x21, 0x53,
		0xd0, 0xa9, 0x87, 0x7c, 0xc6, 0x2a, 0x47, 0x40,
		0x02, 0xdf, 0x32, 0xe5, 0x21, 0x39, 0xf0, 0xa0,
	}),
	BitSize: 256,
	Name:    "SM2", // name from ISO 14888-3
}

// Curve returns an elliptic curve for SM2.
func Curve() elliptic.Curve {
	return curve
}
