package main

import (
	"bufio"
	"log/slog"
	"os"
	"strconv"
	"strings"
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

func One() int {
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

	return sum

}

func Two() int {
	// Parse input file
	lines := ParseInput()

	// Initialize variables
	sum := 0

	// Iterate over each line
	for _, line := range lines {
		first := 0
		last := 0
		var ok bool

		// Find the first number in the string
		for i := 0; i < len(line); i++ {
			if first, ok = Find(line[i:]); ok {
				first = first * 10
				break
			}
		}

		// Find the last number in the string
		for i := len(line) - 1; i >= 0; i-- {
			if last, ok = Find(line[i:]); ok {
				break
			}
		}

		// Add to the sum
		sum += first + last
	}

	return sum
}

func Find(s string) (int, bool) {
	switch {
	case s[0] >= '0' && s[0] <= '9':
		return int(s[0] - '0'), true
	case strings.HasPrefix(s, "one"):
		return 1, true
	case strings.HasPrefix(s, "two"):
		return 2, true
	case strings.HasPrefix(s, "three"):
		return 3, true
	case strings.HasPrefix(s, "four"):
		return 4, true
	case strings.HasPrefix(s, "five"):
		return 5, true
	case strings.HasPrefix(s, "six"):
		return 6, true
	case strings.HasPrefix(s, "seven"):
		return 7, true
	case strings.HasPrefix(s, "eight"):
		return 8, true
	case strings.HasPrefix(s, "nine"):
		return 9, true
	case strings.HasPrefix(s, "zero"):
		return 0, true
	default:
		return 0, false
	}
}

func main() {
	// Run the first part of the challenge
	slog.Info("Part One: " + strconv.Itoa(One()))

	// Run the second part of the challenge
	slog.Info("Part Two: " + strconv.Itoa(Two()))
}
