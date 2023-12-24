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
	i                   int
	j                   int
	partNumbersQuantity int
	partNumbers         []int
}

// struct to hold resultNumbers and their positions (since the same number can be multiple places)
type resultNumbersStruct struct {
	number   int
	line     int
	posBegin int
	posEnd   int
}

var DEBUG, PARTONE, PARTTWO bool

func printGrid(grid [][]rune) {
	fmt.Printf("\n")
	for _, line := range grid {
		for _, l := range line {
			fmt.Printf("%c", l)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func looksForNumber(grid [][]rune, p positionCharStruct) resultNumbersStruct {
	// loop to find beginning of number
	// fmt.Printf("##### %+v %c\n", p, grid[p.i][p.j])
	indexBegin := 0
	indexEnd := 0
	number := ""

	for j := p.j; j >= 0; j-- {
		// fmt.Printf("@@@ %d %c\n", j, grid[p.i][j])
		if !unicode.IsDigit(grid[p.i][j]) {
			indexBegin = j + 1
			break
		}
	}

	// fmt.Println("index:", index)
	// loop to find end of number
	for j := indexBegin; j <= len(grid[p.i])-1; j++ {
		// fmt.Printf(">>> %d %c\n", j, grid[p.i][j])
		if !unicode.IsDigit(grid[p.i][j]) {
			indexEnd = j - 1
			break
		}
		number += string(grid[p.i][j])
	}

	// fmt.Println("found Number:", number)
	value, err := strconv.Atoi(number)

	if err != nil {
		fmt.Printf("Error converting %s to integer:", number)
		return resultNumbersStruct{number: -1, line: -1, posBegin: -1, posEnd: -1}
	}

	return resultNumbersStruct{number: value, line: p.i, posBegin: indexBegin, posEnd: indexEnd}
}

func lookAroundPosition(grid [][]rune, p positionCharStruct) []resultNumbersStruct {
	// fmt.Printf("@@@@@ %+v %c\n", p, grid[p.i][p.j])
	ret := []resultNumbersStruct{}

	relativePositions := []positionCharStruct{
		{p.i - 1, p.j, 0, []int{}},     // Top
		{p.i + 1, p.j, 0, []int{}},     // Bottom
		{p.i, p.j - 1, 0, []int{}},     // Left
		{p.i, p.j + 1, 0, []int{}},     // Right
		{p.i - 1, p.j - 1, 0, []int{}}, // Diagonal top-left
		{p.i - 1, p.j + 1, 0, []int{}}, // Diagonal top-right
		{p.i + 1, p.j - 1, 0, []int{}}, // Diagonal bottom-left
		{p.i + 1, p.j + 1, 0, []int{}}, // Diagonal bottom-right
	}

	for _, pos := range relativePositions {
		// guarantees that stays within the grid
		if pos.i >= 0 && pos.j >= 0 {
			// fmt.Printf("i: %d pos:%+v -> %c\n", i, pos, grid[pos.i][pos.j])
			if pos.i < len(grid) && pos.j < len(grid[pos.i]) {
				if unicode.IsDigit(grid[pos.i][pos.j]) {
					resNumber := looksForNumber(grid, pos)
					ret = append(ret, resNumber)
					// fmt.Printf("number:%d i:%d character:%c\n", number, i, grid[p.i][p.j])
				}

			}
		}
	}

	return ret
}

func contains(numberList []resultNumbersStruct, newNumber resultNumbersStruct) bool {
	for _, list := range numberList {
		if list.number == newNumber.number && list.line == newNumber.line && list.posBegin == newNumber.posBegin && list.posEnd == newNumber.posEnd {
			return true
		}
	}
	return false
}

func main() {

	_, isPresent := os.LookupEnv("DEBUG")
	if isPresent {
		DEBUG = true
	}

	// open file for reading
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var res int
	grid := [][]rune{}

	var positions []positionCharStruct
	var partNumbers []resultNumbersStruct
	// partNumbers := []int{}
	lineNumber := 0

	// reading lines and transforming in to list
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		row := []rune{}
		// offset := 0

		if DEBUG {
			fmt.Println(line)
		}

		// digits := ""
		for charPosition, character := range line {
			// adding character (that's deep)
			row = append(row, character)

			if string(character) != "." && !unicode.IsDigit(character) {
				position := positionCharStruct{i: lineNumber, j: charPosition, partNumbersQuantity: 0}
				positions = append(positions, position)
			}

		}

		grid = append(grid, row)
		lineNumber++

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	for i, p := range positions {
		// fmt.Printf("**** %+v %c\n", p, grid[p.i][p.j])

		partNumbersRes := lookAroundPosition(grid, p)
		for _, number := range partNumbersRes {
			if !contains(partNumbers, number) {
				partNumbers = append(partNumbers, number)

				positions[i].partNumbers = append(positions[i].partNumbers, number.number)
				positions[i].partNumbersQuantity += 1

			}
		}

	}

	if DEBUG {
		for _, p := range positions {
			fmt.Printf("%c (%d, %d) has %+v partNumbers\n", grid[p.i][p.j], p.i, p.j, p.partNumbers)
		}
	}

	for _, position := range positions {
		if len(position.partNumbers) == 2 {
			res += position.partNumbers[0] * position.partNumbers[1]
		}
	}

	fmt.Println(res)

}
