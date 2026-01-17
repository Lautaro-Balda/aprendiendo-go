package protein

import (
	"errors"
)

var ErrStop = errors.New("STOP")
var ErrInvalidBase = errors.New("base inválida")

var mapaCodons = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

func FromRNA(rna string) ([]string, error) {
	var codons = []string{}
	var codon string
	if len(rna)%3 != 0 {
		return codons, ErrInvalidBase
	}
	for _, val := range rna {
		codon += string(val)
		if len(codon) == 3 {
			val, ok := mapaCodons[codon]
			if val == "STOP" {
				break
			}
			if !ok {
				return codons, ErrInvalidBase
			}
			codons = append(codons, val)
			codon = ""
		}
	}
	return codons, nil
}

func FromCodon(codon string) (string, error) {
	val, ok := mapaCodons[codon]
	if !ok || len(codon) != 3 {
		return "", ErrInvalidBase
	}
	if val == "STOP" {
		return "", ErrStop
	}
	return val, nil
}
