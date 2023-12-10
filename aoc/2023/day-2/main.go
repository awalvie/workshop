package main

import (
	"bufio"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

type Cube struct {
	Count int
	Color string
}

type Result struct {
	Cubes []Cube
}

type Set struct {
	Result []Result
}

type Game struct {
	ID   int
	Sets []Set
}

// ParseInput parses the input file and returns a slice of strings
func ParseInput() ([]Game, error) {
	var games []Game

	// Open input.txt for reading
	f, err := os.Open("input.txt")
	if err != nil {
		slog.Error("Error opening input.txt: " + err.Error())
	}
	defer f.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		// Ex: Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
		var game Game

		// Split the line on the colon
		sLine := strings.Split(line, ":")

		// Get the game ID
		ID, err := strconv.Atoi((strings.Split(sLine[0], " ")[1]))
		if err != nil {
			slog.Error("Error converting string to int: " + err.Error())
		}

		// Set the game ID
		game.ID = ID

		// Split sLine on the semicolon
		sSets := strings.Split(sLine[1], ";")

		// Loop through the sets
		for _, sSet := range sSets {
			var set Set

			// Trim the whitespace
			sSet = strings.TrimSpace(sSet)

			// Split the set on the comma
			sResults := strings.Split(sSet, ",")

			// Loop through the results
			for _, sResult := range sResults {
				var result Result

				// Trim the whitespace
				sResult = strings.TrimSpace(sResult)

				// Split the result on the space
				sCubes := strings.Split(sResult, " ")

				// Get the count
				count, err := strconv.Atoi(sCubes[0])
				if err != nil {
					slog.Error("Error converting string to int: " + err.Error())
				}

				// Get the color
				color := sCubes[1]

				// Set the result
				result.Cubes = append(result.Cubes, Cube{Count: count, Color: color})

				// Append the result to the set
				set.Result = append(set.Result, result)
			}

			// Append the set to the game
			game.Sets = append(game.Sets, set)
		}

		// Append the game to the games
		games = append(games, game)
	}

	return games, nil
}

// One runs the first part of the challenge
func One() {
	games, err := ParseInput()
	if err != nil {
		slog.Error("Error parsing input: " + err.Error())
	}

	invalidSum := 0
	gameSum := 0

	// Loop through the games
	for _, game := range games {
		gameSum += game.ID
		found := false
		// Loop through the sets
		for _, set := range game.Sets {
			if !found {
				// Loop through the results
				for _, result := range set.Result {
					// Loop through the cubes
					blueCount := 0
					redCount := 0
					greenCount := 0

					for _, cube := range result.Cubes {
						// Add the count to the color
						if cube.Color == "blue" {
							blueCount += cube.Count
						} else if cube.Color == "red" {
							redCount += cube.Count
						} else if cube.Color == "green" {
							greenCount += cube.Count
						}
					}

					if redCount > 12 || blueCount > 14 || greenCount > 13 {
						invalidSum += game.ID
						found = true
						break
					}
				}
			}
		}
	}

	slog.Info("Solution One: " + strconv.Itoa(gameSum-invalidSum))
}

// Two runs the second part of the challenge
func Two() {
	games, err := ParseInput()
	if err != nil {
		slog.Error("Error parsing input: " + err.Error())
	}

	totalpower := 0

	// Loop through the games
	for _, game := range games {
		var minBlue, minRed, minGreen int
		// Loop through the sets
		for _, set := range game.Sets {
			// Loop through the results
			for _, result := range set.Result {
				// Looop through the cubes
				for _, cube := range result.Cubes {
					if cube.Color == "blue" {
						if cube.Count > minBlue {
							minBlue = cube.Count
						}
					} else if cube.Color == "red" {
						if cube.Count > minRed {
							minRed = cube.Count
						}
					}
					if cube.Color == "green" {
						if cube.Count > minGreen {
							minGreen = cube.Count
						}
					}
				}
			}
		}
		totalpower += (minBlue * minRed * minGreen)
	}
	slog.Info("Solution Two: " + strconv.Itoa(totalpower))
}

func main() {
	// Run the first part of the challenge
	One()

	// Run the second part of the challenge
	Two()
}
