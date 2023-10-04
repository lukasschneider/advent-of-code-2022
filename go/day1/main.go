package main 

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "cmp"
    "slices"
)

func main() {

    var tmp uint32 = 0
    data := []uint32{}

    file, err := os.Open("data.txt")
    
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()

        value, err := strconv.ParseUint(line, 10, 32)
        if err == nil {
            tmp = tmp + uint32(value)
        }

        if line == "" {
            data = append(data, tmp)
            tmp = 0
        }
    }

    slices.SortStableFunc(data, func(a, b uint32) int {
        return cmp.Compare(b, a)
    })

    err = scanner.Err()
    if err != nil {
        log.Fatal(err)
    }

    var top uint32 = data[0] + data[1] + data[2]
    fmt.Printf("Max: %v\n", data[0])
    fmt.Printf("Top: %v\n", top)
}
