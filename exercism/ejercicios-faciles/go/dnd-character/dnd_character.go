package dndcharacter

import (
	"math"
	"math/rand"
	"sort"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {

	return int(math.Floor(float64(score-10) / 2))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	var diceRolls [4]int
	for i := range diceRolls {
		diceRolls[i] = rand.Intn(6) + 1
	}

	sort.Ints(diceRolls[:])
	return diceRolls[1] + diceRolls[2] + diceRolls[3]

}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	char := Character{Strength: Ability(), Dexterity: Ability(), Constitution: Ability(), Intelligence: Ability(), Wisdom: Ability(), Charisma: Ability(), Hitpoints: 10}
	char.Hitpoints += Modifier(char.Constitution)
	return char
}
