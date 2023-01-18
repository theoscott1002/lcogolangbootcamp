package main

import (
	"fmt"
)

func main() {
	var userName string
	var Num1 float64
	var Num2 float64
	var operation string
	var total float64
	fmt.Println("Please Enter your name: ")
	fmt.Scanln(&userName)

	fmt.Printf("WELCOME %v!!! \n", userName)

	fmt.Println("Please Enter first number: ")
	fmt.Scanln(&Num1)

	fmt.Println("Please Enter second number: ")
	fmt.Scanln(&Num2)

	fmt.Println("Please Enter Maths Operation: ")
	fmt.Scanln(&operation)

	if operation == "addition" {
		total = Num1 + Num2
	} else if operation == "subtraction" {
		total = Num1 - Num2
	} else if operation == "division" {
		total = Num1 / Num2
	} else if operation == "multiplication" {
		total = Num1 * Num2
	} else {
		fmt.Println("INVALID OPERATION!!!")
	}

	fmt.Printf("The Answer is: %v\n", total)
}
