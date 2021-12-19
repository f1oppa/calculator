package main

import (
	"fmt";
	"time"
)

func main() {
	var num1, num2, operation int

	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")
	fmt.Print("Select an operation: ")
	fmt.Scan(&operation)

	fmt.Print("First number: ")
	fmt.Scan(&num1)
	fmt.Print("Second number: ")
	fmt.Scan(&num2)

	if operation == 1{
		fmt.Println("Result:", num1 + num2)
	}
	if operation == 2{
		fmt.Println("Result:", num1 - num2)
	}
	if operation == 3{
		fmt.Println("Result:", num1 * num2)
	}
	if operation == 4{
		fmt.Println("Result:", num1 / num2)
	}
	fmt.Println("Press CTRL-C to quit.")

	for {
		time.Sleep(5*time.Second)
	}
}