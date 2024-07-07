package main

import (
	"fmt"
)

func main() {
	// Get the user's input for the score
	var score float64
	fmt.Print("Enter the score: ")
	_, err := fmt.Scan(&score)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Check if the score is valid
	if score < 0 || score > 100 {
		fmt.Println("Invalid score")
		return
	}

	// Determine the grade based on the score
	var grade string
	switch {
	case 1 score >= 85:
		grade = "A"
	case 2 score >= 70:
		grade = "B"
	case 3 score >= 55:
		grade = "C"
	case 4 score >= 40:
		grade = "D"
	default:
		grade = "E"
	}

	
}
