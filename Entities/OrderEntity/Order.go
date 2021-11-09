package OrderEntity

import (
	Base "business/Entities/BaseEntity"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func DecodeJSON(c * gin.Context,v *Order)error{
	decoder:=json.NewDecoder(c.Request.Body)
	return decoder.Decode(v)
}

type Order struct{
	Base.Base
	CustomerID uint32 `binding:"required"`
	ProductID uint32 `binding:"required"`
	Quantity uint32 `binding:"required"`
}

