package sm4

import (
	"bytes"
	"errors"
	flyinox "github.com/flyinox/crypto/sm/sm4"
	mixbee "github.com/mixbee/mixbee-crypto/sm4"
	tjfoc "github.com/tjfoc/gmsm/sm4"
	"testing"
)

func BenchmarkDecrypt(b *testing.B) {
	key := []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10}
	cipher := []byte{0x68, 0x1e, 0xdf, 0x34, 0xd2, 0x06, 0x96, 0x5e, 0x86, 0xb3, 0xe9, 0x4f, 0x53, 0x6e, 0x42, 0x46}
	expected := []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10}
	b.ResetTimer()
	b.Run("mixbee-sm4", func(b *testing.B) {
		block, _ := mixbee.NewCipher(key)
		out := make([]byte, len(cipher))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			block.Decrypt(out, cipher)
			if bytes.Compare(out, expected) != 0 {
				b.Fatal(errors.New("not march"))
			}
		}
	})
	b.Run("tjfoc-sm4", func(b *testing.B) {
		block, _ := tjfoc.NewCipher(key)
		out := make([]byte, len(cipher))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			block.Decrypt(out, cipher)
			if bytes.Compare(out, expected) != 0 {
				b.Fatal(errors.New("not march"))
			}
		}
	})
	b.Run("flyinox-sm4", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = flyinox.Sm4Ecb(key, cipher, flyinox.DEC)
		}
	})
}
