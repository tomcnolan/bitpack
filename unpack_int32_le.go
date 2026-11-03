//go:build 386 || amd64 || amd64p32 || alpha || arm || arm64 || loong64 || mipsle || mips64le || mips64p32le || nios2 || ppc64le || riscv || riscv64 || sh || wasm

package bitpack

import "github.com/parquet-go/bitpack/unsafecast"

func unsafecastBytesToUint32(src []byte) []uint32 {
	return unsafecast.Slice[uint32](src)
}
