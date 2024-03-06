package routes

import (
	"user_service/controllers"

	"github.com/gin-gonic/gin"
)

func ApplicationV1Router(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/member-register", controllers.MemberRegister)
			auth.POST("/member-register-social", controllers.MemberRegisterSocial)
			auth.POST("/member-login", controllers.MemberLogin)
			auth.POST("/member-login-social", controllers.MemberLoginSocial)
			auth.POST("/admin-login", controllers.AdminLogin)
		}
		profile := api.Group("/profile")
		{
			profile.GET("/get-profile-list", controllers.GetUserProfileList)
			profile.GET("/get-profile", controllers.GetUserProfile)
			profile.GET("/get-profile-by/:id", controllers.GetUserProfileById)
			profile.PUT("/update-profile", controllers.UpdateProfile)
			profile.PUT("/update-fcm-token", controllers.UpdateFcmToken)
		}
		socialInfo := api.Group("/social-info")
		{
			socialInfo.POST("/create-or-update", controllers.CreateOrUpdateSocialInfo)
			socialInfo.GET("/get-social-info", controllers.GetSocialInfo)
		}
	}
}
