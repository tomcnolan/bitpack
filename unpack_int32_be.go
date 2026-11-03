//go:build armbe || arm64be || m68k || mips || mips64 || mips64p32 || ppc || ppc64 || s390 || s390x || shbe || sparc || sparc64

package bitpack

import "encoding/binary"

func unsafecastBytesToUint32(src []byte) []uint32 {
	out := make([]uint32, len(src)/4)
	idx := 0
	for k := range out {
		out[k] = binary.LittleEndian.Uint32((src)[idx:(4 + idx)])
		idx += 4
	}
	return out
}
