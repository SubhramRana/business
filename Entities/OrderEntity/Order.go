package OrderEntity

import (
	Base "business/Entities/BaseEntity"
	"encoding/json"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Order struct{
	Base.Base
	CustomerID uint32
	ProductID uint32
	Quantity uint32
}

func DecodeJSON(c * gin.Context,v *Order)error{
	decoder:=json.NewDecoder(c.Request.Body)
	return decoder.Decode(v)
}

func (o Order)Validate() error {
	return validation.ValidateStruct(&o,
			validation.Field(&o.CustomerID,validation.Required.Error("CustomerID is required")),
			validation.Field(&o.ProductID,validation.Required.Error("ProductID is required")),
			validation.Field(&o.Quantity,validation.Required.Error("Quantity is required")),
		)
}