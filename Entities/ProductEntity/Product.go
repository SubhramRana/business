package ProductEntity

import (
	Base "business/Entities/BaseEntity"
	"business/Entities/OrderEntity"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func DecodeJSON(c * gin.Context,v *Product)error{
	decoder:=json.NewDecoder(c.Request.Body)
	return decoder.Decode(v)
}

type Product struct{
	Base.Base
	Name string `gorm:"unique"`
	Quantity uint32 `binding:"required"`
	Price uint32 `binding:"required"`
	Orders []OrderEntity.Order
}


