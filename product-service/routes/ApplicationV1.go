package routes

import (
	"product_service/controllers"

	"github.com/gin-gonic/gin"
)

func ApplicationV1Router(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		product := api.Group("/product")
		{
			product.POST("/create-product", controllers.CreateProduct)
			product.PUT("/update-product/:id", controllers.UpdateProduct)
			product.GET("/product-list", controllers.GetProductList)
			product.GET("/product-detail/:id", controllers.GetProductDetail)
			product.DELETE("/delete-products", controllers.DeleteProducts)
		}
	}
}
