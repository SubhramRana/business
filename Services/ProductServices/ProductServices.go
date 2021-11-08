package ProductServices

import "C"
import (
	"business/Config"
	"business/Entities/ProductEntity"
)

func GetProducts(products *[]ProductEntity.Product) (err error) {
	err=Config.DB.Find(products).Error
	return
}
func AddProduct(product * ProductEntity.Product) (err error) {
	err=Config.DB.Create(product).Error
	return
}
func UpdateProduct(id uint32, product *ProductEntity.Product) (err error) {
	product.ID=id
	err=Config.DB.Model(product).Updates(product).Error
	return
}
