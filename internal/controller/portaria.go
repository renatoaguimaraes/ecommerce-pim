package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/renatoaguimaraes/ecommerce-pim/internal/model"
	"github.com/renatoaguimaraes/ecommerce-pim/internal/repo"
	"github.com/renatoaguimaraes/ecommerce-pim/internal/service"
	"net/http"
)

// CreateCategory create a category
func CreateVisitate(c *gin.Context) {
	m := new(model.Visitante)
	if err := c.ShouldBind(m); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, []error{err})
		return
	}
	s := service.PortariaService{VisitanteRepo: repo.VisitanteRepoMongo{}}
	if err := s.InsertVisitnte(m); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, []error{err})
		return
	}
	c.JSON(http.StatusOK, m)
}

// UpdateProduct update a category
func UpdateVisitate(c *gin.Context) {
	s := service.PortariaService{VisitanteRepo: repo.VisitanteRepoMongo{}}
	m := s.FindVisitnteByID(c.Param("id"))
	if c == nil {
		c.Status(http.StatusNotFound)
		return
	}
	json, _ := c.GetRawData()
	if err := Patch(json, m); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, []error{err})
		return
	}
	if err := s.UpdateVisitnte(m); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, []error{err})
		return
	}
	c.JSON(http.StatusOK, m)
}

// ListCategories list all categories
func ListVisitate(c *gin.Context) {
	s := service.PortariaService{VisitanteRepo: repo.VisitanteRepoMongo{}}
	filter := c.DefaultQuery("filter", "")
	c.JSON(http.StatusOK, s.ListVisitntes(filter))
}

// GetAttributeByID get a product by id
func GetVisitanteByID(c *gin.Context) {
	s := service.PortariaService{VisitanteRepo: repo.VisitanteRepoMongo{}}
	m := s.FindVisitnteByID(c.Param("visitanteId"))
	if m == nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, m)
}
