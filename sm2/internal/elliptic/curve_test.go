package sm2

import (
	"math/big"
	"testing"
)

func Test_curve_IsOnCurve(t *testing.T) {
	tests := []struct {
		name  string
		curve *curve
		x     *big.Int
		y     *big.Int
		want  bool
	}{
		{
			name:  "generator",
			curve: instance,
			x:     instance.Gx,
			y:     instance.Gy,
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.curve.IsOnCurve(tt.x, tt.y); got != tt.want {
				t.Errorf("curve.IsOnCurve() = %v, want %v", got, tt.want)
			}
		})
	}
}
