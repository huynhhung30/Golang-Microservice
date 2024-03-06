package controllers

import (
	"net/http"
	"user_service/models"
	"user_service/utils"
	"user_service/utils/constants"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// Member Register
func MemberRegister(c *gin.Context) {
	// Create models
	requestBody := models.UserModel{}
	err := c.ShouldBindBodyWith(&requestBody, binding.JSON)
	if err != nil {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_INVALID_INPUT, err)
		return
	}
	requestBody.UserType = constants.USER_TYPE_MEMBER
	requestBody.LoginMethod = constants.LOGIN_METHOD_SYSTEM
	// Check input empty
	if requestBody.Email == "" {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_INVALID_INPUT, "email is required")
		return
	}
	if requestBody.Password == "" {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_INVALID_INPUT, "password is required")
		return
	}
	if requestBody.LastName == "" {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_INVALID_INPUT, "lastName is required")
		return
	}
	// Check exist email
	us := models.FindUserProfileByEmail(requestBody.Email, constants.USER_TYPE_MEMBER)
	if us.Id != 0 {
		RES_ERROR_MSG(c, http.StatusConflict, "This email already exists", nil)
		return
	}
	// Hash password
	requestBody.Password = utils.HashPassword(requestBody.Password)
	if err != nil {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_INVALID_INPUT, err.Error())
		return
	}
	userInfo, err := models.CreateUser(&requestBody)
	if err != nil {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_INVALID_INPUT, err.Error())
		return
	}
	RES_SUCCESS_MSG(c, userInfo, "Register member sucessfully!")
}

// Member Register Social
func MemberRegisterSocial(c *gin.Context) {
	// Create models
	requestBody := models.UserModel{}
	err := c.ShouldBindBodyWith(&requestBody, binding.JSON)
	if err != nil {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_INVALID_INPUT, err)
		return
	}
	requestBody.UserType = constants.USER_TYPE_MEMBER
	// Check input empty
	if requestBody.LoginMethod == "" {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_INVALID_INPUT, "login_method is required")
		return
	}
	if requestBody.LoginMethod != constants.LOGIN_METHOD_GOOLE && requestBody.LoginMethod != constants.LOGIN_METHOD_FACEBOOK && requestBody.LoginMethod != constants.LOGIN_METHOD_APPLE {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_INVALID_INPUT, "login_method must be: google, facebook, apple")
		return
	}
	if requestBody.SocialId == "" {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_INVALID_INPUT, "social_id is required")
		return
	}
	if requestBody.LastName == "" {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_INVALID_INPUT, "last_name is required")
		return
	}
	// Check user exist
	userInfoExist := models.FindUserProfileBySocialId(requestBody.LoginMethod, requestBody.SocialId)
	if userInfoExist.Id != 0 {
		RES_SUCCESS_MSG(c, userInfoExist, "Register member social sucessfully!")
		return
	}
	userInfo, err := models.CreateUser(&requestBody)
	if err != nil {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_INVALID_INPUT, err.Error())
		return
	}
	RES_SUCCESS_MSG(c, userInfo, "Register member social sucessfully!")
}

// Member Login
func MemberLogin(c *gin.Context) {
	requestBody := models.UsersLoginRequestModel{}
	err := BindJSON(c, &requestBody)
	if err != nil {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_INVALID_INPUT, err)
		return
	}
	requestBody.LoginMethod = constants.LOGIN_METHOD_SYSTEM
	// Check input empty
	if requestBody.Email == "" || requestBody.Password == "" {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, "email and password is required", err)
		return
	}
	// Find user by email
	userInfo := models.FindUserProfileByEmail(requestBody.Email, constants.USER_TYPE_MEMBER)
	if userInfo.Id == 0 {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_EMAIL_INCORRECT, nil)
		return
	}
	// Check password
	isMathPassword := utils.CheckPasswordHash(requestBody.Password, userInfo.Password)
	if isMathPassword == false {
		RES_ERROR_MSG(c, http.StatusMethodNotAllowed, constants.MSG_PASSWORD_INCORRECT, nil)
		return
	}
	// Generate token
	userInfo.Token = utils.GenerateTokenString(userInfo.Id, constants.USER_TYPE_MEMBER)
	userInfo.Password = ""
	RES_SUCCESS(c, userInfo)
}

// Member Login Social
func MemberLoginSocial(c *gin.Context) {
	requestBody := models.UsersLoginRequestModel{}
	err := BindJSON(c, &requestBody)
	if err != nil {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_INVALID_INPUT, err)
		return
	}
	// Check input empty
	if requestBody.LoginMethod == "" || requestBody.SocialId == "" {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, "email and password is required", err)
		return
	}
	if requestBody.LoginMethod != constants.LOGIN_METHOD_GOOLE && requestBody.LoginMethod != constants.LOGIN_METHOD_FACEBOOK && requestBody.LoginMethod != constants.LOGIN_METHOD_APPLE {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_INVALID_INPUT, "login_method must be: google, facebook, apple")
		return
	}
	// Find user by socical id
	userInfo := models.FindUserProfileBySocialId(requestBody.LoginMethod, requestBody.SocialId)
	if userInfo.Id == 0 {
		RES_ERROR_MSG(c, http.StatusNotFound, constants.MSG_USER_NOT_FOUND, nil)
		return
	}
	// Generate token
	userInfo.Token = utils.GenerateTokenString(userInfo.Id, constants.USER_TYPE_MEMBER)
	RES_SUCCESS(c, userInfo)
}

// Admin Login
func AdminLogin(c *gin.Context) {
	requestBody := models.UsersLoginRequestModel{}
	err := BindJSON(c, &requestBody)
	if err != nil {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_INVALID_INPUT, err)
		return
	}
	requestBody.LoginMethod = constants.LOGIN_METHOD_SYSTEM
	// Check input empty
	if requestBody.Email == "" || requestBody.Password == "" {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, "email and password is required", err)
		return
	}
	// Find user by email
	userInfo := models.FindUserProfileByEmail(requestBody.Email, constants.USER_TYPE_ADMIN)
	if userInfo.Id == 0 {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_EMAIL_INCORRECT, nil)
		return
	}
	// Check password
	isMathPassword := utils.CheckPasswordHash(requestBody.Password, userInfo.Password)
	if isMathPassword == false {
		RES_ERROR_MSG(c, http.StatusMethodNotAllowed, constants.MSG_PASSWORD_INCORRECT, nil)
		return
	}
	// Generate token
	userInfo.Token = utils.GenerateTokenString(userInfo.Id, constants.USER_TYPE_ADMIN)
	userInfo.Password = ""
	RES_SUCCESS(c, userInfo)
}
