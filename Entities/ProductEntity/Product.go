package ProductEntity

import (
	Base "business/Entities/BaseEntity"
	"business/Entities/OrderEntity"
)

type Product struct{
	Base.Base
	Name string `gorm:"unique"`
	Quantity uint32 `binding:"required"`
	Price uint32 `binding:"required"`
	Orders []OrderEntity.Order
}


