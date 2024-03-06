package utils

import (
	"fmt"
	"product_service/config"
	"product_service/utils/functions"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	hashers "github.com/meehow/go-django-hashers"
)

type TokenModel struct {
	UserId   int
	UserType string
	Exp      time.Time
}

// Compare password with hash
func CheckPasswordHash(password string, hash string) bool {
	ok, err := hashers.CheckPassword(password, hash)
	result := true
	if err != nil {
		result = false
	} else if ok != true {
		result = false
	} else {
		result = true
	}
	return result
}

// Hash password
func HashPassword(password string) string {
	passwordHash, _ := hashers.MakePassword(password)
	return passwordHash
}

// Extract token
func ExtractToken(c *gin.Context) string {
	if len(c.Request.Header["Authorization"]) > 0 {
		bearerToken := c.Request.Header["Authorization"][0]
		if len(strings.Split(bearerToken, " ")) == 2 {
			return strings.Split(bearerToken, " ")[1]
		}
		return ""
	}
	return ""
}

// Get token info
func GetTokenInfo(c *gin.Context) (tokenModel TokenModel) {
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.SECRET), nil
	})
	if err != nil {
		tokenModel.UserId = 0
		tokenModel.UserType = ""
		tokenModel.Exp = functions.CurrentTime()
		return tokenModel
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	user_id := claims["user_id"]
	user_type := claims["user_type"]
	expInterface := claims["exp"]
	tokenModel.UserId = int(user_id.(float64))
	tokenModel.UserType = user_type.(string)
	expInt := int64(expInterface.(float64))
	exp := time.Unix(expInt, 0)
	tokenModel.Exp = exp
	return tokenModel
}

// Generate Token String
func GenerateTokenString(user_id int, user_type string) string {
	tokenModel := jwt.New(jwt.SigningMethodHS256)
	claims := tokenModel.Claims.(jwt.MapClaims)
	claims["user_id"] = user_id
	claims["user_type"] = user_type
	claims["exp"] = functions.CurrentTime().Add(time.Hour * config.EXP_HOURS).Unix()
	token, err := tokenModel.SignedString([]byte(config.SECRET))
	if err != nil {
		functions.ShowLog("GenerateTokenStringError", err.Error())
	}
	return token
}
