package functions

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// piggyBank storing the dollar amount and number of each coin in a piggy bank
type piggyBank struct {
	amount   float32 // total dollar amount stored
	dimes    int     // number of dimes
	quarters int     // number of quarters
	nickels  int     // number of nickels
	pennies  int     // number of pennies
}

// Creates a bank and fills it with random coins
func FillPiggyBank() {
	// Create a new random source seeded with the current time
	randomSource := rand.NewSource(time.Now().UnixNano())

	// Create a new random generator using the source
	randomGenerator := rand.New(randomSource)
	// initializes the struct with 0s
	testBank := piggyBank{0.0, 0, 0, 0, 0}
	// creates an immutable float array of size 4
	coins := [4]float32{.10, .25, .05, .01}

	// go while loop
	for testBank.amount < 20.0 {
		// random int index generation
		new_coin := coins[randomGenerator.Intn(len(coins))]
		// use & to pass by reference
		storeInPiggyBank(&testBank, new_coin)
	}

	fmt.Println("Savings goals reached!")
	printPiggyBank(&testBank)
}

// Stores a the given coin in the given piggy bank
//
//	depositBank(*piggyBank): Pointer to the piggy bank
//	coin(float32): The coin to deposit
func storeInPiggyBank(depositBank *piggyBank, coin float32) { // pass by reference with *
	depositBank.amount += coin
	// rounds off float precision errors
	depositBank.amount = float32(math.Round(float64(depositBank.amount)*100) / 100)
	coinName := ""
	if coin == 0.10 {
		depositBank.dimes++
		coinName = "Dime"
	} else if coin == 0.25 {
		depositBank.quarters++
		coinName = "Quarter"
	} else if coin == 0.05 {
		depositBank.nickels++
		coinName = "Nickel"
	} else if coin == 0.01 {
		depositBank.pennies++
		coinName = "Penny"
	}
	fmt.Println(coinName, "deposited, piggy bank now has ", depositBank.amount, " dollars.")
}

// Prints the contents of the piggy bank
//
//	bank(*piggyBank): Pointer to the piggy bank
func printPiggyBank(bank *piggyBank) {
	fmt.Println("The piggy bank contains ", bank.amount, " dollars.")
	fmt.Println(bank.quarters, " quarters")
	fmt.Println(bank.dimes, " dimes")
	fmt.Println(bank.nickels, " nickels")
	fmt.Println(bank.pennies, " pennies")
}
