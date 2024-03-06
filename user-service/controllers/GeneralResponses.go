package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GeneralPaginationModel struct {
	CurrentPage  int `json:"current_page"`
	CurrentCount int `json:"current_count"`
	TotalPage    int `json:"total_page"`
	TotalCount   int `json:"total_count"`
}

// Reponses simple
func RES_SIMPLE(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func RES_SUCCESS_SIMPLE(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"status":     "success",
		"message":    "",
	})
}

// Reponses success
func RES_SUCCESS(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"data":       data,
		"status":     "success",
		"message":    "",
	})
}

// Reponses success
func RES_LIST_SUCCESS(c *gin.Context, data interface{}, meta interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"data":       data,
		"meta":       meta,
		"status":     "success",
		"message":    "",
	})
}

// Reponses succes msg
func RES_SUCCESS_MSG(c *gin.Context, data interface{}, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"data":       data,
		"status":     "success",
		"message":    msg,
	})
}

// Reponses error
func RES_ERROR(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{
		"statusCode": statusCode,
		"data":       data,
		"status":     "error",
		"message":    "error",
	})
}

// Reponses error with msg
func RES_ERROR_MSG(c *gin.Context, statusCode int, msg string, data interface{}) {
	c.JSON(statusCode, gin.H{
		"statusCode": statusCode,
		"data":       data,
		"status":     "error",
		"message":    msg,
	})
}
