package controllers

import (
	"net/http"
	"strconv"
	"sytem_service/models"
	"sytem_service/utils"
	"sytem_service/utils/constants"
	"sytem_service/utils/functions"

	"github.com/gin-gonic/gin"
)

// Create CommonCode
func CreateCommonCode(c *gin.Context) {
	tokenInfo := utils.GetTokenInfo(c)
	if tokenInfo.UserId == 0 {
		RES_ERROR_MSG(c, http.StatusUnauthorized, constants.MSG_INVALID_INPUT, nil)
		return
	}
	requestBody := models.CommonCodeModel{}
	err := BindJSON(c, &requestBody)
	if err != nil {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_INVALID_INPUT, err)
		return
	}
	commonCode, err := models.CreateCommonCode(&requestBody)
	if err != nil {
		RES_ERROR_MSG(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	RES_SUCCESS_MSG(c, commonCode, "Create CommonCode successfully")
}

// Update CommonCode
func UpdateCommonCode(c *gin.Context) {
	tokenInfo := utils.GetTokenInfo(c)
	if tokenInfo.UserId == 0 {
		RES_ERROR_MSG(c, http.StatusUnauthorized, constants.MSG_INVALID_INPUT, nil)
		return
	}
	requestBody := models.CommonCodeModel{}
	err := BindJSON(c, &requestBody)
	if err != nil {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_INVALID_INPUT, err)
		return
	}
	commonCode, err := models.UpdateCommonCode(requestBody)
	if err != nil {
		RES_ERROR_MSG(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	RES_SUCCESS_MSG(c, commonCode, "Update CommonCode successfully")
}

// Get CommonCode List
func GetCommonCodeList(c *gin.Context) {
	queryParams := models.PageLimitQueryModel{}
	err := c.BindQuery(&queryParams)
	if err != nil {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_INVALID_INPUT, err)
		return
	}
	if queryParams.Limit == 0 {
		queryParams.Limit = constants.LIMIT_DEFAULT
	}
	resultList, totalCount := models.FindCommonCodeList(queryParams)
	meta := GeneralPaginationModel{}
	meta.CurrentPage = queryParams.Page
	meta.CurrentCount = len(resultList)
	meta.TotalCount = int(totalCount)
	meta.TotalPage = functions.CalculateTotalPage(meta.TotalCount, queryParams.Limit)
	RES_LIST_SUCCESS(c, resultList, meta)
}

// Get CommonCode Detail
func GetCommonCodeDetail(c *gin.Context) {
	idParam := c.Params.ByName("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_INVALID_INPUT, nil)
		return
	}
	commonCodeDetail := models.FindCommonCodeDetailById(id)
	RES_SUCCESS_MSG(c, commonCodeDetail, "Get CommonCode detail successfully")
}

// Get CommonCode Detail By Code
func GetCommonCodeDetailByCode(c *gin.Context) {
	code := c.Params.ByName("code")
	commonCodeDetail := models.FindCommonCodeDetailByCode(code)
	RES_SUCCESS_MSG(c, commonCodeDetail, "Get CommonCode detail successfully")
}

// Get CommonCodes By ParentCode
func GetCommonCodesByParentCode(c *gin.Context) {
	parentCode := c.Params.ByName("parent_code")
	commonCodeDetail := models.FindCommonCodesByParentCode(parentCode)
	RES_SUCCESS_MSG(c, commonCodeDetail, "Get CommonCode detail successfully")
}

// Delete CommonCodes
func DeleteCommonCodes(c *gin.Context) {
	type Body struct {
		Ids []int `json:"ids"`
	}
	body := Body{}
	err := BindJSON(c, &body)
	if err != nil {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_INVALID_INPUT, err)
		return
	}
	commonCodes := models.DeleteCommonCodesByIds(body.Ids)
	RES_SUCCESS_MSG(c, commonCodes, "Delete CommonCodes successfully")
}
