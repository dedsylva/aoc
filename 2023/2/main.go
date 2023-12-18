package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const maxRedCubes = 12
const maxGreenCubes = 13
const maxBlueCubes = 14

var DEBUG, PARTONE, PARTTWO bool

func checkGameIsValid(set map[string]int) bool {
	if DEBUG {
		fmt.Println("Sets", set["red"], set["green"], set["blue"], set["red"] > maxRedCubes || set["green"] > maxGreenCubes || set["blue"] > maxBlueCubes)
	}

	if set["red"] > maxRedCubes || set["green"] > maxGreenCubes || set["blue"] > maxBlueCubes {
		return false
	}
	return true
}

func getPower(set map[string]int) int {
	if DEBUG {
		fmt.Println("Current Set", set)
	}

	return set["red"] * set["green"] * set["blue"]
}

func main() {

	_, isPresent := os.LookupEnv("DEBUG")
	if isPresent {
		DEBUG = true
	}

	_, isPresent = os.LookupEnv("PARTONE")
	if isPresent {
		PARTONE = true
	}

	_, isPresent = os.LookupEnv("PARTTWO")
	if isPresent {
		PARTTWO = true
	}

	// open file for reading
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var res int

	// reading lines and transforming in to list
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")

		gameNumber, err := strconv.Atoi(strings.TrimSpace(strings.Split(line[0], "Game")[1]))

		var gameIsValid bool

		if err != nil {
			fmt.Println("Error getting game number:", err)
			return
		}

		games := strings.Split(line[1], ";")

		if DEBUG {
			fmt.Printf("Game: %d\n", gameNumber)
		}

		set := map[string]int{
			"red":   0,
			"blue":  0,
			"green": 0,
		}

		for _, game := range games {

			gameIsValid = true

			records := strings.Split(game, ",")

			for _, record := range records {
				color := strings.Fields(record)[1]
				colorValue, err := strconv.Atoi(strings.Fields(record)[0])

				if err != nil {
					fmt.Println("Error getting color:", err)
					return
				}

				if PARTTWO {
					if set[color] < colorValue {
						if DEBUG {
							fmt.Println("Replacing", color, set[color], colorValue)
						}
						set[color] = colorValue
					}
				}
				if PARTONE {
					set[color] = colorValue
				}
			}

			if DEBUG {
				fmt.Println(set)
			}

			// Part 1
			if PARTONE {
				if !checkGameIsValid(set) {
					if DEBUG {
						fmt.Printf("Game %d is invalid\n", gameNumber)
					}

					gameIsValid = false
					break
				}
			}

		}

		// Part 1
		if PARTONE {
			if gameIsValid {
				res = res + gameNumber
				if DEBUG {
					fmt.Printf("Game %d is valid\n", gameNumber)
				}
			}
		}

		if PARTTWO {
			res = res + getPower(set)
		}

	}
	fmt.Println(res)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

}
