package lasagna

// PreparationTime calculates the preparation time for a given number of layers
func PreparationTime(layers []string, avgPreparationTime int) int {
	if avgPreparationTime == 0 {
		avgPreparationTime = 2
	}
	return len(layers) * avgPreparationTime
}

// Quantities calculates the required amounts of noodles and sauce for a given number of layers
func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0
	for _, layer := range layers {
		if layer == "noodles" {
			noodles += 50
		}
		if layer == "sauce" {
			sauce += 0.2
		}
	}
	return noodles, sauce
}

// AddSecretIngredient copy the hidden ingredient from my friends list to my list
func AddSecretIngredient(friendList, myList []string) {
	myList[len(myList)-1] = friendList[len(friendList)-1]
}

// ScaleRecipe scale a recipe based on how many portions are desired
func ScaleRecipe(amounts []float64, portions int) []float64 {
	multiplier := float64(portions) / 2.0
	var resultAmounts []float64
	for _, amount := range amounts {
		resultAmounts = append(resultAmounts, amount*multiplier)
	}
	return resultAmounts
}
