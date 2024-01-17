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

type values struct {
	sourceValue      int
	destinationValue int
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

func findMatch(valueStruct []values, sourceValue int) int {
	for _, v := range valueStruct {
		if v.sourceValue == sourceValue {
			return v.destinationValue
		}
	}

	return sourceValue
}

// I know this is O(n), but I can't control if the array is filled from 0 up
// I could do a offset value and order each time I fill the array, but is just too much work
// func contains(dataArray data, array []values, number int) bool {
// }

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

	ranges := map[string][]values{
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
				// fmt.Println("lineToData", lineToData(lines[0], mapName))

			}
		}

	}

	if DEBUG {
		printAlmanac(almanac)
	}

	for k := range almanac {
		if DEBUG {
			fmt.Println("### ", k, len(almanac[k]))
		}

		for i := 0; i < len(almanac[k]); i++ {

			if DEBUG {
				fmt.Println("line:", i+1)
			}

			for j := 0; j < almanac[k][i].rangeStart; j++ {
				rangeAux := values{sourceValue: almanac[k][i].sourceStart + j, destinationValue: almanac[k][i].destinationStart + j}
				ranges[k] = append(ranges[k], rangeAux)
				// Assuming that there will never be an overlap between ranges of multiple lines
				// if !contains(almanac[k][i], ranges[k], j) {
				// 	ranges[k] = append(ranges[k], j)
				// }

				if DEBUG {
					fmt.Println(j)
				}

			}

		}

	}

	// for k, value := range ranges {
	// 	// slices.Sort(v)
	// 	for _, v := range value {
	// 		fmt.Printf("%s: (%d, %d)\n", k, v.sourceValue, v.destinationValue)
	// 	}
	// }

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

	rangeToRes := map[string]string{
		"seed-to-soil":            "soil",
		"soil-to-fertilizer":      "fertilizer",
		"fertilizer-to-water":     "water",
		"water-to-light":          "light",
		"light-to-temperature":    "temperature",
		"temperature-to-humidity": "humidity",
		"humidity-to-location":    "location",
	}

	keyOrders := [7]string{"seed-to-soil",
		"soil-to-fertilizer",
		"fertilizer-to-water",
		"water-to-light",
		"light-to-temperature",
		"temperature-to-humidity",
		"humidity-to-location"}

	resAux := 0

	for _, seed := range seeds {
		res["seed"] = append(res["seed"], seed)

		resAux = seed
		for k, _ := range keyOrders {
			category := ranges[keyOrders[k]]
			resAux = findMatch(category, resAux)
			// fmt.Println(keyOrders[k], resAux)

			res[rangeToRes[keyOrders[k]]] = append(res[rangeToRes[keyOrders[k]]], resAux)

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
