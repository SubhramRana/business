package CustomerEntity

import (
	Base "business/Entities/BaseEntity"
	"business/Entities/OrderEntity"
	"encoding/json"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"time"
)

type Customer struct{
	Base.Base
	Name string `binding:"required"`
	LastOrderTime time.Time `gorm:"default:NULL"`
	Orders []OrderEntity.Order
}

func DecodeJSON(c * gin.Context,v *Customer)error{
	decoder:=json.NewDecoder(c.Request.Body)
	return decoder.Decode(v)
}

func (c Customer)Validate()(err error){
	validation.ValidateStruct(
		&c,
		validation.Field(&c.Name,validation.Required),
		)
	return
}


