package main

import (
	"fmt"
	"reflect"
)

func main() {

	// 最上位が1の場合
	b := 0x81
	fmt.Printf("%d\n", b)
	fmt.Printf("%08b\n", b)
	// 7bitシフト
	fmt.Printf("%08b\n", b>>7)
	// 8bitシフト
	fmt.Printf("%08b\n", b>>8)
	fmt.Printf("type of b = %s\n", reflect.TypeOf(b))

	// マイナスのint
	i := int8(-10)
	fmt.Printf("%08b\n", byte(i))

	i2 := int8(-100)
	fmt.Printf("%d\n", i2)
	b2 := byte(i2)
	fmt.Printf("%08b\n", b2)
	// 7bitシフト
	fmt.Printf("%08b\n", b2>>7)
	// 8bitシフト
	fmt.Printf("%08b\n", b2>>8)
	// 8bitシフト
	fmt.Printf("%08b\n", b2<<8)

	//unsigned
	fmt.Printf("---unsigned---\n")
	var ui8 uint8 = 10
	fmt.Printf("before shifted   = %d (%08b)\n", ui8, byte(ui8))
	fmt.Printf("1bit right shift = %d (%08b)\n", ui8>>1, byte(ui8>>1))
	fmt.Printf("2bit right shift = %d (%08b)\n", ui8>>2, byte(ui8>>2))
	fmt.Printf("1bit left  shift = %d (%08b)\n", ui8<<1, byte(ui8<<1))
	fmt.Printf("2bit left  shift = %d (%08b)\n", ui8<<2, byte(ui8<<2))

	//signed
	fmt.Printf("---signed---\n")
	var i8 int8 = -10
	fmt.Printf("before shifted   = %d (%08b)\n", i8, byte(i8))
	fmt.Printf("1bit right shift = %d (%08b)\n", i8>>1, byte(i8>>1))
	fmt.Printf("2bit right shift = %d (%08b)\n", i8>>2, byte(i8>>2))
	fmt.Printf("1bit left  shift = %d (%08b)\n", i8<<1, byte(i8<<1))
	fmt.Printf("2bit left  shift = %d (%08b)\n", i8<<2, byte(i8<<2))

}
