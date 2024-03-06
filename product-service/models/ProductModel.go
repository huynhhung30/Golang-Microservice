package models

import (
	"product_service/config"
	"product_service/utils/functions"
	"time"
)

type ProductModel struct {
	Id          int       `json:"id" gorm:"id;primary_key;auto_increment;not_null"`
	Name        string    `json:"name" gorm:"name"`
	Description string    `json:"description" gorm:"description"`
	ThumbImage  string    `json:"thumb_image" gorm:"thumb_image"`
	CreatedBy   int       `json:"created_by" gorm:"created_by"`
	UpdatedBy   int       `json:"updated_by" gorm:"updated_by"`
	CreatedAt   time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"updated_at"`
}

func (t *ProductModel) TableName() string {
	return "products"
}

type ProductFullModel struct {
	ProductModel
	CreatedByUser *UserModel `json:"created_by_user" gorm:"-"`
}

// Create Product
func CreateProduct(productBody *ProductModel) (*ProductModel, error) {
	productBody.CreatedAt = functions.CurrentTime()
	productBody.UpdatedAt = functions.CurrentTime()
	err := config.DB.Create(&productBody).Error
	return productBody, err
}

// Update Product
func UpdateProduct(productBody ProductModel) (ProductModel, error) {
	params := map[string]interface{}{
		"updated_by": productBody.UpdatedBy,
		"updated_at": functions.CurrentTime(),
	}
	if productBody.Name != "" {
		params["name"] = productBody.Name
	}
	if productBody.Description != "" {
		params["description"] = productBody.Description
	}
	if productBody.ThumbImage != "" {
		params["thumb_image"] = productBody.ThumbImage
	}
	err := config.DB.Model(&productBody).Debug().Where("id = ?", productBody.Id).Updates(params).Take(&productBody).Error
	return productBody, err
}

// Find Product List
func FindProductList(params PageLimitQueryModel) (list []ProductFullModel, totalCount int64) {
	config.DB.Table("products").
		Count(&totalCount).
		Limit(params.Limit).
		Offset((params.Page - 1) * params.Limit).
		Find(&list)
	for i := 0; i < len(list); i++ {
		list[i].CreatedByUser = RpcGetUserInfoById(list[i].CreatedBy)
	}
	return list, totalCount
}

// Find Product Detail By Id
func FindProductDetailById(id int) (product ProductFullModel) {
	config.DB.Model(ProductModel{}).Where("id = ?", id).Take(&product)
	product.CreatedByUser = RpcGetUserInfoById(product.CreatedBy)
	return product
}

// Delete Products By Ids
func DeleteProductsByIds(ids []int) (product []ProductModel) {
	idListStr := functions.ConvertIdListToSqlIdList(ids)
	functions.ShowLog("idListStr", idListStr)
	config.DB.Debug().Where("id IN " + functions.ConvertIdListToSqlIdList(ids)).Find(&product).Delete(ProductModel{})
	return product
}
