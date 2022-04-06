package sm2_test

import (
	"crypto/rand"
	"fmt"

	"github.com/need-being/gmcrypto/sm2"
)

func Example() {
	privateKey, err := sm2.GenerateKey(sm2.Curve(), rand.Reader)
	if err != nil {
		panic(err)
	}
	privateKey.ID = []byte("example ID")

	// Sign message
	message := []byte("hello world")
	signature, err := sm2.Sign(rand.Reader, privateKey, message)
	if err != nil {
		panic(err)
	}
	fmt.Println("message signed")

	// Verify message
	publicKey := &privateKey.PublicKey
	if sm2.Verify(publicKey, message, signature) {
		fmt.Println("message verified")
	}

	// Verify with tampered message
	message = []byte("foobar")
	if !sm2.Verify(publicKey, message, signature) {
		fmt.Println("verification failed as expected")
	}
	// Output:
	// message signed
	// message verified
	// verification failed as expected
}
