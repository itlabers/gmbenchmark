package sm3

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	flyinox "github.com/flyinox/crypto/sm/sm2"
	"github.com/mixbee/mixbee-crypto/keypair"
	tjfoc "github.com/tjfoc/gmsm/sm2"
	"testing"
)

func BenchmarkECDSAGenKeyPair(b *testing.B) {
	b.Run("ecdsa", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
	b.Run("mixbee-sm2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _, err := keypair.GenerateKeyPair(keypair.PK_SM2, keypair.SM2P256V1)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
	b.Run("tjfoc-sm2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := tjfoc.GenerateKey()
			if err != nil {
				b.Fatal(err)
			}
		}
	})
	b.Run("flyinox-sm2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := flyinox.GenerateKey(rand.Reader)
			if err != nil {
				b.Fatal(err)
			}
		}
	})

}
