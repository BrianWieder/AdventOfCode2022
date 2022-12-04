package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
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

    var currElf int
    var maxCals [3]int

    for scanner.Scan() {
        line := scanner.Text()
        if strings.TrimSpace(line) == "" {
            checkForTopThree(&maxCals, currElf)
            currElf = 0
            continue
        }
        currCalories, err := strconv.Atoi(scanner.Text())
        if err != nil {
            log.Fatalf("unable to convert %s to int: %w", scanner.Text(), err)
        }
        currElf += currCalories
    }
    checkForTopThree(&maxCals, currElf)
    var total int
    for _, cal := range maxCals {
        total += cal
    }
    fmt.Print(total)
}

func checkForTopThree(maxCals *[3]int, currCals int) {
    if currCals > maxCals[0] {
        maxCals[0], maxCals[1], maxCals[2] = currCals, maxCals[0], maxCals[1]
    } else if currCals > maxCals[1] {
        maxCals[1], maxCals[2] = currCals, maxCals[1]
    } else if currCals > maxCals[2] {
       maxCals[2] = currCals
    }
}