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
        rucksack1 := scanner.Text()
        scanner.Scan()
        rucksack2 := scanner.Text()
        scanner.Scan()
        rucksack3 := scanner.Text()

        firstSackMap := make(map[uint8]int)
        secondSackMap := make(map[uint8]int)
        thirdSackMap := make(map[uint8]int)

        for i := 0; i < len(rucksack1); i++ {
            firstSackMap[rucksack1[i]] += 1
        }
        for i := 0; i < len(rucksack2); i++ {
            secondSackMap[rucksack2[i]] += 1
        }
        for i := 0; i < len(rucksack3); i++ {
            thirdSackMap[rucksack3[i]] += 1
        }

        common := make(map[uint8]bool)

        for key, _ := range firstSackMap {
            if _, ok := secondSackMap[key]; ok {
                common[key] = true
            }
        }

        for key, _ := range common {
            if _, ok := thirdSackMap[key]; ok {
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