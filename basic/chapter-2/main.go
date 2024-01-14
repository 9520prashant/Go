package main

import (
	"errors"
	"fmt"
)

func main() {
	Hello("Praash")
	fmt.Println(sum(2, 3))
	var sum, diff, mul, div, err = airthmeticOp(8, -4)
	if err != nil {
		fmt.Println("Some went wrong!!")
	} else {
		fmt.Printf("The value of Sum %v and Diff %v and Mul %v and div %v \n", sum, diff, mul, div)
	}
	var C_Switch_result bool = Conditonal_Switch(4)
	fmt.Println(C_Switch_result)
}

func Hello(arg string) {
	fmt.Println("Hello " + arg)
}

func sum(a int, b int) int {
	return a + b
}

func airthmeticOp(a int, b int) (int, int, int, int, error) {
	var err error
	if b == 0 {
		err = errors.New("Can't divide!")
		return 0, 0, 0, 0, err
	}
	return a + b, a - b, a * b, a / b, err

}

func Conditonal_Switch(a int) bool {
	switch a {
	case 0:
		return true
	case 1:
		return true
	default:
		return false
	}
}
