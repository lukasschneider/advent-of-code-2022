package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func CreateStack(data []string) {

}

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("%v\n", line)
	}

}
