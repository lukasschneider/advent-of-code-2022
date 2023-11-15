package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "time"
    "slices"
)



func findMarker(data string, number int) (int, byte) {
    for i := range data {
        distinct_chars := make([]byte, number)
        j := 0
        for ; j < number; j++ {
            if !slices.Contains(distinct_chars, data[i+j]) {
                distinct_chars[j] = data[i+j]
            } else {
                break;
            }
        }
        if j == number {
            return i+j, data[i+j]
        }
    }
    return 0, 0
}


func main() {

    file, err := os.Open("data.txt")
    if err != nil {
        log.Fatal(err)
    }

    var data string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        data = scanner.Text()
    }

    start := time.Now()
    pos, val := findMarker(data, 4)
    fmt.Printf("Time for part1: %s\n", time.Since(start))
    fmt.Printf("Positon of %c: %v\n",val, pos)

    start = time.Now()
    pos, val = findMarker(data, 14)
    fmt.Printf("Time for part2: %s\n", time.Since(start))
    fmt.Printf("Positon of %c: %v\n",val, pos)
}
