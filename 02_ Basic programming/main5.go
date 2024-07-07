package main

import (
	"fmt"
)

func main() {
	// Input values (you can modify these)
	budget := 40
	duration := 25
	difficulty := 5

	// Calculate scores for each factor
	budgetScore := calculateScore(budget, []int{0, 50, 80, 100}, []int{4, 3, 2, 1})
	durationScore := calculateScore(duration, []int{0, 20, 30, 50}, []int{10, 5, 2, 1})
	difficultyScore := calculateScore(difficulty, []int{0, 3, 6, 10}, []int{10, 5, 1, 0})

	// Calculate total score
	totalScore := budgetScore + durationScore + difficultyScor