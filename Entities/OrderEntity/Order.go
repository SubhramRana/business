package OrderEntity

import (
	Base "business/Entities/BaseEntity"
)

type Order struct{
	Base.Base
	CustomerID uint32 `binding:"required"`
	ProductID uint32 `binding:"required"`
	Quantity uint32 `binding:"required"`
}

