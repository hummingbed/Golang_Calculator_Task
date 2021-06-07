package main

import "fmt"

func main() {

	//Define  variables
	var num1 float64
	var num2 float64

	var assignmentOperator string

	// output and input values
	fmt.Println("Enter first value :")
	fmt.Scanln(&num1)

	fmt.Println("Enter assignment Operator")
	fmt.Scanln(&assignmentOperator)

	fmt.Println("Enter Second value :")
	fmt.Scanln(&num2)

	//conditional statements

	switch assignmentOperator {
	case "+":
		fmt.Printf("%g %s %g = %g", num1, assignmentOperator, num2, num1+num2)

	case "-":
		fmt.Printf("%g %s %g = %g", num1, assignmentOperator, num2, num1-num2)

	case "*":
		fmt.Printf("%g %s %g = %g", num1, assignmentOperator, num2, num1*num2)

	case "/":
		if num2 == 0.0 {
			fmt.Println("Divide by Zero")
		} else {
			fmt.Printf("%g %s %g = %g", num1, assignmentOperator, num2, num1/num2)
		}
	default:
		fmt.Println("Invalid input value, please try again")
	}

}
