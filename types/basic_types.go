package main

import (
	"fmt"
	"math"
	"strconv"
	"unicode/utf8"
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

	// ======================================================
	// 2. FLOATS - precision, NaN, Inf, math package
	// ======================================================

	fmt.Println("\n 2. Float Types")

	var f32 float32 = 3.1415926
	var f64 float64 = 3.14159265358979323846

	fmt.Printf("float32: %.7f\n", f32)  // less precision
	fmt.Printf("float64: %.15f\n", f64) // more precision

	// Use math package for constants and functions
	fmt.Printf("Pi: %f, E: %f\n", math.Pi, math.E)
	fmt.Printf("Sin(Pi/2): %f\n", math.Sin(math.Pi/2))

	// Special values
	var nan = math.NaN()
	var inf = math.Inf(1)
	var ninf = math.Inf(-1)

	fmt.Printf("NaN: %v, Inf: %v, -Inf: %v\n", nan, inf, ninf)

	// NaN != NaN - always false!
	fmt.Printf("nan == nan? %t\n", nan == nan) // false!
	fmt.Printf("IsNaN(nan)? %t\n", math.IsNaN(nan))

	// ======================================================
	// 3. STRINGS - UTF-8, immutability, indexing, length
	// ======================================================

	fmt.Printf("\n 3.Strings and UTF-8")

	// string with Unicode: emoji, Chinese, etc.
	s := "Hello, 世界 🐹"

	fmt.Printf("String: %s\n", s)
	fmt.Printf("len(s) [bytes]: %d\n", len(s))
	fmt.Printf("Rune count: %d\n", utf8.RuneCountInString(s))

	// Iterate by bytes - DANGEROUS for Unicode
	fmt.Print("Byte iteration: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i]) // will print garbage for multi-byte runes
	}
	fmt.Println()

	// Iterate by runes - CORRECT way
	fmt.Printf("Rune iteration: ")
	for _, r := range s {
		fmt.Printf("%c ", r) // prints each character correctly
	}
	fmt.Println()

	// Accessing by index gives you byte, not rune
	fmt.Printf("s[0] (byte): %c (0x%x)\n", s[0], s[0]) // 'H'
	// s[7] is first byte of '世' - meaningless alone
	fmt.Printf("s[7] (byte in the middle of rune): 0x%x\n", s[7])

	// ======================================================
	// 4. RUNE & BYTE - what they are, how to use them
	// ======================================================

	fmt.Println("\n 4. Rune and Byte")

	// rune = int32 (Unicode code point)
	var heart rune = '💖' // single quotes for rune literals
	fmt.Printf("Heart rune: %c (U+%04X)\n", heart, heart)

	// byte = uint8
	var b byte = 'A'
	fmt.Printf("Byte 'A': %c (0x%x)\n", b, b)

	// Convert string to []rune → modify → back to string
	runes := []rune(s)
	runes[0] = 'J' // Change 'H' to 'J'
	modified := string(runes)
	fmt.Printf("Modified string: %s\n", modified)
}
