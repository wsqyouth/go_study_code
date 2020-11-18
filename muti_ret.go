package main

import (
	"fmt"
	"unicode"
)

func main() {

	const mixed = "abc234de"
	for _, c := range mixed {
		fmt.Printf("For %q: ", c)
		if unicode.IsDigit(c) {
			fmt.Println("is digit rune")
		}

	}

	var x int
	for i := 0; i < len(mixed); {
		x, i = nextInt([]byte(mixed), i)
		fmt.Println(i)
		fmt.Println(x)
	}
}

func nextInt(b []byte, i int) (int, int) {
	for ; i < len(b) && !unicode.IsDigit(rune(b[i])); i++ {

	}
	x := 0
	for ; i < len(b) && unicode.IsDigit(rune(b[i])); i++ {
		x = x*10 + int(b[i]) - '0'
	}
	return x, i
}

//测试rune的用法
