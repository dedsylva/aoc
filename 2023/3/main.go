package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

// struct to hold position of the special characters
type positionCharStruct struct {
	i int
	j int
}

var DEBUG, PARTONE, PARTTWO bool

func printGrid(grid [][]string) {
	fmt.Printf("\n")
	for _, g := range grid {
		fmt.Println(g)
	}
	fmt.Printf("\n")
}

func lookAroundPosition(grid [][]string, p positionCharStruct) map[int]bool {
	ret := make(map[int]bool)

	relativePositions := []positionCharStruct{
		{p.i - 1, p.j},     // Top
		{p.i + 1, p.j},     // Bottom
		{p.i, p.j - 1},     // Left
		{p.i, p.j + 1},     // Right
		{p.i - 1, p.j - 1}, // Diagonal top-left
		{p.i - 1, p.j + 1}, // Diagonal top-right
		{p.i + 1, p.j - 1}, // Diagonal bottom-left
		{p.i + 1, p.j + 1}, // Diagonal bottom-right
	}

	for _, pos := range relativePositions {
		// guarantees that stays within the grid
		if pos.i < len(grid) && pos.j < len(grid[pos.i]) {
			value, err := strconv.Atoi(grid[pos.i][pos.j])

			if err == nil {
				ret[value] = true
			}
		}
	}

	return ret
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
	var checkingNumbers bool
	grid := [][]string{}

	var positions []positionCharStruct
	partNumbers := make(map[int]bool)
	lineNumber := 0

	// reading lines and transforming in to list
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		row := []string{}
		offset := 0

		if DEBUG {
			fmt.Println(line)
		}

		digits := ""

		for charPosition, character := range line {
			// fmt.Printf("@ %d %c\n", i, character)

			if unicode.IsDigit(character) {
				checkingNumbers = true
				digits += string(character)

				// since we need to group chars in one position, charPosition won't be the same as the column position on the grid
			} else {
				checkingNumbers = false
			}

			if !checkingNumbers {
				// adding digits
				if digits != "" {
					offset += len(digits) - 1
					row = append(row, digits)
					digits = ""
				}
				// adding character (that's deep)
				row = append(row, string(character))

				if DEBUG {
					fmt.Printf("%d %d %c\n", lineNumber, charPosition-offset, character)
				}

				if string(character) != "." {
					position := positionCharStruct{i: lineNumber, j: charPosition - offset}
					positions = append(positions, position)
				}

			}

		}
		grid = append(grid, row)
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	printGrid(grid)

	for _, p := range positions {
		if DEBUG {
			fmt.Printf("%+v %s\n", p, grid[p.i][p.j])
		}

		partNumbersRes := lookAroundPosition(grid, p)

		// I did not like this, there must be a better way to do it
		for number, isPart := range partNumbersRes {
			partNumbers[number] = isPart
		}
		fmt.Printf("%+v %s\n", partNumbersRes, grid[p.i][p.j])
	}

	for key, _ := range partNumbers {
		fmt.Println(key)
	}

	fmt.Println("res", res)

}
