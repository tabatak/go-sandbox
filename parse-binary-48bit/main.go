package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	//32bit
	data32 := []byte{0x01, 0x02, 0x03, 0x04}

	//uint32にBigEndianで読み込む
	be32 := binary.BigEndian.Uint32(data32)
	fmt.Printf("BigEndian: %d\n", be32) //16909060

	//uint32にLittleEndianで読み込む
	le32 := binary.LittleEndian.Uint32(data32)
	fmt.Printf("LittleEndian: %d\n", le32) //67305985

	//48bit
	data48 := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06}

	//uint64にBigEndianで読み込む
	be48 := uint64(data48[5]) | uint64(data48[4])<<8 | uint64(data48[3])<<16 |
		uint64(data48[2])<<24 | uint64(data48[1])<<32 | uint64(data48[0])<<40
	fmt.Printf("BigEndian: %d\n", be48) //1108152157446

	//uint64にLittleEndianで読み込む
	le48 := uint64(data48[0]) | uint64(data48[1])<<8 | uint64(data48[2])<<16 |
		uint64(data48[3])<<24 | uint64(data48[4])<<32 | uint64(data48[5])<<40
	fmt.Printf("LittleEndian: %d\n", le48) //6618611909121

}
