package models

import (
	"sytem_service/config"
	"sytem_service/utils/functions"
	"time"
)

type CommonCodeModel struct {
	Id         int       `json:"id" gorm:"id;primary_key;auto_increment;not_null"`
	Code       string    `json:"code" gorm:"code"`
	ParentCode string    `json:"parent_code" gorm:"parent_code"`
	Value      string    `json:"value" gorm:"value"`
	CreatedAt  time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"updated_at"`
}

func (t *CommonCodeModel) TableName() string {
	return "common_codes"
}

// Create CommonCode
func CreateCommonCode(commonCodeBody *CommonCodeModel) (*CommonCodeModel, error) {
	commonCodeBody.CreatedAt = functions.CurrentTime()
	commonCodeBody.UpdatedAt = functions.CurrentTime()
	err := config.DB.Create(&commonCodeBody).Error
	return commonCodeBody, err
}

// Update CommonCode
func UpdateCommonCode(commonCodeBody CommonCodeModel) (CommonCodeModel, error) {
	params := map[string]interface{}{
		"code":        commonCodeBody.Code,
		"parent_code": commonCodeBody.ParentCode,
		"value":       commonCodeBody.Value,
		"updated_at":  functions.CurrentTime(),
	}
	err := config.DB.Model(&commonCodeBody).Debug().Where("id = ?", commonCodeBody.Id).Updates(params).Take(&commonCodeBody).Error
	return commonCodeBody, err
}

// Find CommonCode List
func FindCommonCodeList(params PageLimitQueryModel) (list []CommonCodeModel, totalCount int64) {
	config.DB.Model(CommonCodeModel{}).
		Count(&totalCount).
		Limit(params.Limit).
		Offset((params.Page - 1) * params.Limit).
		Find(&list)
	return list, totalCount
}

// Find CommonCode Detail By Id
func FindCommonCodeDetailById(id int) (commonCode CommonCodeModel) {
	config.DB.Model(CommonCodeModel{}).Where("id = ?", id).Take(&commonCode)
	return commonCode
}

// Find CommonCode Detail By Code
func FindCommonCodeDetailByCode(code string) (commonCode CommonCodeModel) {
	config.DB.Model(CommonCodeModel{}).Where("code = ?", code).Take(&commonCode)
	return commonCode
}

// Find CommonCodes Detail By ParentCode
func FindCommonCodesByParentCode(code string) (commonCodes []CommonCodeModel) {
	config.DB.Model(CommonCodeModel{}).Where("parent_code = ?", code).Find(&commonCodes)
	return commonCodes
}

// Delete CommonCodes By Ids
func DeleteCommonCodesByIds(ids []int) (commonCodes []CommonCodeModel) {
	idListStr := functions.ConvertIdListToSqlIdList(ids)
	functions.ShowLog("idListStr", idListStr)
	config.DB.Debug().Where("id IN " + functions.ConvertIdListToSqlIdList(ids)).Find(&commonCodes).Delete(CommonCodeModel{})
	return commonCodes
}
