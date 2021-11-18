package ProductEntity

import (
	Base "business/Entities/BaseEntity"
	"business/Entities/OrderEntity"
	"encoding/json"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"regexp"
)

type Product struct{
	Base.Base
	Name string `gorm:"unique"`
	Quantity uint32
	Price uint32
	Orders []OrderEntity.Order
}

func DecodeJSON(c * gin.Context,v *Product)error{
	decoder:=json.NewDecoder(c.Request.Body)
	return decoder.Decode(v)
}

func (p Product)Validate() error {
	err:=validation.ValidateStruct(&p,
			validation.Field(&p.Name,validation.Match(regexp.MustCompile("")).
					Error("Product Name can't be changed")),
			validation.Field(&p.Quantity,validation.Required.Error("Quantity is required")),
			validation.Field(&p.Price,validation.Required.Error("Price is required")),
		)
	//log.Warn(err)
	return err
}


