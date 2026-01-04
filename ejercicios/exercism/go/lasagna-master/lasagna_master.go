package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, averageTime int) int {
	if averageTime <= 0 {
		averageTime = 2
	}
	return averageTime * len(layers)
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	var nOfNoodles int
	var nOfSauce float64
	for _, value := range layers {
		if value == "noodles" {
			nOfNoodles++
		}
		if value == "sauce" {
			nOfSauce++
		}
	}
	return nOfNoodles * 50, nOfSauce * 0.2
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendIngredients, ownRecipe []string) {
	secretIngredient := friendIngredients[len(friendIngredients)-1]
	ownRecipe[len(ownRecipe)-1] = secretIngredient
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amountsForTwo []float64, portions int) []float64 {
	newAmounts := make([]float64, len(amountsForTwo))
	for i := range amountsForTwo {
		newAmounts[i] = (amountsForTwo[i] / 2) * float64(portions)
	}
	return newAmounts
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
