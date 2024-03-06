package routes

import (
	"sytem_service/controllers"

	"github.com/gin-gonic/gin"
)

func ApplicationV1Router(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		commonCode := api.Group("/common-code")
		{
			commonCode.POST("/create-common-code", controllers.CreateCommonCode)
			commonCode.PUT("/update-common-code", controllers.UpdateCommonCode)
			commonCode.GET("/common-code-list", controllers.GetCommonCodeList)
			commonCode.GET("/common-code-detail/:id", controllers.GetCommonCodeDetail)
			commonCode.GET("/common-code-detail-by-code/:code", controllers.GetCommonCodeDetailByCode)
			commonCode.GET("/common-codes-by-parent-code/:parent_code", controllers.GetCommonCodesByParentCode)
			commonCode.DELETE("/delete-common-codes", controllers.DeleteCommonCodes)
		}
		media := api.Group("/media")
		{
			media.POST("/upload-files", controllers.UploadFiles)
		}
	}
}
