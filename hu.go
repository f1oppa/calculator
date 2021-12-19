package main

import (
	"fmt";
	"time"
)

func main() {
	var num1, num2, operation int

	fmt.Println("1. Összeadás")
	fmt.Println("2. Kivonás")
	fmt.Println("3. Szorzás")
	fmt.Println("4. Osztás")
	fmt.Print("Válassz műveletet: ")
	fmt.Scan(&operation)

	fmt.Print("Első szám: ")
	fmt.Scan(&num1)
	fmt.Print("Második szám: ")
	fmt.Scan(&num2)

	if operation == 1{
		fmt.Println("Eredmény:", num1 + num2)
	}
	if operation == 2{
		fmt.Println("Eredmény:", num1 - num2)
	}
	if operation == 3{
		fmt.Println("Eredmény:", num1 * num2)
	}
	if operation == 4{
		fmt.Println("Eredmény:", num1 / num2)
	}
	fmt.Println("A kilépéshez nyomd le a CTRL-C billentyűket.")

	for {
		time.Sleep(5*time.Second)
	}
}
