package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "log"
    "strings"
    "unicode"
)

const CRATE_INPUT_SIZE = 4

type bucket []uint8

type move struct {
    numMoved int
    srcStack int
    dstStack int
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

    scanner := bufio.NewScanner(f)

    buckets := getStacks(scanner)

    moves := getMoves(scanner)

    for _, move := range moves {
        makeMove(&buckets, move)
    }
    getTopCrates(buckets)
}

func makeMove(buckets *[]bucket, move move) {
    for i := 0; i < move.numMoved; i++ {
        srcBucket := &(*buckets)[move.srcStack-1]
        dstBucket := &(*buckets)[move.dstStack-1]
        valMoved := (*srcBucket)[len(*srcBucket)-1]
        (*dstBucket) = append(*dstBucket, valMoved)
        (*srcBucket) = (*srcBucket)[:len(*srcBucket)-1]
    }
}

func getStacks(scan *bufio.Scanner) []bucket {
    // Create a map of int->bucket since we do not know how many
    // buckets there are yet
    bucketsMap := make(map[int]bucket)
    for scan.Scan() {
        line := scan.Text()
        if strings.TrimSpace(line) == "" {
            break
        }
        for i := 1; i < len(line); i += 4 {
            bucket := i / CRATE_INPUT_SIZE
            val := line[i]
            if string(val) != " " && !unicode.IsDigit(rune(val)){
                bucketsMap[bucket] = append(bucketsMap[bucket], val)
            }
        }
    }

    // Convert from map to 2D array
    buckets := make([]bucket, len(bucketsMap))
    for i := 0; i < len(bucketsMap); i++ {
        if bucketFromMap, ok := bucketsMap[i]; ok {
            buckets[i] = make(bucket, len(bucketFromMap))
            for x, crate := range bucketFromMap {
                buckets[i][len(bucketFromMap) - x - 1] = crate
            }
        }
    }
    return buckets
}

func getMoves(scan *bufio.Scanner) []move {
    var moves []move
    var numMoved, src, dst int
    for scan.Scan() {
        line := scan.Text()
        n, err := fmt.Sscanf(line, "move %d from %d to %d\n", &numMoved, &src, &dst)
        if err == io.EOF {
            break
        } else if err != nil {
            log.Fatalf("unable to scan for move: %q", err)
        }
        if n != 3 {
            log.Fatal("did not scan 3 elements")
        }
        moves = append(moves, move{
            numMoved: numMoved,
            srcStack: src,
            dstStack: dst,
        })
    }
    return moves
}

func printBuckets(buckets []bucket) {
    for i := 0; i < len(buckets); i++ {
        fmt.Print(i, ": ")
        for x, val := range buckets[i] {
            fmt.Print(string(val))
            if x + 1 != len(buckets[i]) {
                fmt.Print(",")
            }
        }
        fmt.Print("\n")
    }
}

func getTopCrates(buckets []bucket) {
    for i := 0; i < len(buckets); i++ {
        fmt.Print(string(buckets[i][len(buckets[i])-1]))
    }
}