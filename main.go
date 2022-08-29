package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxTurns = 5
	usage    = `Welcome to the Lucky Number game!

The program will pick %d random numbers.
Your mission is to guess one of those numbers

The greater you number is, the harder it gets.

Wanna play?`
)

func main() {
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Printf(usage, maxTurns)
		return
	}

	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Not a number.")
		return
	}

	if guess <= 0 {
		fmt.Println("Please pick a positive number.")
		return
	}

	min := 10
	if guess > min {
		min = guess
	}

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(guess) + 1

		if n == guess {
			if turn == 1 {
				fmt.Println("First Guess Winner!")
			} else {
				fmt.Println("You Won!")
			}
			return
		}
	}
	fmt.Println("You Lost...Try again?")

}
