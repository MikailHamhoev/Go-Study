package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("=== GO BASIC TYPES - DEEP DIVE ===")

	// ======================================================
	// 1. INTEGERS - signed, unsigned, sizes, aliases
	// ======================================================

	fmt.Println(" 1. Integer Types")

	var i8 int8 = 127 // -128 to 127
	var i16 int16 = 32767
	var i32 int32 = 2147483647
	var i64 int64 = 9223372036854775807

	var u8 uint8 = 255 // 0 to 255 (same as 'byte')
	var u16 uint16 = 65535
	var u32 uint32 = 4294967295
	var u64 uint64 = 18446744073709551615

	fmt.Printf("int8: %d, int16:%d, int32: %d, int64: %d\n", i8, i16, i32, i64)
	fmt.Printf("uint8: %d, uint32: %d, uint32: %d, uint64: %d\n", u8, u16, u32, u64)

	// 'int' and 'uint' are platform-dependent (usually 64-bit on modern systems)
	var platformInt int = -42
	var platformUint uint = 42
	fmt.Printf("platform int: %d (size: %d bits)\n", platformInt, strconv.IntSize)
	fmt.Printf("platform uint: %d\n", platformUint)

	// uintptr - for pointer arithmetic (rarely needed)
	var addr uintptr = 0x12345678
	fmt.Printf("uintptr (hex): 0x%x\n", addr)
}
