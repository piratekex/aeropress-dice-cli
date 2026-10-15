package main

type recipe struct {
	grindSize    string
	brewTime     int
	brewTemp     int
	coffeeAmount int
	waterAmount  int
	method       string
	bloomTime    int
	bloomAmount  int
	agitation    string
	dilute       bool
}

func generateRecipe() recipe {

	/*
	 * Dice rolls
	 */

	//grind roll
	grindRoll := diceRoll(grindFaces)

	//temperature roll
	temperatureRoll := diceRoll(temperatureFaces)

	//ratio roll
	ratioRoll := diceRoll(ratioFaces)

	//method roll
	methodRoll := diceRoll(methodFaces)

	//agitation roll
	agitationRoll := diceRoll(agitationFaces)

	// creating a recipe and populating the fields from the dice rolls
	createdRecipe := recipe{
		grindSize:    grindRoll.size,
		brewTime:     grindRoll.time,
		brewTemp:     temperatureRoll.celsius,
		coffeeAmount: ratioRoll.coffee,
		waterAmount:  ratioRoll.water,
		method:       methodRoll.method,
		bloomTime:    methodRoll.time,
		bloomAmount:  methodRoll.water,
		agitation:    agitationRoll,
		dilute:       ratioRoll.dilute,
	}

	return createdRecipe
}
