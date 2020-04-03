package sm3

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	flyinox "github.com/flyinox/crypto/sm/sm2"
	"github.com/mixbee/mixbee-crypto/keypair"
	"github.com/mixbee/mixbee-crypto/signature"
	tjfoc "github.com/tjfoc/gmsm/sm2"
	"testing"
)

func BenchmarkSign(b *testing.B) {
	msg := []byte("hello world")
	h := sha256.New()
	h.Write(msg)
	hash := h.Sum(nil)
	b.Run("ecdsa", func(b *testing.B) {
		priKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _, err := ecdsa.Sign(rand.Reader, priKey, hash)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
	b.Run("mixbee-sm2", func(b *testing.B) {
		priKey, _, _ := keypair.GenerateKeyPair(keypair.PK_SM2, keypair.SM2P256V1)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, err := signature.Sign(signature.SM3withSM2, priKey, hash, "")
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("tjfoc-sm2", func(b *testing.B) {
		priKey, _ := tjfoc.GenerateKey()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _, err := tjfoc.Sign(priKey, hash)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
	b.Run("flyinox-sm2", func(b *testing.B) {
		priKey, _ := flyinox.GenerateKey(rand.Reader)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _, err := flyinox.Sign(rand.Reader, priKey, hash)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}
