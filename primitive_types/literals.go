package literals

import "fmt"

func main() {
	integer()
	floating()
	rune()
	string()
}

func integer() {
	bin := 0b10
	octal := 0o10
	octal2 := 010 // Do not use it
	hex := 0x10
	bigNumber := 1_234
	fmt.Println(bin, octal, octal2, hex, bigNumber)
}

func floating() {
	num := 3.30e2
	num2 := 0xa0p1 // p for exponent
	fmt.Println(num, num2)
}

func rune() {
	// below are a in rune
	a := 'a'
	oct := '\141'       // 8-bit octal numbers
	hex := '\x61'       // 8-bit hexadecimal numbers
	hex2 := '\u0061'    // 16-bit hexadecimal numbers
	uni := '\U00000061' // 32-bit unicode

	fmt.Println(a, oct, hex, hex2, uni)
}

func string() {
	str := "Greeting and\n\"Salutations\""
	backquote := `don't need to use / to be "escaped"`
	fmt.Println(str)
	fmt.Println(backquote)
}
