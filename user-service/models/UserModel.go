package models

import (
	"time"

	"user_service/config"
	"user_service/utils"

	"user_service/utils/constants"
	"user_service/utils/functions"
)

type UserModel struct {
	Id            int       `json:"id" gorm:"id;primary_key;auto_increment;not_null"`
	UserType      string    `json:"user_type" gorm:"user_type"`
	FirstName     string    `json:"first_name" gorm:"first_name"`
	LastName      string    `json:"last_name" gorm:"last_name"`
	Avatar        string    `json:"avatar" gorm:"avatar"`
	Email         string    `json:"email" gorm:"email"`
	Password      string    `json:"password" gorm:"password"`
	Status        string    `json:"status" gorm:"status"`
	EmailVerified bool      `json:"email_verified" gorm:"email_verified"`
	Address       string    `json:"address" gorm:"address"`
	PhoneNumber   string    `json:"phone_number" gorm:"phone_number"`
	LoginMethod   string    `json:"login_method" gorm:"login_method"`
	SocialId      string    `json:"social_id" gorm:"social_id"`
	CreatedAt     time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"updated_at"`
}

func (t *UserModel) TableName() string {
	return "users"
}

type UsersLoginRequestModel struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	LoginMethod string `json:"login_method"`
	UserType    string `json:"user_type"`
	SocialId    string `json:"social_id"`
}

type ProfileModel struct {
	UserModel
	Token string `json:"token" gorm:"-"`
}

// Create User
func CreateUser(userBody *UserModel) (user ProfileModel, err error) {
	userBody.CreatedAt = functions.CurrentTime()
	userBody.UpdatedAt = functions.CurrentTime()
	userBody.EmailVerified = false
	userBody.Status = constants.USER_STATUS_ACTIVE
	if userBody.Avatar == "" {
		userBody.Avatar = constants.USER_AVATAR_DEFAULT
	}
	err = config.DB.Debug().Create(&userBody).Error
	user.UserModel = *userBody
	user.Password = ""
	user.Token = utils.GenerateTokenString(user.Id, constants.USER_TYPE_MEMBER)
	return user, err
}

// Find User Info By Email
func FindUserProfileByEmail(email string, user_type string) (user *ProfileModel) {
	config.DB.Where("email = ? and user_type = ? and login_method = ?", email, user_type, constants.LOGIN_METHOD_SYSTEM).First(&user)
	return user
}

// Find User Profile List
func CheckUserTypeIsAdmin(user_id int) bool {
	var userInfo ProfileModel
	config.DB.Where("id = ?", user_id).Take(&userInfo)
	if userInfo.UserType != constants.USER_TYPE_ADMIN {
		return true
	} else {
		return false
	}
}

// Find User Profile List
func FindUserProfileList(params PageLimitQueryModel) (user []ProfileModel, totalCount int64) {
	sortParams := "id asc"
	keywordParams := ""
	if params.Keyword != "" {
		// keywordParams = ""
		keywordParams = "first_name LIKE" + "'%" + params.Keyword + "%'" + "OR last_name LIKE" + "'%" + params.Keyword + "%'" + "OR email LIKE" + "'%" + params.Keyword + "%'"

	}
	if params.Sort == constants.SORT_PARAMS_DESC {
		sortParams = "id desc"
	}
	config.DB.Table("users").
		Debug().
		Order(sortParams).
		Where(keywordParams).
		Count(&totalCount).
		Limit(params.Limit).
		Offset((params.Page - 1) * params.Limit).
		Find(&user)
	for index := range user {
		user[index].Password = ""
	}
	return user, totalCount
}

// Find User Profile By Id
func FindUserProfileById(user_id int) (user ProfileModel) {
	config.DB.Debug().
		Table("users").
		Where("id = ?", user_id).
		Take(&user)
	user.Password = ""
	return user
}

// Find User Info By Id
func FindUserInfoById(user_id int) (user UserModel) {
	config.DB.Where("id = ?", user_id).Take(&user)
	user.Password = ""
	return user
}

// Find User Profile By Social Id
func FindUserProfileBySocialId(login_method string, social_id string) (user *ProfileModel) {
	config.DB.Model(&user).Where("login_method = ? AND social_id = ?", login_method, social_id).Scan(&user)
	user.Password = ""
	return user
}

// Update User
func UpdateUser(user UserModel) (UserModel, error) {
	params := map[string]interface{}{
		"updated_at": functions.CurrentTime(),
	}
	if user.FirstName != "" {
		params["first_name"] = user.FirstName
	}
	if user.LastName != "" {
		params["last_name"] = user.LastName
	}
	if user.Avatar != "" {
		params["avatar"] = user.Avatar
	}
	if user.Address != "" {
		params["address"] = user.Address
	}
	if user.PhoneNumber != "" {
		params["phone_number"] = user.PhoneNumber
	}
	functions.ShowLog("params", params)
	err := config.DB.Model(&user).Debug().Where("id = ?", user.Id).Updates(params).Take(&user).Error
	return user, err
}
