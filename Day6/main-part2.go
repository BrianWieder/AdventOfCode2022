package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)
const N_DISTINCT = 14

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
    recent := make(map[uint8]int, N_DISTINCT)
    for i := 0; i < len(data); i++ {
        if i < N_DISTINCT {
            recent[data[i]] += 1
        } else {
            oneToDecrement := data[i-N_DISTINCT]
            recent[oneToDecrement] -= 1
            if recent[oneToDecrement] == 0 {
                delete(recent, oneToDecrement)
            }
            recent[data[i]] += 1
            if len(recent) == N_DISTINCT {
                return i + 1
            }
        }
    }
    fmt.Print(recent)
    return 0
}