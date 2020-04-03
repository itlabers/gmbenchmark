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

func BenchmarkVerify(b *testing.B) {
	msg := []byte("hello world")
	h := sha256.New()
	h.Write(msg)
	hash := h.Sum(nil)
	b.Run("ecdsa", func(b *testing.B) {
		priKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		r, s, err := ecdsa.Sign(rand.Reader, priKey, hash)
		pubKey := priKey.PublicKey
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ok := ecdsa.Verify(&pubKey, hash, r, s)
			if !ok {
				b.Fatal(err)
			}
		}
	})
	b.Run("mixbee-sm2", func(b *testing.B) {
		priKey, pubKey, _ := keypair.GenerateKeyPair(keypair.PK_SM2, keypair.SM2P256V1)
		sign, err := signature.Sign(signature.SM3withSM2, priKey, hash, "")

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ok := signature.Verify(pubKey, hash, sign)
			if !ok {
				b.Fatal(err)
			}
		}
	})

	b.Run("tjfoc-sm2", func(b *testing.B) {
		priKey, _ := tjfoc.GenerateKey()
		r, s, err := tjfoc.Sign(priKey, hash)
		pubKey := priKey.PublicKey
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ok := tjfoc.Verify(&pubKey, hash, r, s)
			if !ok {
				b.Fatal(err)
			}
		}
	})
	b.Run("flyinox-sm2", func(b *testing.B) {
		priKey, _ := flyinox.GenerateKey(rand.Reader)
		r, s, err := flyinox.Sign(rand.Reader, priKey, hash)
		pubKey := priKey.PublicKey
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ok := flyinox.Verify(&pubKey, hash, r, s)
			if !ok {
				b.Fatal(err)
			}
		}
	})
}
