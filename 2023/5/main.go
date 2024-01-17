package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type data struct {
	sourceName       string
	sourceStart      int
	destinationName  string
	destinationStart int
	rangeStart       int
}

func toNumber(slice []string) []int {
	res := []int{}

	for _, s := range slice {
		value, err := strconv.Atoi(string(s))

		if err == nil {
			res = append(res, value)
		}
	}

	return res
}

func lineToData(line string, fullName string) data {
	var res data
	lineAux := toNumber(strings.Split(strings.TrimSpace(line), " "))

	if DEBUG {
		fmt.Println("lineAux", lineAux)
	}

	res.destinationStart = lineAux[0]
	res.destinationName = strings.Split(fullName, "-")[2]
	res.sourceStart = lineAux[1]
	res.sourceName = strings.Split(fullName, "-")[0]
	res.rangeStart = lineAux[2]

	return res

}

func printAlmanac(almanac map[string][]data) {

	for k := range almanac {
		for _, d := range almanac[k] {
			fmt.Printf("name: %s\n data: %+v\n", k, d)
		}
		fmt.Printf("\n\n")
	}
}

func findRange(almanac []data, sourceValue int) int {
	for _, almanacData := range almanac {
		sourceOffset := 0
		for i := almanacData.sourceStart; i < almanacData.sourceStart+almanacData.rangeStart; i++ {
			if sourceValue == i {
				sourceOffset = i - almanacData.sourceStart
				return almanacData.destinationStart + sourceOffset
			}
		}

	}
	return sourceValue
}

var DEBUG bool

func main() {

	_, isPresent := os.LookupEnv("DEBUG")
	if isPresent {
		DEBUG = true
	}

	var seeds []int

	// open file for reading
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	mapName := ""

	almanac := map[string][]data{
		"seed-to-soil":            {},
		"soil-to-fertilizer":      {},
		"fertilizer-to-water":     {},
		"water-to-light":          {},
		"light-to-temperature":    {},
		"temperature-to-humidity": {},
		"humidity-to-location":    {},
	}

	// reading lines and transforming in to list
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// non-empty line
		if len(scanner.Text()) > 0 {
			lines := strings.Split(scanner.Text(), ":")

			if len(lines) > 1 {
				// first line
				if lines[0] == "seeds" {
					seeds = toNumber(strings.Split(strings.TrimSpace(lines[1]), " "))

					if DEBUG {
						fmt.Println("seeds", seeds)
					}
				} else {
					// map definition
					mapName = strings.Split(lines[0], " ")[0]

					if DEBUG {
						fmt.Println(mapName)
					}
				}

			} else {
				// getting ranges
				almanac[mapName] = append(almanac[mapName], lineToData(lines[0], mapName))

			}
		}

	}

	// printAlmanac(almanac)

	almanacToCategory := map[int]string{
		0: "seed-to-soil",
		1: "soil-to-fertilizer",
		2: "fertilizer-to-water",
		3: "water-to-light",
		4: "light-to-temperature",
		5: "temperature-to-humidity",
		6: "humidity-to-location",
	}

	res := map[string][]int{
		"seed":        {},
		"soil":        {},
		"fertilizer":  {},
		"water":       {},
		"light":       {},
		"temperature": {},
		"humidity":    {},
		"location":    {},
	}

	var sourceValue int
	for _, seed := range seeds {
		res["seed"] = append(res["seed"], seed)

		sourceValue = seed
		for i := 0; i < len(almanacToCategory); i++ {
			destinationValue := findRange(almanac[almanacToCategory[i]], sourceValue)
			// it's ugly I know
			res[almanac[almanacToCategory[i]][0].destinationName] = append(res[almanac[almanacToCategory[i]][0].destinationName], destinationValue)
			sourceValue = destinationValue
		}
	}

	lowestLocation := res["location"][0]
	for i := 0; i < len(res["seed"]); i++ {
		if res["location"][i] < lowestLocation {
			lowestLocation = res["location"][i]
		}
		fmt.Printf("seed: %d, soil: %d, fertilizer: %d, water: %d, light: %d, temperature: %d, humidity: %d, location: %d\n", res["seed"][i], res["soil"][i], res["fertilizer"][i], res["water"][i], res["light"][i], res["temperature"][i], res["humidity"][i], res["location"][i])
	}
	fmt.Println(lowestLocation)
}
