package sm2

import (
	"bytes"
	"crypto/elliptic"
	"crypto/rand"
	"io"
	"math/big"
	"reflect"
	"testing"
)

func TestGenerateKey(t *testing.T) {
	tests := []struct {
		name    string
		c       elliptic.Curve
		rand    io.Reader
		want    *PrivateKey
		wantErr bool
	}{
		{
			name: "GB/T 32918.5-2017 A.2",
			c:    Curve(),
			rand: io.MultiReader(
				bytes.NewReader([]byte{
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // zeros
					0x39, 0x45, 0x20, 0x8f, 0x7b, 0x21, 0x44, 0xb1,
					0x3f, 0x36, 0xe3, 0x8a, 0xc6, 0xd3, 0x9f, 0x95,
					0x88, 0x93, 0x93, 0x69, 0x28, 0x60, 0xb5, 0x1a,
					0x42, 0xfb, 0x81, 0xef, 0x4d, 0xf7, 0xc5, 0xb7, // minus 1
				}),
				rand.Reader,
			),
			want: &PrivateKey{
				PublicKey: PublicKey{
					Curve: curve,
					X: new(big.Int).SetBytes([]byte{
						0x09, 0xf9, 0xdf, 0x31, 0x1e, 0x54, 0x21, 0xa1,
						0x50, 0xdd, 0x7d, 0x16, 0x1e, 0x4b, 0xc5, 0xc6,
						0x72, 0x17, 0x9f, 0xad, 0x18, 0x33, 0xfc, 0x07,
						0x6b, 0xb0, 0x8f, 0xf3, 0x56, 0xf3, 0x50, 0x20,
					}),
					Y: new(big.Int).SetBytes([]byte{
						0xcc, 0xea, 0x49, 0x0c, 0xe2, 0x67, 0x75, 0xa5,
						0x2d, 0xc6, 0xea, 0x71, 0x8c, 0xc1, 0xaa, 0x60,
						0x0a, 0xed, 0x05, 0xfb, 0xf3, 0x5e, 0x08, 0x4a,
						0x66, 0x32, 0xf6, 0x07, 0x2d, 0xa9, 0xad, 0x13,
					}),
				},
				D: new(big.Int).SetBytes([]byte{
					0x39, 0x45, 0x20, 0x8f, 0x7b, 0x21, 0x44, 0xb1,
					0x3f, 0x36, 0xe3, 0x8a, 0xc6, 0xd3, 0x9f, 0x95,
					0x88, 0x93, 0x93, 0x69, 0x28, 0x60, 0xb5, 0x1a,
					0x42, 0xfb, 0x81, 0xef, 0x4d, 0xf7, 0xc5, 0xb8,
				}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateKey(tt.c, tt.rand)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSign(t *testing.T) {
	tests := []struct {
		name    string
		rand    io.Reader
		priv    *PrivateKey
		message []byte
		want    []byte
		wantErr bool
	}{
		{
			name: "GB/T 32918.5-2017 A.2",
			rand: io.MultiReader(
				bytes.NewReader([]byte{
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // zeros
					0x59, 0x27, 0x6e, 0x27, 0xd5, 0x06, 0x86, 0x1a,
					0x16, 0x68, 0x0f, 0x3a, 0xd9, 0xc0, 0x2d, 0xcc,
					0xef, 0x3c, 0xc1, 0xfa, 0x3c, 0xdb, 0xe4, 0xce,
					0x6d, 0x54, 0xb8, 0x0d, 0xea, 0xc1, 0xbc, 0x20, // minus 1
				}),
				rand.Reader,
			),
			priv: &PrivateKey{
				PublicKey: PublicKey{
					Curve: curve,
					X: new(big.Int).SetBytes([]byte{
						0x09, 0xf9, 0xdf, 0x31, 0x1e, 0x54, 0x21, 0xa1,
						0x50, 0xdd, 0x7d, 0x16, 0x1e, 0x4b, 0xc5, 0xc6,
						0x72, 0x17, 0x9f, 0xad, 0x18, 0x33, 0xfc, 0x07,
						0x6b, 0xb0, 0x8f, 0xf3, 0x56, 0xf3, 0x50, 0x20,
					}),
					Y: new(big.Int).SetBytes([]byte{
						0xcc, 0xea, 0x49, 0x0c, 0xe2, 0x67, 0x75, 0xa5,
						0x2d, 0xc6, 0xea, 0x71, 0x8c, 0xc1, 0xaa, 0x60,
						0x0a, 0xed, 0x05, 0xfb, 0xf3, 0x5e, 0x08, 0x4a,
						0x66, 0x32, 0xf6, 0x07, 0x2d, 0xa9, 0xad, 0x13,
					}),
					ID: []byte{
						0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38,
						0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38,
					},
				},
				D: new(big.Int).SetBytes([]byte{
					0x39, 0x45, 0x20, 0x8f, 0x7b, 0x21, 0x44, 0xb1,
					0x3f, 0x36, 0xe3, 0x8a, 0xc6, 0xd3, 0x9f, 0x95,
					0x88, 0x93, 0x93, 0x69, 0x28, 0x60, 0xb5, 0x1a,
					0x42, 0xfb, 0x81, 0xef, 0x4d, 0xf7, 0xc5, 0xb8,
				}),
			},
			message: []byte{
				0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x20,
				0x64, 0x69, 0x67, 0x65, 0x73, 0x74,
			},
			want: []byte{
				// r
				0xf5, 0xa0, 0x3b, 0x06, 0x48, 0xd2, 0xc4, 0x63,
				0x0e, 0xea, 0xc5, 0x13, 0xe1, 0xbb, 0x81, 0xa1,
				0x59, 0x44, 0xda, 0x38, 0x27, 0xd5, 0xb7, 0x41,
				0x43, 0xac, 0x7e, 0xac, 0xee, 0xe7, 0x20, 0xb3,

				// s
				0xb1, 0xb6, 0xaa, 0x29, 0xdf, 0x21, 0x2f, 0xd8,
				0x76, 0x31, 0x82, 0xbc, 0x0d, 0x42, 0x1c, 0xa1,
				0xbb, 0x90, 0x38, 0xfd, 0x1f, 0x7f, 0x42, 0xd4,
				0x84, 0x0b, 0x69, 0xc4, 0x85, 0xbb, 0xc1, 0xaa,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Sign(tt.rand, tt.priv, tt.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sign() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVerify(t *testing.T) {
	tests := []struct {
		name    string
		pub     *PublicKey
		message []byte
		sig     []byte
		want    bool
	}{
		{
			name: "GB/T 32918.5-2017 A.2",
			pub: &PublicKey{
				Curve: curve,
				X: new(big.Int).SetBytes([]byte{
					0x09, 0xf9, 0xdf, 0x31, 0x1e, 0x54, 0x21, 0xa1,
					0x50, 0xdd, 0x7d, 0x16, 0x1e, 0x4b, 0xc5, 0xc6,
					0x72, 0x17, 0x9f, 0xad, 0x18, 0x33, 0xfc, 0x07,
					0x6b, 0xb0, 0x8f, 0xf3, 0x56, 0xf3, 0x50, 0x20,
				}),
				Y: new(big.Int).SetBytes([]byte{
					0xcc, 0xea, 0x49, 0x0c, 0xe2, 0x67, 0x75, 0xa5,
					0x2d, 0xc6, 0xea, 0x71, 0x8c, 0xc1, 0xaa, 0x60,
					0x0a, 0xed, 0x05, 0xfb, 0xf3, 0x5e, 0x08, 0x4a,
					0x66, 0x32, 0xf6, 0x07, 0x2d, 0xa9, 0xad, 0x13,
				}),
				ID: []byte{
					0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38,
					0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38,
				},
			},
			message: []byte{
				0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x20,
				0x64, 0x69, 0x67, 0x65, 0x73, 0x74,
			},
			sig: []byte{
				// r
				0xf5, 0xa0, 0x3b, 0x06, 0x48, 0xd2, 0xc4, 0x63,
				0x0e, 0xea, 0xc5, 0x13, 0xe1, 0xbb, 0x81, 0xa1,
				0x59, 0x44, 0xda, 0x38, 0x27, 0xd5, 0xb7, 0x41,
				0x43, 0xac, 0x7e, 0xac, 0xee, 0xe7, 0x20, 0xb3,

				// s
				0xb1, 0xb6, 0xaa, 0x29, 0xdf, 0x21, 0x2f, 0xd8,
				0x76, 0x31, 0x82, 0xbc, 0x0d, 0x42, 0x1c, 0xa1,
				0xbb, 0x90, 0x38, 0xfd, 0x1f, 0x7f, 0x42, 0xd4,
				0x84, 0x0b, 0x69, 0xc4, 0x85, 0xbb, 0xc1, 0xaa,
			},
			want: true,
		},
		{
			name: "GB/T 32918.5-2017 A.2: ID Changed",
			pub: &PublicKey{
				Curve: curve,
				X: new(big.Int).SetBytes([]byte{
					0x09, 0xf9, 0xdf, 0x31, 0x1e, 0x54, 0x21, 0xa1,
					0x50, 0xdd, 0x7d, 0x16, 0x1e, 0x4b, 0xc5, 0xc6,
					0x72, 0x17, 0x9f, 0xad, 0x18, 0x33, 0xfc, 0x07,
					0x6b, 0xb0, 0x8f, 0xf3, 0x56, 0xf3, 0x50, 0x20,
				}),
				Y: new(big.Int).SetBytes([]byte{
					0xcc, 0xea, 0x49, 0x0c, 0xe2, 0x67, 0x75, 0xa5,
					0x2d, 0xc6, 0xea, 0x71, 0x8c, 0xc1, 0xaa, 0x60,
					0x0a, 0xed, 0x05, 0xfb, 0xf3, 0x5e, 0x08, 0x4a,
					0x66, 0x32, 0xf6, 0x07, 0x2d, 0xa9, 0xad, 0x13,
				}),
				ID: []byte{
					0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38,
					0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0xff,
				},
			},
			message: []byte{
				0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x20,
				0x64, 0x69, 0x67, 0x65, 0x73, 0x74,
			},
			sig: []byte{
				// r
				0xf5, 0xa0, 0x3b, 0x06, 0x48, 0xd2, 0xc4, 0x63,
				0x0e, 0xea, 0xc5, 0x13, 0xe1, 0xbb, 0x81, 0xa1,
				0x59, 0x44, 0xda, 0x38, 0x27, 0xd5, 0xb7, 0x41,
				0x43, 0xac, 0x7e, 0xac, 0xee, 0xe7, 0x20, 0xb3,

				// s
				0xb1, 0xb6, 0xaa, 0x29, 0xdf, 0x21, 0x2f, 0xd8,
				0x76, 0x31, 0x82, 0xbc, 0x0d, 0x42, 0x1c, 0xa1,
				0xbb, 0x90, 0x38, 0xfd, 0x1f, 0x7f, 0x42, 0xd4,
				0x84, 0x0b, 0x69, 0xc4, 0x85, 0xbb, 0xc1, 0xaa,
			},
			want: false,
		},
		{
			name: "GB/T 32918.5-2017 A.2: message tampered",
			pub: &PublicKey{
				Curve: curve,
				X: new(big.Int).SetBytes([]byte{
					0x09, 0xf9, 0xdf, 0x31, 0x1e, 0x54, 0x21, 0xa1,
					0x50, 0xdd, 0x7d, 0x16, 0x1e, 0x4b, 0xc5, 0xc6,
					0x72, 0x17, 0x9f, 0xad, 0x18, 0x33, 0xfc, 0x07,
					0x6b, 0xb0, 0x8f, 0xf3, 0x56, 0xf3, 0x50, 0x20,
				}),
				Y: new(big.Int).SetBytes([]byte{
					0xcc, 0xea, 0x49, 0x0c, 0xe2, 0x67, 0x75, 0xa5,
					0x2d, 0xc6, 0xea, 0x71, 0x8c, 0xc1, 0xaa, 0x60,
					0x0a, 0xed, 0x05, 0xfb, 0xf3, 0x5e, 0x08, 0x4a,
					0x66, 0x32, 0xf6, 0x07, 0x2d, 0xa9, 0xad, 0x13,
				}),
				ID: []byte{
					0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38,
					0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38,
				},
			},
			message: []byte{
				0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x20,
				0x64, 0x69, 0x67, 0x65, 0x73, 0xff,
			},
			sig: []byte{
				// r
				0xf5, 0xa0, 0x3b, 0x06, 0x48, 0xd2, 0xc4, 0x63,
				0x0e, 0xea, 0xc5, 0x13, 0xe1, 0xbb, 0x81, 0xa1,
				0x59, 0x44, 0xda, 0x38, 0x27, 0xd5, 0xb7, 0x41,
				0x43, 0xac, 0x7e, 0xac, 0xee, 0xe7, 0x20, 0xb3,

				// s
				0xb1, 0xb6, 0xaa, 0x29, 0xdf, 0x21, 0x2f, 0xd8,
				0x76, 0x31, 0x82, 0xbc, 0x0d, 0x42, 0x1c, 0xa1,
				0xbb, 0x90, 0x38, 0xfd, 0x1f, 0x7f, 0x42, 0xd4,
				0x84, 0x0b, 0x69, 0xc4, 0x85, 0xbb, 0xc1, 0xaa,
			},
			want: false,
		},
		{
			name: "GB/T 32918.5-2017 A.2: signature corrupted",
			pub: &PublicKey{
				Curve: curve,
				X: new(big.Int).SetBytes([]byte{
					0x09, 0xf9, 0xdf, 0x31, 0x1e, 0x54, 0x21, 0xa1,
					0x50, 0xdd, 0x7d, 0x16, 0x1e, 0x4b, 0xc5, 0xc6,
					0x72, 0x17, 0x9f, 0xad, 0x18, 0x33, 0xfc, 0x07,
					0x6b, 0xb0, 0x8f, 0xf3, 0x56, 0xf3, 0x50, 0x20,
				}),
				Y: new(big.Int).SetBytes([]byte{
					0xcc, 0xea, 0x49, 0x0c, 0xe2, 0x67, 0x75, 0xa5,
					0x2d, 0xc6, 0xea, 0x71, 0x8c, 0xc1, 0xaa, 0x60,
					0x0a, 0xed, 0x05, 0xfb, 0xf3, 0x5e, 0x08, 0x4a,
					0x66, 0x32, 0xf6, 0x07, 0x2d, 0xa9, 0xad, 0x13,
				}),
				ID: []byte{
					0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38,
					0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38,
				},
			},
			message: []byte{
				0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x20,
				0x64, 0x69, 0x67, 0x65, 0x73, 0x74,
			},
			sig: []byte{
				// r
				0xf5, 0xa0, 0x3b, 0x06, 0x48, 0xd2, 0xc4, 0x63,
				0x0e, 0xea, 0xc5, 0x13, 0xe1, 0xbb, 0x81, 0xa1,
				0x59, 0x44, 0xda, 0x38, 0x27, 0xd5, 0xb7, 0x41,
				0x43, 0xac, 0x7e, 0xac, 0xee, 0xe7, 0x20, 0xff,

				// s
				0xb1, 0xb6, 0xaa, 0x29, 0xdf, 0x21, 0x2f, 0xd8,
				0x76, 0x31, 0x82, 0xbc, 0x0d, 0x42, 0x1c, 0xa1,
				0xbb, 0x90, 0x38, 0xfd, 0x1f, 0x7f, 0x42, 0xd4,
				0x84, 0x0b, 0x69, 0xc4, 0x85, 0xbb, 0xc1, 0xff,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Verify(tt.pub, tt.message, tt.sig); got != tt.want {
				t.Errorf("Verify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSign(b *testing.B) {
	priv, err := GenerateKey(Curve(), rand.Reader)
	if err != nil {
		b.Fatal("GenerateKey:", err)
	}
	priv.ID = []byte("benchmark")
	message := []byte("message")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err = Sign(rand.Reader, priv, message)
		if err != nil {
			b.Fatal("Sign:", err)
		}
	}
}

func BenchmarkVerify(b *testing.B) {
	priv, err := GenerateKey(Curve(), rand.Reader)
	if err != nil {
		b.Fatal("GenerateKey:", err)
	}
	priv.ID = []byte("benchmark")
	message := []byte("message")
	sig, err := Sign(rand.Reader, priv, message)
	if err != nil {
		b.Fatal("Sign:", err)
	}
	pub := &priv.PublicKey
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if !Verify(pub, message, sig) {
			b.Fatal("Verify failed")
		}
	}
}

func BenchmarkVerifyFailed(b *testing.B) {
	priv, err := GenerateKey(Curve(), rand.Reader)
	if err != nil {
		b.Fatal("GenerateKey:", err)
	}
	priv.ID = []byte("benchmark")
	message := []byte("message")
	sig, err := Sign(rand.Reader, priv, message)
	if err != nil {
		b.Fatal("Sign:", err)
	}
	pub := &priv.PublicKey

	// tamper message
	message = []byte("malicious")

	// benchmark
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if Verify(pub, message, sig) {
			b.Fatal("Verify should faile")
		}
	}
}