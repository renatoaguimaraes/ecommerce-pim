package service

import (
	"github.com/renatoaguimaraes/ecommerce-pim/internal/model"
	"github.com/renatoaguimaraes/ecommerce-pim/internal/repo"
)

// VariantService service
type VariantService struct {
	Repo repo.VariantRepoMongo
}

// Insert a variant
func (vs VariantService) Insert(variants []model.Variant) error {
	return vs.Repo.Insert(variants)
}

// FindVariantByID a variant
func (vs VariantService) FindVariantByID(productID string, variantID string) *model.Variant {
	return vs.Repo.FindVariantByID(productID, variantID)
}

// Update updata a variant
func (vs VariantService) Update(v *model.Variant) error {
	return vs.Repo.Update(v)
}

// ListVariants list all variants
func (vs VariantService) ListVariants(productID string) []model.Variant {
	return vs.Repo.ListVariants(productID)
}

// Calculate all possible variants based on product variations
func (vs VariantService) Calculate(p *model.Product) []model.Variant {
	if len(p.Variation) == 0 {
		return []model.Variant{}
	}
	combinations := MapToMatrix(p.Variation)
	permutations := Permutation(combinations)
	variants := make([]model.Variant, len(permutations))
	n := 0
	for _, permutation := range permutations {
		variants[n] = model.Variant{ProductID: p.ID, Combination: ArrayToMap(permutation)}
		n++
	}
	return variants
}
