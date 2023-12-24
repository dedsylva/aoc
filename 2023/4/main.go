package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type scratchCard struct {
	cardNumber             int
	winningNumbersQuantity int
	instances              int
}

func parseNumbers(slice []string) []int {
	res := []int{}

	for _, s := range slice {
		value, err := strconv.Atoi(s)

		if err == nil {
			res = append(res, value)
		}
	}

	return res
}

func checkWinningNumbers(cardNumber int, winningNumbers []int, myNumbers []int) int {
	var matches []int

	if DEBUG {
		fmt.Println("cardNumber", cardNumber)
	}

	for _, m := range myNumbers {
		for _, w := range winningNumbers {
			if m == w {
				matches = append(matches, m)
				if DEBUG {
					fmt.Println("found match", w)
				}
			}
		}
	}

	return len(matches)
}

var DEBUG bool

func main() {

	_, isPresent := os.LookupEnv("DEBUG")
	if isPresent {
		DEBUG = true
	}

	scratchCards := []scratchCard{}
	res := 0

	// open file for reading
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// reading lines and transforming in to list
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "|")

		leftSide := strings.Split(strings.Split(line[0], ":")[1], " ")
		rightSide := strings.Split(line[1], " ")
		beginning := strings.Split(line[0], ":")[0]
		lengthBeginning := len(strings.Split(beginning, " "))

		winningNumbers := parseNumbers(leftSide)
		myNumbers := parseNumbers(rightSide)

		cardNumber, err := strconv.Atoi(strings.Split(beginning, " ")[lengthBeginning-1])

		if DEBUG {
			fmt.Println(line)
			fmt.Println(cardNumber)
			fmt.Println(winningNumbers, len(winningNumbers))
			fmt.Println(myNumbers, len(myNumbers))
		}

		if err != nil {
			fmt.Println("Error getting card number:", err)
			return
		}

		winningNumbersQuantity := checkWinningNumbers(cardNumber, winningNumbers, myNumbers)
		scratchCardsLine := scratchCard{cardNumber: cardNumber, winningNumbersQuantity: winningNumbersQuantity, instances: 1}

		scratchCards = append(scratchCards, scratchCardsLine)

	}
	for s, scratchCard := range scratchCards {
		if DEBUG {
			fmt.Printf("before scratchCard: %+v\n", scratchCard)
		}

		// loops through instances, checks how many winningNumbers and adds for next ones
		for i := 1; i <= scratchCard.instances; i++ {
			for w := 1; w <= scratchCard.winningNumbersQuantity; w++ {
				if w < len(scratchCards) {
					scratchCards[s+w].instances += 1
				}
			}
		}

	}
	for _, scratchCard := range scratchCards {
		res += scratchCard.instances
		if DEBUG {
			fmt.Printf("after scratchCard: %+v\n", scratchCard)
		}
	}

	fmt.Println(res)

}
