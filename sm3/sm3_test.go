package sm3

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"testing"

	flyinox "github.com/flyinox/crypto/sm/sm3"
	mixbee "github.com/mixbee/mixbee-crypto/sm3"
	tjfoc "github.com/tjfoc/gmsm/sm3"
)

func BenchmarkSM3(b *testing.B) {
	msg := []byte("abc")
	hash, _ := hex.DecodeString("66c7f0f462eeedd9d1f2d46bdc10e4e24167c4875cf2f7a2297da02b8f4ba8e0")

	sha256Hash, _ := hex.DecodeString("ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad")
	b.ResetTimer()
	b.Run("openssl-sm3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			h := New()
			h.Write(msg)
			hashed := h.Sum(nil)
			if bytes.Compare(hash, hashed) != 0 {
				b.Fatal(errors.New("not march"))
			}
		}
	})
	b.Run("sha256", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			h := sha256.New()
			h.Write(msg)
			hashed := h.Sum(nil)
			if bytes.Compare(sha256Hash, hashed) != 0 {
				b.Fatal(errors.New("not march"))
			}
		}
	})

	b.Run("mixbee-sm3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			h := mixbee.New()
			h.Write(msg)
			hashed := h.Sum(nil)
			if bytes.Compare(hash, hashed) != 0 {
				b.Fatal(errors.New("not march"))
			}
		}
	})
	b.Run("tjfoc-sm3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			h := tjfoc.New()
			h.Write(msg)
			hashed := h.Sum(nil)
			if bytes.Compare(hash, hashed) != 0 {
				b.Fatal(errors.New("not march"))
			}
		}
	})
	b.Run("flyinox-sm3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			h := flyinox.New()
			h.Write(msg)
			hashed := h.Sum(nil)
			if bytes.Compare(hash, hashed) != 0 {
				b.Fatal(errors.New("not march"))
			}
		}
	})
}
