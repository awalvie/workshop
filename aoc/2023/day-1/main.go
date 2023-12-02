package main

import (
	"bufio"
	"log/slog"
	"os"
	"strconv"
)

// ParseInput parses the input file and returns a slice of strings
func ParseInput() []string {
	slog.Info("Parsing input file")

	// Open the input file
	file, err := os.Open("input.txt")
	if err != nil {
		slog.Error("Error opening input file: %v", err)
	}
	defer file.Close()

	// Read file and append each line to a slice
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Check for errors
	if err := scanner.Err(); err != nil {
		slog.Error("Error scanning input file: %v", err)
	}

	// Return the slice of strings
	return lines
}

func main() {
	// Parse input file
	lines := ParseInput()

	// Initialize variables
	sum := 0

	// Loop through each line
	for _, line := range lines {
		integers := make([]int, 0)
		// Loop through each character in the line
		for _, char := range line {
			// Check if the character is a number
			if char >= '0' && char <= '9' {
				// Append the number to the slice
				integers = append(integers, int(char-'0'))
			}
		}

		// Get the first and last index of the slice
		first := integers[0]
		last := integers[len(integers)-1]

		number := first*10 + last

		// Add to the sum
		sum += number
	}

	slog.Info("Sum: " + strconv.Itoa(sum))
}
