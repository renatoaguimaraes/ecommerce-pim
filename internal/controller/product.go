package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatoaguimaraes/ecm-pim/internal/model"
	"github.com/renatoaguimaraes/ecm-pim/internal/repo"
	"github.com/renatoaguimaraes/ecm-pim/internal/service"
)

// CreateProduct create a product
func CreateProduct(c *gin.Context) {
	prod := new(model.Product)
	if err := c.ShouldBind(prod); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, []error{err})
		return
	}
	ps := service.ProductService{Repo: repo.ProductRepoMongo{}, VariantService: service.VariantService{Repo: repo.VariantRepoMongo{}}}
	if err := ps.Insert(prod); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, []error{err})
		return
	}
	c.JSON(http.StatusOK, prod)
}

// UpdateProduct update a product
func UpdateProduct(c *gin.Context) {
	ps := service.ProductService{Repo: repo.ProductRepoMongo{}}
	prod := ps.FindProductByID(c.Param("productId"))
	if prod == nil {
		c.Status(http.StatusNotFound)
		return
	}
	json, _ := c.GetRawData()
	if err := Patch(json, prod); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, []error{err})
		return
	}
	if err := ps.Update(prod); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, []error{err})
		return
	}
	c.JSON(http.StatusOK, prod)
}

// GetProductByID get a product by id
func GetProductByID(c *gin.Context) {
	ps := service.ProductService{Repo: repo.ProductRepoMongo{}}
	prod := ps.FindProductByID(c.Param("productId"))
	if prod == nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, prod)
}

// ListProducts list all product
func ListProducts(c *gin.Context) {
	vs := service.ProductService{Repo: repo.ProductRepoMongo{}}
	filter, _ := c.GetQueryMap("filter")
	c.JSON(http.StatusOK, vs.ListProducts(filter))
}
