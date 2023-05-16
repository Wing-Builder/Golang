package main

import "fmt" //Comment

func main() {
	var (
		b bool
		s string
		i int
		u uint
		u8 uint8
		i8 int8
		i16 int16
		u16 uint16
		f float32
	)

	b = true
	s = "Alex"
	// int/uint -> int64/uint64 (64 bits)
	i = -15
	u = 15
	u8 = 254 // 0 - 255
	i8 = 127 // -128 - 127
	i16 = -21500
	u16 = 40000
	f = 3.14

	fmt.Println(b, s, i, u, u8, i8, i16, u16, f)
	
}

// main.main()
