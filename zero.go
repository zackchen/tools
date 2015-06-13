package main

import (
	"encoding/binary"
	"github.com/deepglint/muses/util/io"
)

func main() {
	var i int16 = 0
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, uint16(i))
	// a := 0
	io.CreateFileByBytes("zero.dat", b)

}
