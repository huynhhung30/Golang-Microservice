package controllers

import (
	"net/http"
	"strconv"
	"user_service/models"
	"user_service/utils"
	"user_service/utils/constants"
	"user_service/utils/functions"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// Get User Profile List
func GetUserProfileList(c *gin.Context) {
	queryParams := models.PageLimitQueryModel{}
	err := c.BindQuery(&queryParams)
	if err != nil {
		RES_ERROR_MSG(c, http.StatusForbidden, constants.MSG_INVALID_INPUT, err)
		return
	}
	if queryParams.Limit == 0 {
		queryParams.Limit = constants.LIMIT_DEFAULT
	}
	tokenInfo := utils.GetTokenInfo(c)
	if tokenInfo.UserId == 0 {
		RES_ERROR_MSG(c, http.StatusUnauthorized, constants.MSG_TOKEN_NOT_FOUND, nil)
		return
	}
	checkUsertypeIsAdmin := models.CheckUserTypeIsAdmin(tokenInfo.UserId)
	if checkUsertypeIsAdmin != true {
		RES_ERROR_MSG(c, http.StatusMethodNotAllowed, constants.MSG_USER_NOT_ROLE, nil)
	}
	userList, totalCount := models.FindUserProfileList(queryParams)
	meta := GeneralPaginationModel{}
	meta.CurrentPage = queryParams.Page
	meta.CurrentCount = len(userList)
	meta.TotalCount = int(totalCount)
	meta.TotalPage = (int(totalCount) / queryParams.Limit) + 1
	RES_LIST_SUCCESS(c, userList, meta)
}

// Get profile user
func GetUserProfile(c *gin.Context) {
	tokenInfo := utils.GetTokenInfo(c)
	if tokenInfo.UserId == 0 {
		RES_ERROR_MSG(c, http.StatusUnauthorized, constants.MSG_TOKEN_NOT_FOUND, nil)
		return
	}
	userInfo := models.FindUserProfileById(tokenInfo.UserId)
	RES_SUCCESS(c, userInfo)
}

// GetUserProfileById

func GetUserProfileById(c *gin.Context) {
	tokenInfo := utils.GetTokenInfo(c)
	if tokenInfo.UserId == 0 {
		RES_ERROR_MSG(c, http.StatusUnauthorized, constants.MSG_TOKEN_NOT_FOUND, nil)
		return
	}
	idParam := c.Params.ByName("id")
	id, err := strconv.Atoi(idParam)
	functions.ShowLog("=-=-=-=-=-=-===========testt===============================")
	if err != nil {
		RES_ERROR_MSG(c, http.StatusUnauthorized, constants.MSG_INVALID_INPUT, nil)
		return
	}
	userInfo := models.FindUserProfileById(id)
	if userInfo.Id == 0 {
		RES_SUCCESS_SIMPLE(c, nil)
	} else {
		RES_SUCCESS(c, userInfo)
	}
}

// Update Profile
func UpdateProfile(c *gin.Context) {
	tokenInfo := utils.GetTokenInfo(c)
	if tokenInfo.UserId == 0 {
		RES_ERROR_MSG(c, http.StatusUnauthorized, constants.MSG_TOKEN_NOT_FOUND, nil)
		return
	}
	userUpdateBody := models.UserModel{}
	if err := c.ShouldBindBodyWith(&userUpdateBody, binding.JSON); err != nil {
		RES_ERROR_MSG(c, http.StatusNotFound, constants.MSG_INVALID_INPUT, err)
		return
	}
	userInfo := models.FindUserInfoById(tokenInfo.UserId)
	if userInfo.Id == 0 {
		RES_ERROR_MSG(c, http.StatusUnauthorized, constants.MSG_USER_NOT_FOUND, nil)
		return
	}
	userUpdateBody.Id = tokenInfo.UserId
	userInfo, err := models.UpdateUser(userUpdateBody)
	if err != nil {
		RES_ERROR_MSG(c, http.StatusNotFound, err.Error(), nil)
		return
	}
	RES_SUCCESS_MSG(c, userInfo, "Update profile successfully")
}

// Update Fcm Token
func UpdateFcmToken(c *gin.Context) {
	tokenInfo := utils.GetTokenInfo(c)
	if tokenInfo.UserId == 0 {
		RES_ERROR_MSG(c, http.StatusUnauthorized, constants.MSG_TOKEN_NOT_FOUND, nil)
		return
	}
	fcmTokenBody := models.FcmTokenModel{}
	if err := c.ShouldBindBodyWith(&fcmTokenBody, binding.JSON); err != nil {
		RES_ERROR_MSG(c, http.StatusNotFound, constants.MSG_INVALID_INPUT, err)
		return
	}
	fcmTokenInfo := models.FindFcmToken(fcmTokenBody.FcmToken)
	if fcmTokenInfo.Id != 0 {
		RES_SUCCESS_MSG(c, fcmTokenInfo, "Update profile successfully")
		return
	}
	fcmTokenBody.Id = tokenInfo.UserId
	fcmTokenInfoNew, err := models.UpdateFcmToken(fcmTokenBody)
	if err != nil {
		RES_ERROR_MSG(c, http.StatusNotFound, err.Error(), nil)
		return
	}
	RES_SUCCESS_MSG(c, fcmTokenInfoNew, "Update fcm token successfully")
}
