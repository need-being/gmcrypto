package sm2

import "testing"

func Test_Curve(t *testing.T) {
	if !curve.IsOnCurve(curve.Gx, curve.Gy) {
		t.Errorf("curve.IsOnCurve() = %v, want %v", false, true)
	}
}
