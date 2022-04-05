package convert

import (
	"errors"
	"math/big"
)

// IntegerToBytes defined by GB/T 32918.1-2016 4.2.2
func IntegerToBytes(x *big.Int, buf []byte) error {
	if x.Sign() < 0 {
		return errors.New("negative integer")
	}
	if len(x.Bits()) > len(buf) {
		return errors.New("integer too large")
	}
	x.FillBytes(buf)
	return nil
}

// BytesToInteger defined by GB/T 32918.1-2016 4.2.3.
func BytesToInteger(b []byte) *big.Int {
	return new(big.Int).SetBytes(b)
}

// FieldToBytes defined by GB/T 32918.1-2016 4.2.6.
func FieldToBytes(x, q *big.Int, buf []byte) error {
	if x.Sign() < 0 {
		return errors.New("negative integer")
	}
	if x.Cmp(q) >= 0 {
		return errors.New("integer not in the field")
	}
	if len(q.Bits()) > len(buf) {
		return errors.New("buffer too small")
	}
	x.FillBytes(buf)
	return nil
}
