package main

import (
	"fmt"
)

func main() {
	var input int
	fmt.Print("Enter a positive integer: ")
	_, err := fmt.Scan(&input)
	if err != nil || input <= 0 {
		fmt.Println("Invalid input. Please enter a positive integer.")
		return
	}

	