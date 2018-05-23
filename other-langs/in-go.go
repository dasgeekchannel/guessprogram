package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

var (
	highNumber int64 = 100
)

func randomNumber() int64 {
	res, _ := rand.Int(rand.Reader, big.NewInt(highNumber))
	return res.Int64()

}

func intro() {
	fmt.Println("------------------------------------")
	fmt.Println("       Guess The Number Game        ")
	fmt.Println("------------------------------------")
	fmt.Printf("Guess a number between 0 and %d:\n", highNumber)
}

func userInput() int64 {
	var input string
	fmt.Scanln(&input)
	inputInt, _ := strconv.Atoi(input)
	inputInt64 := int64(inputInt)
	return inputInt64
}

func main() {
	intro()
	theNumber := randomNumber()
	for x := 0; x < int(highNumber); x++ {
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
