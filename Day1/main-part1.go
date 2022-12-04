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
    var maxCal int

    for scanner.Scan() {
        line := scanner.Text()
        if strings.TrimSpace(line) == "" {
            if currElf > maxCal {
                maxCal = currElf
            }
            currElf = 0
            continue
        }
        currCalories, err := strconv.Atoi(scanner.Text())
        if err != nil {
            log.Fatalf("unable to convert %s to int: %w", scanner.Text(), err)
        }
        currElf += currCalories
    }
    if currElf > maxCal {
        maxCal = currElf
    }
    fmt.Print(maxCal)
}