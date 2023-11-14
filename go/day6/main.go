package main

import (
    "bufio"
    "errors"
    "fmt"
    "log"
    "os"
)

type vec struct {
    arr [4]byte
}

func (vec *vec) contains(c byte) bool {
    for _, e := range vec.arr {
        if  e == c {
            return true;
        }
    }
    return false
}
func (vec *vec) append(c byte) {
    
}

func part1(data string) (int, error) {
    var distinc_char vec
    for _, c := range data {
        if !distinc_char.contains(byte(c)) {

        }
    }
    return 0, errors.New("No j in given string")
}


func main() {

    file, err := os.Open("data.txt")
    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        pos, err := part1(scanner.Text())
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("Positon of j: %v\n", pos)
    }

}
