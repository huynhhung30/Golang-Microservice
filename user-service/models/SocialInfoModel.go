package models

import (
	"time"
	"user_service/config"
	"user_service/utils/functions"
)

type SocialInfoModel struct {
	Id          int       `json:"id" gorm:"id;primary_key;auto_increment;not_null"`
	SocialType  string    `json:"social_type" gorm:"social_type"`
	SocialId    string    `json:"social_id" gorm:"social_id"`
	FirstName   string    `json:"first_name" gorm:"first_name"`
	LastName    string    `json:"last_name" gorm:"last_name"`
	Avatar      string    `json:"avatar" gorm:"avatar"`
	Email       string    `json:"email" gorm:"email"`
	PhoneNumber string    `json:"phone_number" gorm:"phone_number"`
	CreatedAt   time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"updated_at"`
}

func (t *SocialInfoModel) TableName() string {
	return "social_infos"
}

// Create SocialInfo
func CreateSocialInfo(socialInfoBody *SocialInfoModel) (*SocialInfoModel, error) {
	socialInfoBody.CreatedAt = functions.CurrentTime()
	socialInfoBody.UpdatedAt = functions.CurrentTime()
	err := config.DB.Debug().Create(&socialInfoBody).Error
	return socialInfoBody, err
}

// Find SocialInfo Detail By Type And Id
func FindSocialInfoDetailByTypeAndId(social_type string, social_id string) (socialInfo SocialInfoModel) {
	config.DB.Debug().Model(SocialInfoModel{}).Where("social_type = ? AND social_id = ?", social_type, social_id).Take(&socialInfo)
	return socialInfo
}
