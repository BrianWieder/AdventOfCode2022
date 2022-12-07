package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        log.Fatal("missing file to read in args")
    }
    f, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatalf("unable to open %s: %w", os.Args[1], err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)
    if !scanner.Scan() {
        log.Fatal("failed to scan for input!")
    }
    input := scanner.Text()
    fmt.Println(processDatastream(input))
}

func processDatastream(data string) int {
    recentFour := make(map[uint8]int, 4)
    for i := 0; i < len(data); i++ {
        if i < 4 {
            recentFour[data[i]] += 1
        } else {
            oneToDecrement := data[i-4]
            recentFour[oneToDecrement] -= 1
            if recentFour[oneToDecrement] == 0 {
                delete(recentFour, oneToDecrement)
            }
            recentFour[data[i]] += 1
            if len(recentFour) == 4 {
                return i + 1
            }
        }
    }
    fmt.Print(recentFour)
    return 0
}