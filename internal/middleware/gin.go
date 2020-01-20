package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/renatoaguimaraes/ecommerce-pim/internal/controller"
)

// Start server api
func Start() {
	r := gin.Default()
	//
	r.POST("/portaria/visitantes", controller.CreateVisitate)
	r.GET("/portaria/visitantes", controller.ListVisitate)
	r.GET("/portaria/visitantes/:visitanteId", controller.GetVisitanteByID)
	r.PATCH("/portaria/visitantes/:visitanteId", controller.UpdateVisitate)
	// products
	r.POST("/products", controller.CreateProduct)
	r.GET("/products", controller.ListProducts)
	r.GET("/products/:productId", controller.GetProductByID)
	r.PATCH("/products/:productId", controller.UpdateProduct)
	// product variants
	r.GET("/products/:productId/variants", controller.ListVariants)
	r.GET("/products/:productId/variants/:variantId", controller.GetVariantByID)
	r.PATCH("/products/:productId/variants/:variantId", controller.UpdateVariant)
	// categories
	r.POST("/categories", controller.CreateCategory)
	r.GET("/categories", controller.ListCategories)
	r.PATCH("/categories/:categoryId", controller.UpdateCategory)
	// product assets
	r.POST("/products/:productId/assets", func(c *gin.Context) {})
	r.GET("/products/:productId/assets", func(c *gin.Context) {})
	r.GET("/products/:productId/assets/:assetId", func(c *gin.Context) {})
	r.PATCH("/products/:productId/assets/:assetId", func(c *gin.Context) {})
	r.DELETE("/products/:productId/assets/:assetId", func(c *gin.Context) {})
	// product prices
	r.POST("/products/:productId/prices", func(c *gin.Context) {})
	r.GET("/products/:productId/prices", func(c *gin.Context) {})
	r.GET("/products/:productId/prices/:priceId", func(c *gin.Context) {})
	r.PATCH("/products/:productId/prices/:priceId", func(c *gin.Context) {})
	r.DELETE("/products/:productId/prices/:priceId", func(c *gin.Context) {})
	// attribute
	r.GET("/attributes", func(c *gin.Context) {})
	r.POST("/attributes/:group", controller.CreateAttribute)
	r.GET("/attributes/:group", func(c *gin.Context) {})
	r.GET("/attributes/:group/:attributeId", controller.GetAttributeByID)
	r.PATCH("/attributes/:group/:attributeId", func(c *gin.Context) {})
	// attribute value
	r.POST("/attributes/:group/:attributeId/values", controller.CreateAttributeValue)
	r.GET("/attributes/:group/:attributeId/values", func(c *gin.Context) {})
	r.GET("/attributes/:group/:attributeId/values/:attributeValueId", controller.GetGetAttributeValueByID)
	r.PATCH("/attributes/:group/:attributeId/values/:attributeValueId", func(c *gin.Context) {})
	r.Run()
}
