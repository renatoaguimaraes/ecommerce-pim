package service

import (
	"testing"

	"github.com/renatoaguimaraes/ecommerce-pim/internal/model"
)

func TestMapToMatrix(t *testing.T) {
	variation := model.Variation{}
	variation[model.AttributeRef("color")] = []model.AttributeValueRef{"red", "green"}
	variation[model.AttributeRef("size")] = []model.AttributeValueRef{"s", "m"}
	matrix := MapToMatrix(variation)
	// check matrix lenght
	if len(matrix) != 2 {
		t.Error("Matrix size should be 2")
	}
	if len(matrix[0]) != 2 || len(matrix[1]) != 2 {
		t.Error("Matrix dimension size should be 2")
	}
	// check matrix elements
	if matrix[0][0]["color"] != model.AttributeValueRef("red") {
		t.Error("Color should be red")
	}
	if matrix[0][1]["color"] != model.AttributeValueRef("green") {
		t.Error("Color should be green")
	}
	if matrix[1][0]["size"] != model.AttributeValueRef("s") {
		t.Error("Size should be s")
	}
	if matrix[1][1]["size"] != model.AttributeValueRef("m") {
		t.Error("Size should be m")
	}
}

func TestMapToMatrixEmpty(t *testing.T) {
	variation := model.Variation{}
	matrix := MapToMatrix(variation)
	if len(matrix) != 0 {
		t.Error("Size of matrix should be 0")
	}
}

func TestPermutation(t *testing.T) {
	variations := make([][]model.Combination, 2)
	variations[0] = []model.Combination{{"size": "s"}, {"size": "m"}}
	variations[1] = []model.Combination{{"color": "red"}, {"color": "blue"}}
	combinations := Permutation(variations)
	if len(combinations) != 4 {
		t.Error("Number of combinations should be 4")
	}
	if combinations[0][0]["size"] != model.AttributeValueRef("s") {
		t.Error("Number of combinations should be 2")
	}
	if combinations[0][1]["color"] != model.AttributeValueRef("red") {
		t.Error("Color should be red")
	}
	if combinations[1][0]["size"] != model.AttributeValueRef("s") {
		t.Error("Size should be s")
	}
	if combinations[1][1]["color"] != model.AttributeValueRef("blue") {
		t.Error("Color should be blue")
	}
	if combinations[2][0]["size"] != model.AttributeValueRef("m") {
		t.Error("Size should be m")
	}
	if combinations[2][1]["color"] != model.AttributeValueRef("red") {
		t.Error("Color should be red")
	}
	if combinations[3][0]["size"] != model.AttributeValueRef("m") {
		t.Error("Size should be m")
	}
	if combinations[3][1]["color"] != model.AttributeValueRef("blue") {
		t.Error("Color should be blue")
	}
}

func TestPermutationEmpty(t *testing.T) {
	var variations [][]model.Combination
	combinations := Permutation(variations)
	if len(combinations) != 0 {
		t.Error("Number of combinations should be 0")
	}
}

func TestArrayToMap(t *testing.T) {
	variations := make([]model.Combination, 2)
	variations[0] = model.Combination{"size": "s"}
	variations[1] = model.Combination{"color": "red"}
	variationMap := ArrayToMap(variations)
	if len(variationMap) != 2 {
		t.Error("Number of variations should be 2")
	}
	if variationMap["size"] != model.AttributeValueRef("s") {
		t.Error("Size should be s")
	}
	if variationMap["color"] != model.AttributeValueRef("red") {
		t.Error("Color should be blue")
	}
}

func TestArrayToMapEmpty(t *testing.T) {
	var variations []model.Combination
	variationMap := ArrayToMap(variations)
	if len(variationMap) != 0 {
		t.Error("Number of variations should be 0")
	}
}

// benchmark

func BenchmarkMapToMatrix(b *testing.B) {
	for n := 0; n < b.N; n++ {
		variation := model.Variation{}
		variation[model.AttributeRef("color")] = []model.AttributeValueRef{"red", "green"}
		variation[model.AttributeRef("size")] = []model.AttributeValueRef{"s", "m"}
		MapToMatrix(variation)
	}
}

func benchmarkPermutation(variations [][]model.Combination, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Permutation(variations)
	}
}

func BenchmarkPermutation4(b *testing.B) {
	variations := make([][]model.Combination, 2)
	variations[0] = []model.Combination{{"size": "s"}, {"size": "m"}}
	variations[1] = []model.Combination{{"color": "red"}, {"color": "blue"}}
	benchmarkPermutation(variations, b)
}

func BenchmarkPermutation8(b *testing.B) {
	variations := make([][]model.Combination, 3)
	variations[0] = []model.Combination{{"size": "s"}, {"size": "m"}}
	variations[1] = []model.Combination{{"color": "red"}, {"color": "blue"}}
	variations[2] = []model.Combination{{"sleeve": "short"}, {"sleeve": "long"}}
	benchmarkPermutation(variations, b)
}

func BenchmarkPermutation16(b *testing.B) {
	variations := make([][]model.Combination, 4)
	variations[0] = []model.Combination{{"size": "s"}, {"size": "m"}}
	variations[1] = []model.Combination{{"color": "red"}, {"color": "blue"}}
	variations[2] = []model.Combination{{"sleeve": "short"}, {"sleeve": "long"}}
	variations[3] = []model.Combination{{"condition": "new"}, {"sleeve": "pre-owned"}}
	benchmarkPermutation(variations, b)
}
func BenchmarkPermutation40(b *testing.B) {
	variations := make([][]model.Combination, 4)
	variations[0] = []model.Combination{{"size": "xs"}, {"size": "s"}, {"size": "m"}, {"size": "l"}, {"size": "xl"}}
	variations[1] = []model.Combination{{"color": "red"}, {"color": "blue"}}
	variations[2] = []model.Combination{{"sleeve": "short"}, {"sleeve": "long"}}
	variations[3] = []model.Combination{{"condition": "new"}, {"sleeve": "pre-owned"}}
	benchmarkPermutation(variations, b)
}

func BenchmarkArrayToMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		variations := make([]model.Combination, 2)
		variations[0] = model.Combination{"size": "s"}
		variations[1] = model.Combination{"color": "red"}
		ArrayToMap(variations)
	}
}
