package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/renatoaguimaraes/ecm-pim/internal/model"
	"github.com/renatoaguimaraes/ecm-pim/internal/repo"
	"github.com/renatoaguimaraes/ecm-pim/internal/service"
	"net/http"
)

// CreateCategory create a category
func CreateCategory(c *gin.Context) {
	ca := new(model.Category)
	if err := c.ShouldBind(ca); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, []error{err})
		return
	}
	cs := service.CategoryService{Repo: repo.CategoryRepoMongo{}}
	if err := cs.Insert(ca); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, []error{err})
		return
	}
	c.JSON(http.StatusOK, ca)
}

// UpdateProduct update a category
func UpdateCategory(c *gin.Context) {
	cs := service.CategoryService{Repo: repo.CategoryRepoMongo{}}
	ca := cs.FindByID(c.Param("categoryId"))
	if ca == nil {
		c.Status(http.StatusNotFound)
		return
	}
	json, _ := c.GetRawData()
	if err := Patch(json, ca); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, []error{err})
		return
	}
	if err := cs.Update(ca); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, []error{err})
		return
	}
	c.JSON(http.StatusOK, ca)
}

// ListCategories list all categories
func ListCategories(c *gin.Context) {
	vs := service.CategoryService{Repo: repo.CategoryRepoMongo{}}
	path := c.DefaultQuery("path", "/")
	if path == "" {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, vs.ListCategories(path))
}
