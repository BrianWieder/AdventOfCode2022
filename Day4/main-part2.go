package main

import (
    "fmt"
    "io"
    "log"
    "os"
)

type assignment struct {
    start int
    end int
}

func main() {
    if len(os.Args) < 2 {
        log.Fatal("missing file to read in args")
    }
    f, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatalf("unable to open %s: %w", os.Args[1], err)
    }

    defer f.Close()

    var numOverlap int

    for {
        var firstAssignment assignment
        var secondAssignment assignment
        n, err := fmt.Fscanf(f, "%d-%d,%d-%d\n", &firstAssignment.start, &firstAssignment.end, &secondAssignment.start, &secondAssignment.end)

        if err == io.EOF {
            break
        }

        if n != 4 {
            log.Fatal("could not scan 4 items")
        }
        if overlaps(firstAssignment, secondAssignment) {
            numOverlap += 1
        }
    }

    fmt.Println(numOverlap)
}

func overlaps(firstAssignment, secondAssignment assignment) bool {
    // Second Assignment start overlaps with first assignment
    if firstAssignment.start <= secondAssignment.start && firstAssignment.end >= secondAssignment.start {
        return true
    }
    // Second Assignment end overlaps with first assignment
    if firstAssignment.start <= secondAssignment.end && firstAssignment.end >= secondAssignment.end {
        return true
    }
    // First Assignment start overlaps with second assignment
    if secondAssignment.start <= firstAssignment.start && secondAssignment.end >= firstAssignment.start {
        return true
    }
    // First Assignment start overlaps with second assignment
    if secondAssignment.start <= firstAssignment.end && secondAssignment.end >= firstAssignment.end {
        return true
    }
    return false
}