package main

import (
	"fmt"
)

func main() {
	// int data type --------------------------------------------------------
	var num int8 = 10
	fmt.Println(num)

	// string data type --------------------------------------------------------
	name := "Raginda Firdaus"
	fmt.Println(name)

	// declare multiple variable at once --------------------------------------------------------
	var (
		firstName = "Eko"
		lastName  = "Khannedy"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)

	// const data type --------------------------------------------------------
	const firstNameConst string = "Raginda"
	const lastNameConst = "Firdaus"
	// firstNameConst = "Cannot be modified"
	// lastNameConst = "Cannot be modified"

	// data type conversion --------------------------------------------------------
	var num32 int32 = 100000
	var num64 int64 = int64(num32)
	var num8 int8 = int8(num32)
	fmt.Println(num32)
	fmt.Println(num64)
	fmt.Println(num8) // overflow

	// string slicing --------------------------------------------------------
	var name2 = "Agi"
	var e = name2[0] // this slicing automatically change data type into byte
	eString := string(e)
	fmt.Println(name2)
	fmt.Println(eString)

	// type declarations --------------------------------------------------------
	type isMarried bool
	var marriedStatus isMarried = true
	fmt.Println(marriedStatus)

	// math operations --------------------------------------------------------
	var a = 10
	var b = 7
	var c = a + b
	fmt.Println(c)

	// augmented assignments
	a += 2
	fmt.Println(a)
	a -= 2
	fmt.Println(a)
	a *= 2
	fmt.Println(a)
	a /= 2
	fmt.Println(a)
	a %= 2
	fmt.Println(a)

	// unary operator
	a++
	fmt.Println(a)
	a--
	fmt.Println(a)

	// comparison operations --------------------------------------------------------
	var name1 = "agi"
	name2 = "baba"
	var compareResult bool = name1 < name2
	fmt.Println("compare result 1: ", compareResult)

	num1 := 100
	num2 := 100
	compareResult = num1 != num2
	fmt.Println("compare result 2: ", compareResult)

	// boolean operations --------------------------------------------------------
	finalScore := 90
	absenceScore := 80

	successFinalScore := finalScore > 80
	successAbsence := absenceScore > 80

	successScore := successFinalScore && successAbsence

	fmt.Println("success score 1: ", successScore)
	fmt.Println("success score 2: ", finalScore > 80 && absenceScore > 80)

	// array data type --------------------------------------------------------

	// manual array declaration
	var arrOfNames [3]string
	arrOfNames[0] = "kle"
	arrOfNames[1] = "syah"
	arrOfNames[2] = "putra"

	fmt.Println(arrOfNames)
	fmt.Println(arrOfNames[0])
	fmt.Println(arrOfNames[1])
	fmt.Println(arrOfNames[2])

	// direct array declaration
	var arrOfNums = [4]int{
		90,
		80,
		95,
	}
	fmt.Println(arrOfNums)

	// print length of array
	fmt.Println(len(arrOfNums))

	// declare array without specify the length at the beginning
	var newArr = [...]int{
		11,
		22,
		33,
		44,
		55,
		66,
		77,
	}
	fmt.Println(newArr)
	fmt.Println(len(newArr))

	// modify element on array
	newArr[0] = 100
	fmt.Println(newArr)

	// slice data type --------------------------------------------------------

}
