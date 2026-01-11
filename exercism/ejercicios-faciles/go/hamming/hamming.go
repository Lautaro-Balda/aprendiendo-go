package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("La longitud de los strings es diferente!")
	}
	var hammingCounter int
	for i := range a {
		if a[i] != b[i] {
			hammingCounter++

		}
	}
	return hammingCounter, nil
}
