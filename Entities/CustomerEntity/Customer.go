package CustomerEntity

import (
	Base "business/Entities/BaseEntity"
	"business/Entities/OrderEntity"
	validation "github.com/go-ozzo/ozzo-validation"
	"time"
)

type Customer struct{
	Base.Base
	Name string `binding:"required"`
	LastOrderTime time.Time `gorm:"default:NULL"`
	Orders []OrderEntity.Order
}

func (c Customer)Validate()(err error){
	validation.ValidateStruct(
		&c,
		validation.Field(&c.Name,validation.Required),
		)
	return
}


