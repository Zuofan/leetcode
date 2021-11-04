package main

import (
	"bytes"
	"fmt"
	"math/big"
	"unsafe"
)

func main() {

	i := -1 // type int
	fmt.Printf("%d %x %d %x\n", i, i, uint(i), uint(i))
	i = -9
	fmt.Printf("uint64=%x uint32=%x\n", uint64(i), uint32(i))

	//something wrong
	fmt.Printf("binary=%x\n", *(*[8]byte)(unsafe.Pointer(&i)))

	fmt.Printf("twoComplement(%d)=%x\n", i, TwoComplement(int64(i)))
}


func TwoComplement(val int64) []byte {
	n := big.NewInt(val)

	var r []byte
	if n.Cmp(big.NewInt(0)) != -1 {
		r = n.Bytes()
	} else {
		mask := big.NewInt(1)
		mask.Lsh(mask, 64)

		r = n.Add(n, mask).Bytes()
	}

	res := bytes.NewBuffer([]byte{})
	for i := 0; i < 8-len(r); i++ {
		res.WriteByte(0)
	}
	res.Write(r)
	return res.Bytes()
}
