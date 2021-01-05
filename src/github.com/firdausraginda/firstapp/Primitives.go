package main

import (
	"fmt"
)

func main() {
	// boolean type -------------------------------------------------------
	var l bool = true
	fmt.Printf("%v, %T\n", l, l)

	n := 1 == 1
	m := 1 == 2
	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", m, m)

	// bit operator -------------------------------------------------------
	a := 10            // 1010
	b := 3             // 0011
	fmt.Println(a & b) // 0010
	fmt.Println(a | b) // 1011
	fmt.Println(a ^ b) // exclusive or 1001

	// bit shifting -------------------------------------------------------
	c := 8              // 2^3
	fmt.Println(c << 3) // 2^3 * 2^3 = 2^6
	fmt.Println(c >> 3) // 2^3 / 2^3 = 2^0

	// float type -------------------------------------------------------
	d := 10.2
	e := 3.7
	fmt.Println(d + e)
	fmt.Println(d - e)
	fmt.Println(d * e)
	fmt.Println(d / e)

	// complex type -------------------------------------------------------
	f := 1 + 2i
	g := 2 + 5.2i
	fmt.Println(f + g)
	fmt.Println(f - g)
	fmt.Println(f * g)
	fmt.Println(f / g)

	// string slicing
	exString := "this is golang"
	fmt.Println(string(exString[1:6])) // need to convert exString to string first before do slicing

	// string concatination -------------------------------------------------------
	string1 := "this is a string"
	string2 := "this is also a string"
	fmt.Printf("%v, %T\n", string1+string2, string1+string2)

	// char definition -------------------------------------------------------
	char := 'a'
	fmt.Printf("%v, %T\n", string(char), char)
}
