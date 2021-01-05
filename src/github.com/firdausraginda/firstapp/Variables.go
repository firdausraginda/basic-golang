// reference: https://www.youtube.com/watch?v=YS4e4q9oBaU&ab_channel=freeCodeCamp.org

package main

import (
	"fmt"
	"strconv"
)

// defined variable in package level
var i int = 42

// defined several variables in once step because its related in package level
var (
	actorName    string = "Elisabeth Sladen"
	companion    string = "Sarah Jane Smith"
	doctorNumber int    = 3
	season       int    = 11
)

var (
	counter int = 0
)

func main() {
	// 3 ways to defined a variable -------------------------------------------------------

	// 1st way
	var i int
	i = 60

	// 2nd way
	var j int = 27

	// 3rd way
	k := 99

	// 2 ways to print a result -------------------------------------------------------

	fmt.Printf("%v, %T", actorName, actorName) // print with formating string

	fmt.Println(i, j, k)

	// conversion from int to float -------------------------------------------------------
	var a int = 88
	fmt.Printf("%v, %T\n", a, a)

	var b float32
	b = float32(a)
	fmt.Printf("%v, %T\n", b, b)

	// conversion from int to string -------------------------------------------------------
	var c int = 77
	fmt.Printf("%v, %T\n", c, c)

	var d string
	d = strconv.Itoa(c) //Itoa is function inside strconv package to convert variable from integer to string
	fmt.Printf("%v, %T\n", d, d)
}
