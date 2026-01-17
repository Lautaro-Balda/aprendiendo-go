package proverb

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	resultado := make([]string, 0, len(rhyme))
	if len(rhyme) == 0 {
		return resultado
	}
	for i := 0; i < len(rhyme)-1; i++ {
		resultado = append(resultado, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))

	}
	resultado = append(resultado, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))

	return resultado
}
