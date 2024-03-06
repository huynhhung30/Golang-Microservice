package models

type PageLimitQueryModel struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}
