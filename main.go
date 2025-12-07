/*
 * Author: Rabira Motuma
 * Version: 1.0.0
 * Date: 2025-12-05
 * Fileoverview: This program asks the user for the number of bolts, nuts, and washers in their purchase
 * and then calculates and prints out the total cost of the order.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// variables
	var boltsAsString string
  var nutsAsString string
  var washersAsString string
  var boltsAsNumber int
  var nutsAsNumber int
  var washersAsNumber int
  var priceOfBolts int
  var priceOfNuts int
  var priceOfWashers int
  var totalPrice int

	reader := bufio.NewReader(os.Stdin)

	// inputs
	fmt.Print("How many bolts would you like to purchase? " )
	boltsAsString, _ = reader.ReadString('\n')
	boltsAsString = strings.TrimSpace(boltsAsString)
	boltsAsNumber, _ = strconv.Atoi(boltsAsString)

	fmt.Print("How many nuts would you like to purchase? ")
	nutsAsString, _ = reader.ReadString('\n')
	nutsAsString = strings.TrimSpace(nutsAsString)
	nutsAsNumber, _ = strconv.Atoi(nutsAsString)

	fmt.Print("How many washers would you like to purchase? ")
	washersAsString, _ = reader.ReadString('\n')
	washersAsString = strings.TrimSpace(washersAsString)
	washersAsNumber, _ = strconv.Atoi(washersAsString)

	// process
	priceOfBolts = (boltsAsNumber*10)
	priceOfNuts = (nutsAsNumber*5)
	priceOfWashers = (washersAsNumber*3)
	totalPrice = (priceOfBolts+priceOfNuts+priceOfWashers)

	// output
	fmt.Printf("Number of bolts: %5d\n", boltsAsNumber)
	fmt.Printf("Number of nuts: %5d\n", nutsAsNumber)
	fmt.Printf("Number of washers: %3d\n", washersAsNumber)


	// check parts
	if (boltsAsNumber>nutsAsNumber) {
		  fmt.Println("Check the Order - not enough nuts for the bolts you purchased.")
	} else if (washersAsNumber<boltsAsNumber) {
		  fmt.Println("Check the Order - not enough washers for the bolts you purchased.")
	} else {
		  fmt.Println("Order is OK.")
	}

	// mandatory output
	fmt.Printf("Your total cost of the order is %d cents.\n", totalPrice)

	fmt.Println("\nDone.")
}
