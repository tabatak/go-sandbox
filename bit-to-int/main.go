package main

import "fmt"

func main() {

	var b byte = 0xAA //10101010
	var m byte = 0x3C //00111100

	fmt.Printf("b=%08b\n", b)
	fmt.Printf("m=%08b\n", m)
	fmt.Printf("b&m=%08b\n", b&m)

	// 最下位ビットのインデックスを0として、2～5ビットの値をuint8として取り出す

	// 対象ビットだけをマスクして残す
	// 10101010 & 00111100 = 00101000
	var i = uint8(b & m)

	// 不要な下位ビットをシフトして落とす
	// 00101000 >> 2 → 00001010
	i = i >> 2
	fmt.Printf("i=%d\n", i)

	var b2 byte = 0xAA //10101010
	fmt.Printf("%08b\n", b2)

	// 両側にシフトして必要部分だけを取り出す
	// まずは、上位ビットを落とす
	var tmp = b2 << 2
	fmt.Printf("%08b\n", tmp)

	// 下位ビットを落とす(先に左に2ビットシフトした分を加算)
	tmp = tmp >> (2 + 2)
	fmt.Printf("%08b\n", tmp)

	// uint8で取り出す
	i2 := uint8(tmp)
	fmt.Printf("%d\n", i2)

}
