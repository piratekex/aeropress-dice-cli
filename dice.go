package main

import "math/rand/v2"

/*
 * Creating the structs for each dice
 */

// temp dice struct
type tempFace struct {
	celsius int
}

// grind dice struct
type grindFace struct {
	time int
	size string
}

// ratio dice struct
type ratioFace struct {
	coffee int
	water  int
	dilute bool
}

// method dice struct
type methodFace struct {
	method string
	time   int
	water  int
}

/*
 * Creating the functions for the dice
 */

// dice roll

func diceRoll[T any](faces []T) T {
	return faces[rand.IntN(len(faces))]
}

/*
 * Creating the actual dice
 */

// temperature dice
var temperatureFaces = []tempFace{
	{celsius: 95},
	{celsius: 75},
	{celsius: 90},
	{celsius: 80},
	{celsius: 85},
	{celsius: 0},
}

// grind dice
var grindFaces = []grindFace{
	{time: 30, size: "very fine"},
	{time: 60, size: "fine"},
	{time: 90, size: "medium fine"},
	{time: 120, size: "medium"},
	{time: 240, size: "coarse"},
	{size: "your choice"},
}

// ratio dice
var ratioFaces = []ratioFace{
	{coffee: 30, water: 200, dilute: true},
	{coffee: 24, water: 200, dilute: true},
	{coffee: 15, water: 250, dilute: true},
	{coffee: 15, water: 200, dilute: false},
	{coffee: 12, water: 200, dilute: false},
	{dilute: false},
}

// method dice
var methodFaces = []methodFace{
	{method: "inverted", time: 30, water: 60},
	{method: "standard", time: 30, water: 60},
	{method: "inverted", time: 30, water: 30},
	{method: "standard", time: 30, water: 30},
	{method: "standard"},
	{method: "inverted"},
}

// agitation dice
var agitationFaces = []string{
	"No stir",
	"stir ×2 before press",
	"CW once + CCW once before press",
	"stir ×1 before press",
	"N/S/E/W stir",
	"your choice",
}
