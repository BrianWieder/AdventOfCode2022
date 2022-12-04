package main

import (
    "fmt"
    "io"
    "log"
    "os"
)

const OPP_ROCK = "A"
const OPP_PAPER = "B"
const OPP_SCISSORS = "C"

const PLAYER_ROCK = "X"
const PLAYER_PAPER = "Y"
const PLAYER_SCISSORS = "Z"

func main() {
    if len(os.Args) < 2 {
        log.Fatal("missing file to read in args")
    }
    f, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatalf("unable to open %s: %w", os.Args[1], err)
    }

    defer f.Close()

    var score int

    var opp string
    var player string

    for {
        n, err := fmt.Fscanln(f, &opp, &player)
        if err == io.EOF {
            break
        }
        if n != 2 {
            log.Fatal("fscanln did not return 2 items")
        }
        score += getScore(opp, player)
    }
    fmt.Print(score)
}

func getScore(opp, player string) int {
    score := 0
    switch player {
    case PLAYER_ROCK:
        score += 1
    case PLAYER_PAPER:
        score += 2
    case PLAYER_SCISSORS:
        score += 3
    }
    switch opp {
    case OPP_ROCK:
        switch player {
        case PLAYER_ROCK:
            score += 3
        case PLAYER_PAPER:
            score += 6
        case PLAYER_SCISSORS:
            score += 0
        }
    case OPP_PAPER:
        switch player {
        case PLAYER_ROCK:
            score += 0
        case PLAYER_PAPER:
            score += 3
        case PLAYER_SCISSORS:
            score += 6
        }
    case OPP_SCISSORS:
        switch player {
        case PLAYER_ROCK:
            score += 6
        case PLAYER_PAPER:
            score += 0
        case PLAYER_SCISSORS:
            score += 3
        }
    default:
        log.Fatalf("unsupported opp move: %s", opp)
    }
    return score
}