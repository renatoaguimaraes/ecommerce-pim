package service

import "github.com/renatoaguimaraes/ecommerce-pim/internal/model"

//MapToMatrix convert a variation to bi-dimensional matrix
func MapToMatrix(variation model.Variation) [][]model.Combination {
	combinations := make([][]model.Combination, len(variation))
	if len(variation) == 0 {
		return combinations
	}
	n := 0
	for attr, values := range variation {
		combination := make([]model.Combination, len(values))
		for i, value := range values {
			combination[i] = model.Combination{attr: value}
		}
		combinations[n] = combination
		n++
	}
	return combinations
}

//Permutation combine all elements
func Permutation(variations [][]model.Combination) [][]model.Combination {
	var combinations [][]model.Combination
	if len(variations) == 0 {
		return combinations
	}
	if len(variations) == 1 {
		for _, variation := range variations[0] {
			combinations = append(combinations, []model.Combination{variation})
		}
		return combinations
	}
	calculated := Permutation(variations[1:])
	for _, variation := range variations[0] {
		for _, value := range calculated {
			combinationsToAdd := append([]model.Combination{variation}, value...)
			combinations = append(combinations, combinationsToAdd)
		}
	}
	return combinations
}

// ArrayToMap flat
func ArrayToMap(permutation []model.Combination) model.Combination {
	combination := model.Combination{}
	for _, pairs := range permutation {
		for k, v := range pairs {
			combination[k] = v
		}
	}
	return combination
}
