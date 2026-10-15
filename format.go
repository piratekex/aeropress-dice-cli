package main

import (
	"flag"
	"fmt"
)

//time format

func formatTime(seconds int) string {
	minute := seconds / 60
	second := seconds % 60
	return fmt.Sprintf("%d:%02d", minute, second)
}

// converting celsius to fahrenheit
func toFahrenheit(celsius int) int {

	return (celsius * 9 / 5) + 32
}

// creating the -f fahrenheit flag
var fahrenheit = flag.Bool("f", false, "display the temperature in Fahrenheit")

func printRecipe(generatedRecipe recipe) {

	/*
	 * printing the generated recipe in to the terminal
	 */

	//grind dice
	if generatedRecipe.brewTime == 0 {
		fmt.Println("Grind size: Your choice")
		fmt.Println("Brew time: Your choice")
	} else {
		fmt.Println("Grind size:", generatedRecipe.grindSize)
		fmt.Println("Brew time:", formatTime(generatedRecipe.brewTime))
	}

	//temperature dice
	if generatedRecipe.brewTemp == 0 {
		fmt.Println("Brew temperature: Your choice")
	} else if *fahrenheit {
		fmt.Println("Brew temperature:", toFahrenheit(generatedRecipe.brewTemp), "°F")
	} else {
		fmt.Println("Brew temperature:", generatedRecipe.brewTemp, "°C")
	}

	//ratio dice
	if generatedRecipe.coffeeAmount == 0 {
		fmt.Println("Coffee amount: Your choice")
		fmt.Println("Water amount: Your choice")
	} else if !generatedRecipe.dilute {
		fmt.Println("Coffee amount:", generatedRecipe.coffeeAmount, "g")
		fmt.Println("Water amount:", generatedRecipe.waterAmount, "g")
	} else {
		fmt.Println("Coffee amount:", generatedRecipe.coffeeAmount, "g")
		fmt.Println("Water amount:", generatedRecipe.waterAmount, "g", "+ Dilute")
	}

	//method dice
	if generatedRecipe.bloomAmount == 0 {
		fmt.Println("Method:", generatedRecipe.method)
		fmt.Println("Bloom:", "No bloom")
	} else {
		fmt.Println("Method:", generatedRecipe.method)
		fmt.Println("Bloom water:", generatedRecipe.bloomAmount, "g")
		fmt.Println("Bloom time:", formatTime(generatedRecipe.bloomTime))
	}

	//agitaton dice
	fmt.Println("Agitation:", generatedRecipe.agitation)
}
