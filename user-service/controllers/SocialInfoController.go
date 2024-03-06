package controllers

import (
	"net/http"
	"user_service/models"
	"user_service/utils/constants"

	"github.com/gin-gonic/gin"
)

// Create Or Update Social Info
func CreateOrUpdateSocialInfo(c *gin.Context) {
	requestBody := models.SocialInfoModel{}
	err := BindJSON(c, &requestBody)
	if err != nil {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_INVALID_INPUT, err)
		return
	}
	socialInfo := models.FindSocialInfoDetailByTypeAndId(requestBody.SocialType, requestBody.SocialId)
	if socialInfo.Id == 0 {
		socialInfoNew, err := models.CreateSocialInfo(&requestBody)
		if err != nil {
			RES_ERROR_MSG(c, http.StatusBadRequest, err.Error(), nil)
			return
		}
		socialInfo = *socialInfoNew
	}
	RES_SUCCESS(c, socialInfo)
}

// Get Social Info
func GetSocialInfo(c *gin.Context) {
	type QueryParams struct {
		SocialType string `form:"social_type"`
		SocialId   string `form:"social_id"`
	}
	queryParams := QueryParams{}
	err := c.BindQuery(&queryParams)
	if err != nil {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_INVALID_INPUT, err)
		return
	}
	socialInfo := models.FindSocialInfoDetailByTypeAndId(queryParams.SocialType, queryParams.SocialId)
	RES_SUCCESS(c, socialInfo)
}
