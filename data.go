package main

import "math/rand/v2"

// teams represents a fixed set of 32 countries taking part in the World Cup.
var teams = [32]string{
	"Netherlands",
	"Senegal",
	"Ecuador",
	"Qatar",
	"England",
	"USA",
	"Iran",
	"Wales",
	"Argentina",
	"Poland",
	"Mexico",
	"Saudi Arabia",
	"France",
	"Australia",
	"Tunisia",
	"Denmark",
	"Japan",
	"Spain",
	"Germany",
	"Costa Rica",
	"Morocco",
	"Croatia",
	"Belgium",
	"Canada",
	"Brazil",
	"Switzerland",
	"Cameroon",
	"Serbia",
	"Portugal",
	"South Korea",
	"Uruguay",
	"Ghana",
}

// mixOrder is a helper function capable of mixing order an array's objects.
func mixOrder(t [32]string) [32]string {
	for i := 0; i < len(t)-1; i++ {
		team := t[i]
		swapIndex := randRange(0, len(t))
		swappedTeam := teams[swapIndex]
		t[i] = swappedTeam
		t[swapIndex] = team
	}

	return t
}

// randRange provides a random value form a specified min-max range.
func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}
