package main

import "fmt"

func FizzBuzz(n int) {
	for i := 1; i < n; i++ {
		printFizzBuzzValue(i)
		fmt.Print(", ")
	}
	printFizzBuzzValue(n)
	fmt.Println()
}

func printFizzBuzzValue(n int) {
	if n%3 == 0 && n%5 == 0 {
		fmt.Print("Fizz Buzz")
	} else if n%3 == 0 {
		fmt.Print("Fizz")
	} else if n%5 == 0 {
		fmt.Print("Buzz")
	} else {
		fmt.Print(n)
	}
}
