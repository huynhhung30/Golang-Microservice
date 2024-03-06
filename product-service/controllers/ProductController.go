package controllers

import (
	"net/http"
	"product_service/models"
	"product_service/utils"
	"product_service/utils/constants"
	"product_service/utils/functions"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create Product
func CreateProduct(c *gin.Context) {
	tokenInfo := utils.GetTokenInfo(c)
	if tokenInfo.UserId == 0 {
		RES_ERROR_MSG(c, http.StatusUnauthorized, constants.MSG_INVALID_INPUT, nil)
		return
	}
	requestBody := models.ProductModel{}
	err := BindJSON(c, &requestBody)
	if err != nil {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_INVALID_INPUT, err)
		return
	}
	requestBody.CreatedBy = tokenInfo.UserId
	product, err := models.CreateProduct(&requestBody)
	if err != nil {
		RES_ERROR_MSG(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	RES_SUCCESS_MSG(c, product, "Create product successfully")
}

// Update Product
func UpdateProduct(c *gin.Context) {
	tokenInfo := utils.GetTokenInfo(c)
	if tokenInfo.UserId == 0 {
		RES_ERROR_MSG(c, http.StatusUnauthorized, constants.MSG_INVALID_INPUT, nil)
		return
	}
	requestBody := models.ProductModel{}
	err := BindJSON(c, &requestBody)
	if err != nil {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_INVALID_INPUT, err)
		return
	}
	requestBody.UpdatedBy = tokenInfo.UserId
	product, err := models.UpdateProduct(requestBody)
	if err != nil {
		RES_ERROR_MSG(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	RES_SUCCESS_MSG(c, product, "Update Product successfully")
}

// Get Product List
func GetProductList(c *gin.Context) {
	queryParams := models.PageLimitQueryModel{}
	err := c.BindQuery(&queryParams)
	if err != nil {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_INVALID_INPUT, err)
		return
	}
	if queryParams.Limit == 0 {
		queryParams.Limit = constants.LIMIT_DEFAULT
	}
	resultList, totalCount := models.FindProductList(queryParams)
	meta := GeneralPaginationModel{}
	meta.CurrentPage = queryParams.Page
	meta.CurrentCount = len(resultList)
	meta.TotalCount = int(totalCount)
	meta.TotalPage = functions.CalculateTotalPage(meta.TotalCount, queryParams.Limit)
	RES_LIST_SUCCESS(c, resultList, meta)
}

// Get Product Detail
func GetProductDetail(c *gin.Context) {
	idParam := c.Params.ByName("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_INVALID_INPUT, nil)
		return
	}
	productDetail := models.FindProductDetailById(id)
	RES_SUCCESS_MSG(c, productDetail, "Get Product detail successfully")
}

// Delete Product
func DeleteProducts(c *gin.Context) {
	type Body struct {
		Ids []int `json:"ids"`
	}
	body := Body{}
	err := BindJSON(c, &body)
	if err != nil {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_INVALID_INPUT, err)
		return
	}
	products := models.DeleteProductsByIds(body.Ids)
	RES_SUCCESS_MSG(c, products, "Delete Products successfully")
}
