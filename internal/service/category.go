package service

import (
	"github.com/renatoaguimaraes/ecommerce-pim/internal/model"
	"github.com/renatoaguimaraes/ecommerce-pim/internal/repo"
)

// CategoryService service
type CategoryService struct {
	Repo repo.CategoryRepo
}

// Insert  a category
func (ps CategoryService) Insert(p *model.Category) error {
	if p.ParentPath == "" {
		p.ParentPath = "/"
	}
	if p.Path == "" {
		p.Path = "/"
	}
	if err := ps.Repo.Insert(p); err != nil {
		return err
	}
	return nil
}

// FindByID find by id
func (ps CategoryService) FindByID(ID string) *model.Category {
	return ps.Repo.FindByID(ID)
}

// Update a category
func (ps CategoryService) Update(p *model.Category) error {
	return ps.Repo.Update(p)
}

// List tree
func (ps CategoryService) ListAll() []model.Category {
	return ps.Repo.ListAll()
}

// ListCategories tree
func (ps CategoryService) ListCategories(path string) []model.Category {
	return ps.Repo.List(path)
}
