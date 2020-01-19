package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatoaguimaraes/ecm-pim/internal/repo"
	"github.com/renatoaguimaraes/ecm-pim/internal/service"
)

// ListVariants list all variants of a product
func ListVariants(c *gin.Context) {
	pID := c.Param("productId")
	if pID == "" {
		c.Status(http.StatusNotFound)
		return
	}
	vs := service.VariantService{Repo: repo.VariantRepoMongo{}}
	c.JSON(http.StatusOK, vs.ListVariants(c.Param("productId")))
}

// GetVariantByID get a specific variant by id
func GetVariantByID(c *gin.Context) {
	vs := service.VariantService{Repo: repo.VariantRepoMongo{}}
	v := vs.FindVariantByID(c.Param("productId"), c.Param("variantId"))
	if v == nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, v)
}

// UpdateVariant update a product
func UpdateVariant(c *gin.Context) {
	vs := service.VariantService{Repo: repo.VariantRepoMongo{}}
	v := vs.FindVariantByID(c.Param("productId"), c.Param("variantId"))
	if v == nil {
		c.Status(http.StatusNotFound)
		return
	}
	json, _ := c.GetRawData()
	if err := Patch(json, v); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := vs.Update(v); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, []error{err})
		return
	}
	c.JSON(http.StatusOK, v)
}
