package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
    "time"
)

func FindCommonChar(stringsToCompare ...string) uint32 {
	if len(stringsToCompare) < 1 {
		return 0
	}

	firstString := stringsToCompare[0]

	for _, char := range firstString {
		charExistsInAll := true

		for _, str := range stringsToCompare[1:] {
			if !strings.ContainsRune(str, char) {
				charExistsInAll = false
				break
			}
		}

		if charExistsInAll {
			if char >= 65 && char <= 90 {
				return uint32(char - 38)
			} else {
				return uint32(char - 96)
			}
		}
	}

	return 0
}

func part1(data []string) uint32 {
    start := time.Now()
	var score uint32 = 0

	for _, line := range data {
		slice := len(line) / 2
		comp1 := line[:slice]
		comp2 := line[slice:]
		score += FindCommonChar(comp1, comp2)
	}
    fmt.Printf("Part1: %v\n", time.Since(start))

	return score
}

func part2(data []string) uint32 {
    start := time.Now()
	var score uint32 = 0
	for i := 0; i < len(data); i += 3 {
		score += FindCommonChar(data[i], data[i+1], data[i+2])
	}
    fmt.Printf("Part2: %v\n", time.Since(start))
	return score
}

func main() {
	file, err := os.Open("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal("failed to close File")
		}
	}(file)

	scanner := bufio.NewScanner(file)
	var data []string

	for i := 0; scanner.Scan(); i++ {
		data = append(data, scanner.Text())
	}

	fmt.Printf("Part1: Score: %v\n", part1(data))
	fmt.Printf("Part2: Score: %v\n", part2(data))

}
