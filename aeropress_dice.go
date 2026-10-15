package main

import (
	"flag"
)

func main() {

	// parsing the flags
	flag.Parse()
	// generating a recipe
	generatedRecipe := generateRecipe()
	//printing the generated recipe
	printRecipe(generatedRecipe)

}
