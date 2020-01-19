package service

import (
	"github.com/renatoaguimaraes/ecm-pim/internal/model"
	"github.com/renatoaguimaraes/ecm-pim/internal/repo"
)

// ProductService service
type ProductService struct {
	Repo           repo.ProductRepo
	VariantService VariantService
}

// Insert  a product
func (ps ProductService) Insert(p *model.Product) error {
	if len(p.Categories) == 0 {
		p.Categories = []string{"/"}
	}

	if err := ps.Repo.Insert(p); err != nil {
		return err
	}

	// calculate variants based on product variation
	variants := ps.VariantService.Calculate(p)

	if err := ps.VariantService.Insert(variants); err != nil {
		return err
	}

	return nil
}

// FindProductByID find a product
func (ps ProductService) FindProductByID(productID string) *model.Product {
	return ps.Repo.FindProductByID(productID)
}

// Update a product
func (ps ProductService) Update(p *model.Product) error {
	return ps.Repo.Update(p)
}

// List tree
func (ps ProductService) ListProducts(filter map[string]string) []model.Product {
	return ps.Repo.List(filter)
}
