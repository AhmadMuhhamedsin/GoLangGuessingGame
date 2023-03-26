package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	secret := getRandomNumber()
	for matching := false; !matching; {
		guess := getUserInput()
		// fmt.Println(secret, guess)
		matching = isMatching(secret, guess)
		// fmt.Println(matching)
	}

}
func isMatching(secret, guess int) bool {
	if guess > secret {
		fmt.Println("Your guess is too big ")
		return false
	} else if guess < secret {
		fmt.Println("Your guess is too small ")
		return false
	} else {
		fmt.Println("Your guess is correct ")
		return true
	}

}
func getUserInput() int {
	fmt.Println("Please Input your guess: ")
	var input int
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Failed to parse your input")
	} else {
		fmt.Println("You guessed", input)
	}
	return input
}
func getRandomNumber() int {
	rand.Seed(time.Now().Unix())
	return rand.Int() % 11
}
