package models

import (
	"time"
	"user_service/config"
	"user_service/utils/functions"
)

type FcmTokenModel struct {
	Id        int       `json:"id" gorm:"id;primary_key;auto_increment;not_null"`
	UserId    string    `json:"user_id" gorm:"user_id"`
	FcmToken  string    `json:"fcm_token" gorm:"fcm_token"`
	Os        string    `json:"os" gorm:"os"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
}

func (t *FcmTokenModel) TableName() string {
	return "fcm_tokens"
}

// Update Fcm Token
func UpdateFcmToken(fcmTokenBody FcmTokenModel) (FcmTokenModel, error) {
	fcmTokenBody.CreatedAt = functions.CurrentTime()
	fcmTokenBody.UpdatedAt = functions.CurrentTime()
	err := config.DB.Create(&fcmTokenBody).Error
	return fcmTokenBody, err
}

// Find SocialInfo Detail By Type And Id
func FindFcmToken(fcm_token string) (fcmTokenInfo FcmTokenModel) {
	config.DB.Model(FcmTokenModel{}).Where("fcm_token = ?", fcm_token).Take(&fcmTokenInfo)
	return fcmTokenInfo
}
