package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func checkNumber(input string) (int, error) {
	res := map[string]int{
		"zero":      0,
		"one":       1,
		"two":       2,
		"three":     3,
		"four":      4,
		"five":      5,
		"six":       6,
		"seven":     7,
		"eight":     8,
		"nine":      9,
		"ten":       10,
		"eleven":    11,
		"twelve":    12,
		"thirteen":  13,
		"fourteen":  14,
		"fifteen":   15,
		"sixteen":   16,
		"seventeen": 17,
		"eighteen":  18,
		"nineteen":  19,
		"twenty":    20,
		"thirty":    30,
		"forty":     40,
		"fifty":     50,
		"sixty":     60,
		"seventy":   70,
		"eighty":    80,
		"ninety":    90,
	}

	for i := 1; i <= len(input)+1; i++ {
		for k, v := range res {
			if k == input[:i-1] {
				// fmt.Println("aa", k, v, input[:i-1])
				return v, nil
			}

		}

	}
	return -1, errors.New("Invalid digit")
}

func addStrings(line []string, l int) int {
	var concatenated = 0
	var err error

	if l == 1 {
		concatenated, err = strconv.Atoi(line[0] + line[0])
	} else {
		concatenated, err = strconv.Atoi(line[0] + line[l-1])
	}

	if err == nil {
		return concatenated
	} else {
		return 0
	}

}

func main() {
	// open file for reading
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	res := [][]string{}
	sum := 0

	// reading lines and transforming in to list
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// numeric characters of each line
		numericChars := []string{}

		for i, j := range line {
			// fmt.Printf("%c %s\n", j, line[i:])
			number, err := checkNumber(line[i:])

			// fmt.Println("numbeer", number, err)

			if err == nil {
				numericChars = append(numericChars, strconv.Itoa(number))
				// fmt.Println("numericChars: ", numericChars)
			} else if unicode.IsDigit(j) {
				numericChars = append(numericChars, string(j))
			}

		}

		res = append(res, numericChars)

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	fmt.Println(res)

	for _, line := range res {
		sum += addStrings(line, len(line))
	}

	fmt.Println(sum)
}
