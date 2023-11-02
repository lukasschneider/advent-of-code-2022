package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "day5/stack"
)

var s1 = stack.New[byte]()
var s2 = stack.New[byte]()
var s3 = stack.New[byte]()
var s4 = stack.New[byte]()
var s5 = stack.New[byte]()
var s6 = stack.New[byte]()
var s7 = stack.New[byte]()
var s8 = stack.New[byte]()
var s9 = stack.New[byte]()


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
    createDS()
    fmt.Println(s1.Size())
    fmt.Println(s2.Size())
}

func createDS() {
    s1.Push('Q')
    s1.Push('W')
    s1.Push('P')
    s1.Push('S')
    s1.Push('Z')
    s1.Push('R')
    s1.Push('H')
    s1.Push('D')

    s2.Push('V')
    s2.Push('B')
    s2.Push('R')
    s2.Push('W')
    s2.Push('Q')
    s2.Push('H')
    s2.Push('F')

    s3.Push('C')
    s3.Push('V')
    s3.Push('S')
    s3.Push('H')

    s4.Push('H')
    s4.Push('F')
    s4.Push('G')

    s5.Push('P')
    s5.Push('G')
    s5.Push('J')
    s5.Push('B')
    s5.Push('Z')

}
