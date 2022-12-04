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

    var prioritySum int

    for scanner.Scan() {
        rucksack := scanner.Text()

        firstSack := rucksack[:len(rucksack)/2]
        secondSack := rucksack[len(rucksack)/2:]

        firstSackMap := make(map[uint8]int)
        secondSackMap := make(map[uint8]int)

        for i := 0; i < len(firstSack); i++ {
            firstSackMap[firstSack[i]] += 1
            secondSackMap[secondSack[i]] += 1
        }

        for key, _ := range firstSackMap {
            if _, ok := secondSackMap[key]; ok {
                prioritySum += getPriority(key)
            }
        }
    }

    fmt.Println(prioritySum)
}

func getPriority(item uint8) int {
    if item >= 'a' && item <= 'z' {
        return int(item - 'a') + 1
    } else if item >= 'A' && item <= 'Z' {
        return int(item - 'A') + 27
    }
    return 0
}