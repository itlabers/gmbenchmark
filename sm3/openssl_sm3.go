package sm3

/*
#cgo LDFLAGS: -lcrypto -L"/usr/local/opt/openssl@1.1/lib"
#cgo CPPFLAGS: -I"/usr/local/opt/openssl@1.1/include"

#include "openssl/evp.h"
*/
import "C"
import (
	"hash"
)
import "unsafe"

type digest struct {
	ctx *C.EVP_MD_CTX
}

func (d *digest) Write(p []byte) (n int, err error) {
	r := C.EVP_DigestUpdate(d.ctx, unsafe.Pointer(&p[0]), C.ulong(len(p)))
	return int(r), nil
}

func New() hash.Hash {
	d := new(digest)
	d.ctx = C.EVP_MD_CTX_new()
	d.Reset()
	return d
}

func (d *digest) Sum(b []byte) []byte {
	if len(b) == 0 {
		b = make([]byte, d.Size())
	}
	var size C.uint
	C.EVP_DigestFinal(d.ctx, (*C.uchar)(&b[0]), &size)
	defer C.EVP_MD_CTX_free(d.ctx)
	return b
}

// Reset resets the Hash to its initial state.
func (d *digest) Reset() {
	C.EVP_DigestInit(d.ctx, C.EVP_sm3())
}

// Size returns the number of bytes Sum will return.
func (*digest) Size() int {
	return 32
}

// BlockSize returns the hash's underlying block size.
// The Write method must be able to accept any amount
// of data, but it may operate more efficiently if all writes
// are a multiple of the block size.
func (*digest) BlockSize() int {
	return 64
}
