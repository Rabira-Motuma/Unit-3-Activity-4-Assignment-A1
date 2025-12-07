/**
 * @author Rabira Motuma
 * @version 1.0.0
 * @date 2025-12-05
 * @fileoverview This program asks the user for the number of bolts, nuts, and washers in their purchase
 * and then calculates and prints out the total cost of the order.
 */

// variables
let boltsAsString: string;
let nutsAsString: string;
let washersAsString: string;
let boltsAsNumber: number;
let nutsAsNumber: number;
let washersAsNumber: number;
let priceOfBolts: number;
let priceOfNuts: number;
let priceOfWashers: number;
let totalPrice: number;

// input number of pieces
boltsAsString = prompt("How many bolts would you like to purchase? ") || "0";
boltsAsNumber = parseInt(boltsAsString);

nutsAsString = prompt("How many nuts would you like to purchase? ") || "0";
nutsAsNumber = parseInt(nutsAsString);

washersAsString = prompt("How many washers would you like to purchase? ") || "0";
washersAsNumber = parseInt(washersAsString);

// process
priceOfBolts = (boltsAsNumber * 10);
priceOfNuts = (nutsAsNumber * 5);
priceOfWashers = (washersAsNumber * 3);
totalPrice = (priceOfBolts + priceOfNuts + priceOfWashers);

//output
console.log(``)

// check parts
if (boltsAsNumber > nutsAsNumber) {
    console.log("Check the Order - not enough nuts for the bolts you purchased.");
} else if (washersAsNumber < boltsAsNumber) {
    console.log("Check the Order - not enough washers for the bolts you purchased.");
} else {
  console.log("Order is OK.");
}

// mandatory output
console.log(`Your total cost of the order is ${totalPrice} cents.`);

console.log("\nDone.")
