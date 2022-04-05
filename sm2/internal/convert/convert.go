package convert

import (
	"errors"
	"math/big"
)

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
