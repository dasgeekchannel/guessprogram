package main

import (
	"fmt"
	"crypto/rand"
	"math/big"
	"strconv"
	"os"
)

func randomNumber() int64 {
	res, _ := rand.Int(rand.Reader, big.NewInt(100))
	return res.Int64()

}

func intro() {
    fmt.Println("------------------------------------")
	fmt.Println("       Guess The Number Game        ")
	fmt.Println("------------------------------------")
	fmt.Println("Guess a number between 0 and 100:")
}

func userInput() int64 {
	var input string
	fmt.Scanln(&input)
	inputInt ,_ := strconv.Atoi(input)
	inputInt64 := int64(inputInt)
	return inputInt64
}

func main() {
	intro()
	theNumber := randomNumber()
	for x := 0 ; x < 1000 ; x++{ // if u need a thousand guesses .....
		theGuess := userInput()
		if theGuess < theNumber {
			fmt.Printf("%d is too low\n", theGuess)
		}
		if theGuess > theNumber {
			fmt.Printf("%d is too high\n", theGuess)
		}
		if theGuess == theNumber {
			fmt.Printf("you win\nyour score was %d\n", x+1)
			os.Exit(0)
		}
	}

}
