package main

import (
	"fmt"
	"strings"
)

// Function to check if two words are anagrams
func areAnagrams(word1, word2 string) bool {
	// Remove spaces and convert to lowercase for comparison
	word1 = strings.ReplaceAll(strings.ToLower(word1), " ", "")
	word2 = strings.ReplaceAll(strings.ToLower(word2), " ", "")

	// Create maps to store letter frequencies
	freq1 := make(map[rune]int)
	// Populate the maps with letter frequencies
	for _, char := range word1 {
		freq1[char]++
	}
	for _, char := range word2 {
		freq2[char]++
	}

	// Compare the maps
	return compareMaps(freq1, freq2)
}

// Function to compare two maps
func compareMaps(map1, map2 map[rune]int) bool {
	if len(map1) != len(map2) {
		return false

}
