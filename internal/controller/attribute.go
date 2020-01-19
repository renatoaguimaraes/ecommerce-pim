package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatoaguimaraes/ecm-pim/internal/model"
	"github.com/renatoaguimaraes/ecm-pim/internal/repo"
	"github.com/renatoaguimaraes/ecm-pim/internal/service"
)

// CreateAttribute create an attribute
func CreateAttribute(c *gin.Context) {
	a := new(model.Attribute)
	a.Group = c.Param("group")
	if err := c.ShouldBind(a); err == nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, []error{err})
		return
	}
	as := service.AttributeService{RepoAttr: repo.AttributeRepoMongo{}}
	if err := as.InsertAttr(a); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, []error{err})
		return
	}
	c.JSON(http.StatusOK, a)
}

// GetAttributeByID get a product by id
func GetAttributeByID(c *gin.Context) {
	as := service.AttributeService{RepoAttr: repo.AttributeRepoMongo{}}
	attr := as.FindAttrByID(c.Param("group"), c.Param("attributeId"))
	if attr == nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, attr)
}

// CreateAttributeValue create an attribute value
func CreateAttributeValue(c *gin.Context) {
	as := service.AttributeService{RepoAttrValue: repo.AttributeValueRepoMongo{}}
	av := new(model.AttributeValue)
	if err := c.ShouldBind(av); err == nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, []error{err})
		return
	}
	aID := c.Param("attributeId")
	if aID == "" {
		c.Status(http.StatusNotFound)
		return
	}
	if err := av.SetAttributeID(aID); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, []error{err})
		return

	}
	if err := as.InsertAttrValue(av); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, []error{err})
		return
	}
	c.JSON(http.StatusOK, av)
}

// GetGetAttributeValueByID get a specific variant by id
func GetGetAttributeValueByID(c *gin.Context) {
	as := service.AttributeService{RepoAttrValue: repo.AttributeValueRepoMongo{}}
	av := as.FindAttrValueByID(c.Param("attributeId"), c.Param("attributeValueId"))
	if av == nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, av)
}
