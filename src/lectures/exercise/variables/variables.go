//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	favouriteColor := "white"
	fmt.Println("my favourite color is " + favouriteColor)
	birthYear, age := 1991, 31
	fmt.Println(fmt.Sprintf("I am %d years old, having been born in %d", age, birthYear))
	var (
		initialFirst = 'R'
		initialLast  = 'B'
	)
	initials := fmt.Sprintf("%s.%s.", string(initialFirst), string(initialLast))
	fmt.Println(fmt.Sprintf("My initials are %s", initials))
	var ageInDays int
	ageInDays = age * 365
	fmt.Println(fmt.Sprintf("I am more than %d days old!", ageInDays))
	maxAge := 35
	maxDays := maxAge * 365
	remainingDays := maxDays - ageInDays
	fmt.Println(fmt.Sprintf("I have to make the most of my next %d days... because after I will be %d", remainingDays, maxAge))
}
