package main

import (
	"bufio"
	"day5/stack"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type instruction struct {
    number byte
    src byte
    dest byte
}

var instructionArr []*instruction


func NewInstruction(s []string) (*instruction, error) {

    if len(s) == 0 {
        return nil, errors.New("string is empty")
    }

    number, err := strconv.ParseInt(s[1], 10, 8)
    if err != nil {
        return nil, errors.New("there is no number in string at position 1")
    }

    src, err := strconv.ParseInt(s[3], 10, 8)
    if err != nil {
        return nil, errors.New("there is no number in string at position 3")
    }

    dest, err := strconv.ParseInt(s[5], 10, 8)
    if err != nil {
        return nil, errors.New("there is no number in string at position 5")
    }

    ins := &instruction{
        number: byte(number),
        src: byte(src),
        dest: byte(dest),
    }

    return ins, nil
}

func cratemoover9000(insLookUp map[int]*stack.Stack[byte], ins *instruction) error {
    srcStack := insLookUp[int(ins.src)]
    destStack := insLookUp[int(ins.dest)]
    for i := 1; i <= int(ins.number); i++ {
        srcElement, err := srcStack.Pop()
        if err != nil {
            continue
        }
        destStack.Push(srcElement)
    }
    return nil
}

func cratemoover9001(insLookUp map[int]*stack.Stack[byte], ins *instruction) error {
    var cache []byte
    srcStack := insLookUp[int(ins.src)]
    destStack := insLookUp[int(ins.dest)]
    for i := 1; i <= int(ins.number); i++ {
        srcElement, err := srcStack.Pop()
        if err != nil {
            continue
        }
        cache = append(cache, srcElement)
    }
    for i := len(cache)-1; i >= 0; i-- {
        destStack.Push(cache[i])
    }
    return nil
}

func part1() {

    start := time.Now()
    var insLookUp = map[int]*stack.Stack[byte] {
        1: stack.New[byte]('Q', 'W', 'P', 'S', 'Z', 'R', 'H', 'D'),
        2: stack.New[byte]('V', 'B', 'R', 'W', 'Q', 'H', 'F'),
        3: stack.New[byte]('C', 'V', 'S', 'H'),
        4: stack.New[byte]('H', 'F', 'G'),
        5: stack.New[byte]('P', 'G', 'J', 'B', 'Z'),
        6: stack.New[byte]('Q', 'T', 'J', 'H', 'W', 'F', 'L'),
        7: stack.New[byte]('Z', 'T', 'W', 'D', 'L', 'V', 'J', 'N'),
        8: stack.New[byte]('D', 'T', 'Z', 'C', 'J', 'G', 'H', 'F'),
        9: stack.New[byte]('W', 'P', 'V', 'M', 'B', 'H'),
    }

    for _, e := range instructionArr {
        err := cratemoover9000(insLookUp, e) 
        if err != nil {
            log.Fatal(err)
        }
    }

    result, err := result(insLookUp)
    if err != nil {
        log.Fatal(err, "Could not get result")
    }
    fmt.Printf("Part 1 took %s\n", time.Since(start))
    fmt.Printf("Result: %s\n", result)
}

func part2() {
    start := time.Now()
    var insLookUp = map[int]*stack.Stack[byte] {
        1: stack.New[byte]('Q', 'W', 'P', 'S', 'Z', 'R', 'H', 'D'),
        2: stack.New[byte]('V', 'B', 'R', 'W', 'Q', 'H', 'F'),
        3: stack.New[byte]('C', 'V', 'S', 'H'),
        4: stack.New[byte]('H', 'F', 'G'),
        5: stack.New[byte]('P', 'G', 'J', 'B', 'Z'),
        6: stack.New[byte]('Q', 'T', 'J', 'H', 'W', 'F', 'L'),
        7: stack.New[byte]('Z', 'T', 'W', 'D', 'L', 'V', 'J', 'N'),
        8: stack.New[byte]('D', 'T', 'Z', 'C', 'J', 'G', 'H', 'F'),
        9: stack.New[byte]('W', 'P', 'V', 'M', 'B', 'H'),
    }

    for _, e := range instructionArr {
        err := cratemoover9001(insLookUp, e) 
        if err != nil {
            log.Fatal(err)
        }
    }

    result, err := result(insLookUp)
    if err != nil {
        log.Fatal(err, "Could not get result")
    }

    fmt.Printf("Part 2 took %s\n", time.Since(start))
    fmt.Printf("Result: %s\n", result)
}

func result(insLookUp map[int]*stack.Stack[byte]) ([]byte, error) {
    var result []byte
    var keys []int
    for k,_ := range insLookUp {
        keys = append(keys, k)
    }
    sort.Ints(keys)
    for _, k := range keys {
        stack := insLookUp[k]
        v, err := stack.Peek()
        if err != nil {
            return nil, err 
        }
        result = append(result, v)
    }
    return result, nil
}

func main() {

    file, err := os.Open("data.txt")
    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)
    lineCount := 0    

    for scanner.Scan() {
        lineCount++
        if lineCount >= 11 {
            line := scanner.Text()
            lineArr := strings.Split(line, " ") 
            ins, err := NewInstruction(lineArr)
            if err != nil {
                log.Fatal(err)
            }
            instructionArr = append(instructionArr, ins)
        }
    }

    part1()
    part2()
}

