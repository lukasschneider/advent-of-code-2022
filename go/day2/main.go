package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type game struct {
	Rock     byte
	Paper    byte
	Scissors byte
}

type predict struct {
	Loss byte
	Draw byte
	Win  byte
}

const (
	ROCK     = 1
	PAPER    = 2
	SCISSORS = 3
	WON      = 6
	DRAW     = 3
	LOSS     = 0
)

func CalcScore(data []string) uint32 {

	enemy := game{
		Rock:     'A',
		Paper:    'B',
		Scissors: 'C',
	}

	player := game{
		Rock:     'X',
		Paper:    'Y',
		Scissors: 'Z',
	}
	var score uint32 = 0

	var scoreTable = map[byte]map[byte]uint32{
		enemy.Rock: {
			player.Paper:    WON + PAPER,
			player.Rock:     DRAW + ROCK,
			player.Scissors: LOSS + SCISSORS,
		},
		enemy.Scissors: {
			player.Paper:    LOSS + PAPER,
			player.Rock:     WON + ROCK,
			player.Scissors: DRAW + SCISSORS,
		},
		enemy.Paper: {
			player.Paper:    DRAW + PAPER,
			player.Rock:     LOSS + ROCK,
			player.Scissors: WON + SCISSORS,
		},
	}

	for _, line := range data {
		first := line[0]
		last := line[2]

		scoreMap, ok := scoreTable[first]
		if ok {
			value, ok := scoreMap[last]
			if ok {
				score += value
			} else {
				log.Fatalf("last not found in scoreMap: %c\n", last)
			}
		} else {
			log.Fatalf("first not found in scoreTable: %c\n", first)
		}
	}

	return score
}

func CalcScorePredict(data []string) uint32 {
	var score uint32 = 0
	enemy := game{
		Rock:     'A',
		Paper:    'B',
		Scissors: 'C',
	}

	decider := predict{
		Loss: 'X',
		Draw: 'Y',
		Win:  'Z',
	}

	var scoreTable = map[byte]map[byte]uint32{
		enemy.Rock: {
			decider.Loss: LOSS + SCISSORS,
			decider.Draw: DRAW + ROCK,
			decider.Win:  WON + PAPER,
		},
		enemy.Scissors: {
			decider.Loss: LOSS + PAPER,
			decider.Draw: DRAW + SCISSORS,
			decider.Win:  WON + ROCK,
		},
		enemy.Paper: {
			decider.Loss: LOSS + ROCK,
			decider.Draw: DRAW + PAPER,
			decider.Win:  WON + SCISSORS,
		},
	}

	for _, line := range data {
		first := line[0]
		last := line[2]

		scoreMap, ok := scoreTable[first]
		if ok {
			value, ok := scoreMap[last]
			if ok {
				score += value
			} else {
				log.Fatalf("last not found in scoreMap: %c\n", last)
			}
		} else {
			log.Fatalf("first not found in scoreTable: %c\n", first)
		}
	}

	return score
}

func main() {

	var score uint32 = 0
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

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	score = CalcScorePredict(data)
	fmt.Printf("Predict Final Score: %v\n", score)

	score = CalcScore(data)
	fmt.Printf("Final Score: %v\n", score)

}
