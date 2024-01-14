package main

import (
	"fmt"
	// "unicode/utf8"
)

func main() {
	// Hello World in Go 🥳

	// fmt.Println("Hello, World!")

	// Variables in Go 🥳

	// var intVar int16 = 32767
	// intVar = intVar + 1
	// fmt.Println("intVar", intVar)

	// var floatVar float32 = 314344
	// fmt.Println("floatVar", floatVar)

	// var stringVar string = "Hello \n World"
	// fmt.Println("stringVar\n", stringVar)

	// var stringVar string = `hello world`
	// fmt.Println("stringVar\n", stringVar)

	stringVar := "Short Way of Creating String"
	fmt.Println("StringVar", stringVar)

	//fmt.Println(len("γ")) // will return/give the number of Byte in String which is  2
	// *******NOTE********
	// A string contains characters as unicode points, not bytes. len(string) returns the number of bytes in this string, but not the number of characters. Therefore, we are converting the string to rune array and then finding the array length.

	//fmt.Println(len([]rune("γ")))            // will return/give the number of character in String which is 1
	//fmt.Println(utf8.RuneCountInString("γ")) // will return/give the number of character in String which is 1

	//var boolVar bool = true
	//fmt.Println(!boolVar)

	// TypeCasting in Go 🥳

	// var result int = int(intVar) + int(floatVar)
	// fmt.Println("result", result)

}
