package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1(data []byte) uint32 {
	var count uint32 = 0
	for i := 0; i < len(data); i += 4 {
		if (data[i] >= data[i+2]) && (data[i+1] <= data[i+3]) {
			count++
		} else if (data[i] <= data[i+2]) && (data[i+1] >= data[i+3]) {
			count++
		}
	}
	return count
}

func part2(data []byte) uint32 {
	var count uint32 = 0

	for i := 0; i < len(data); i += 4 {
		if data[i+1] >= data[i+2] && data[i+3] >= data[i] {
			count++
		}
	}

	return count
}

func main() {
	file, err := os.Open("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal("failed to close file")
		}
	}(file)

	scanner := bufio.NewScanner(file)
	var data []byte

	for scanner.Scan() {

		line := scanner.Text()
		commaPosition := strings.IndexRune(line, ',')

		firstSet := line[:commaPosition]
		secondSet := line[commaPosition+1:]

		dashPosition := strings.IndexRune(firstSet, '-')
		first, err := strconv.ParseInt(firstSet[:dashPosition], 10, 8)
		second, err := strconv.ParseInt(firstSet[dashPosition+1:], 10, 8)

		dashPosition = strings.IndexRune(secondSet, '-')
		third, err := strconv.ParseInt(secondSet[:dashPosition], 10, 8)
		fourth, err := strconv.ParseInt(secondSet[dashPosition+1:], 10, 8)

		if err != nil {
			log.Fatal(err)
		}

		data = append(data, byte(first))
		data = append(data, byte(second))
		data = append(data, byte(third))
		data = append(data, byte(fourth))
	}

	fmt.Printf("contains count: %v\n", part1(data))
	fmt.Printf("overlap count: %v\n", part2(data))
}
