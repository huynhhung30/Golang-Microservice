package models

import "time"

type OrderConfirmationRequestQueryModel struct {
	Page         int     `form:"page"`
	Limit        int     `form:"limit"`
	Keyword      string  `form:"keyword"`
	CartDetailId []int   `form:"cart_detail_id"`
	Discount     float64 `form:"discount"`
	Vat          float64 `form:"vat"`
	DiscountId   int     `form:"discount_id"`
}

type OrderConfirmationRequestPremiumAuctionModel struct {
	TypeOfMembership        string    `form:"type_of_membership"`
	Vat                     float64   `form:"vat"`
	DiscountId              int       `form:"discount_id"`
	PinkPrice               float64   `form:"pink_price"`
	GreenPrice              float64   `form:"green_price"`
	BluePrice               float64   `form:"blue_price"`
	PremiumAuctionAreaId    int       `form:"permium_auction_region_id"`
	VisaPayment             bool      `form:"visa_payment"`
	RequestARedBill         bool      `form:"request_a_red_bill"`
	CompanyName             string    `form:"company_name"`
	Email                   string    `form:"email"`
	TaxCode                 string    `form:"tax_code"`
	Address                 string    `form:"address"`
	BidDateFrom             time.Time `form:"bid_date_from"`
	PremiumAuctionSettingId int       `form:"premium_auction_setting_id"`
}

type PageLimitQueryModel struct {
	Page     int    `form:"page"`
	Limit    int    `form:"limit"`
	Keyword  string `form:"keyword"`
	Status   string `form:"status"`
	Sort     string `form:"sort"`
	Type     string `form:"type"`
	FromDate string `form:"from_date"`
	ToDate   string `form:"to_date"`
}

type MonthYearQueryModel struct {
	PageLimitQueryModel
	Month          int    `form:"month"`
	Year           int    `form:"year"`
	Keyword        string `form:"keyword"`
	CityId         int    `form:"city_id"`
	MembershipType string `form:"membership_type"`
}

type LikeOrDislikeRequestModel struct {
	Id     int  `json:"id"`
	IsLike bool `json:"is_like"`
}

type SaveOrNotBookMarkRequestModel struct {
	Id         int  `json:"id"`
	IsBookmark bool `json:"is_bookmark"`
}
type DesignQuery struct {
	ProjectName string `query:"project_name"`
	RoomType    string `query:"room_type"`
	Style       string `query:"style"`
	DesignType  string `query:"design_type"`
	TypeOfHouse string `query:"type_of_house"`
	Page        int    `form:"page"`
	Limit       int    `form:"limit"`
	Keyword     string `form:"keyword"`
}

type ChangeStatusTypeModel struct {
	StatusType string `json:"status_type"`
	Status     bool   `json:"status"`
}

type ContractorByTypeOfMembershipRequestQueryModel struct {
	PageLimitQueryModel
	TypeOfMembership string `form:"type_of_membership"`
}

type QueryAuctionModel struct {
	Region   string `form:"region" json:"region"`
	FromDate string `form:"from_date" json:"from_date"`
	ToDate   string `form:"to_date" json:"to_date"`
}
